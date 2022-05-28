package grpcds

import (
	"context"
	"fmt"
	"reflect"
	"time"

	pb "github.com/guseggert/go-ds-grpc/proto"
	ds "github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
	"github.com/ipfs/go-datastore/scoped"
	"google.golang.org/protobuf/types/known/durationpb"
)

// Datastore forwards all requests to a gRPC server.
type Datastore struct {
	QueryFilterCodecs map[reflect.Type]QueryFilterCodec
	QueryOrderCodecs  map[reflect.Type]QueryOrderCodec
	Client            pb.DatastoreClient
}

type Options struct {
	QueryFilterCodecs map[reflect.Type]QueryFilterCodec
	QueryOrderCodecs  map[reflect.Type]QueryOrderCodec
}

// WithQueryFilterEncoders sets the given query filter encoders. These map from a concrete type that implements
func WithQueryFilterCodec(codec QueryFilterCodec) func(o *Options) {
	return func(o *Options) {
		typ := codec.Type()
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		o.QueryFilterCodecs[typ] = codec
	}
}

func WithQueryOrderCodec(codec QueryOrderCodec) func(o *Options) {
	return func(o *Options) {
		typ := codec.Type()
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		o.QueryOrderCodecs[typ] = codec
	}
}

// New constructs a new shim client that implements the Datastore interface.
// This negotiates the datastore features that the remote datastore supports,
// and constructs an underlying representation that only implements those features,
// so that type assertions on feature interfaces work as expected.
func New(ctx context.Context, client pb.DatastoreClient, optFns ...func(o *Options)) (ds.Datastore, error) {
	opts := &Options{
		QueryFilterCodecs: map[reflect.Type]QueryFilterCodec{},
		QueryOrderCodecs:  map[reflect.Type]QueryOrderCodec{},
	}

	// always register the built-ins
	WithQueryFilterCodec(&FilterKeyCompareCodec{})(opts)
	WithQueryFilterCodec(&FilterValueCompareCodec{})(opts)
	WithQueryFilterCodec(&FilterKeyPrefixCodec{})(opts)
	WithQueryOrderCodec(&OrderByKeyCodec{})(opts)
	WithQueryOrderCodec(&OrderByKeyDescCodec{})(opts)
	WithQueryOrderCodec(&OrderByValueCodec{})(opts)
	WithQueryOrderCodec(&OrderByValueDescCodec{})(opts)

	for _, optFn := range optFns {
		optFn(opts)
	}

	grpcDS := &Datastore{
		Client:            client,
		QueryFilterCodecs: opts.QueryFilterCodecs,
		QueryOrderCodecs:  opts.QueryOrderCodecs,
	}
	resp, err := client.Features(ctx, &pb.FeaturesRequest{})
	if err != nil {
		return nil, fmt.Errorf("negotiating datastore features: %w", err)
	}

	// scope down the concrete type to implement only the methods that the
	// remote datastore supports

	var featureNames []string
	for _, f := range resp.Features {
		switch f {
		case pb.FeaturesResponse_BATCHING:
			featureNames = append(featureNames, "Batching")
		case pb.FeaturesResponse_CHECKED:
			featureNames = append(featureNames, "Checked")
		case pb.FeaturesResponse_GC:
			featureNames = append(featureNames, "GC")
		case pb.FeaturesResponse_SCRUBBED:
			featureNames = append(featureNames, "Scrubbed")
		case pb.FeaturesResponse_PERSISTENT:
			featureNames = append(featureNames, "Persistent")
		case pb.FeaturesResponse_TTL:
			featureNames = append(featureNames, "TTL")
		case pb.FeaturesResponse_TRANSACTION:
			featureNames = append(featureNames, "Transactionn")
		default:
			return nil, fmt.Errorf("unknown feature %q", f.String())
		}
	}
	features := ds.FeaturesByName(featureNames...)
	scopedDS := scoped.WithFeatures(grpcDS, features)
	return scopedDS, nil
}

func (d *Datastore) Get(ctx context.Context, key ds.Key) (value []byte, err error) {
	res, err := d.Client.Get(ctx, &pb.GetRequest{Key: key.String()})
	if err != nil {
		return nil, GRPCToDSError(err)
	}
	return res.Value, nil
}

func (d *Datastore) Has(ctx context.Context, key ds.Key) (exists bool, err error) {
	res, err := d.Client.Has(ctx, &pb.HasRequest{Key: key.String()})
	if err != nil {
		return false, GRPCToDSError(err)
	}
	return res.Has, nil
}

func (d *Datastore) GetSize(ctx context.Context, key ds.Key) (size int, err error) {
	res, err := d.Client.GetSize(ctx, &pb.GetSizeRequest{Key: key.String()})
	if err != nil {
		return 0, GRPCToDSError(err)
	}
	return int(res.Size), nil
}
func (d *Datastore) Put(ctx context.Context, key ds.Key, value []byte) error {
	_, err := d.Client.Put(ctx, &pb.PutRequest{Key: key.String(), Value: value})
	return GRPCToDSError(err)
}

func (d *Datastore) Delete(ctx context.Context, key ds.Key) error {
	_, err := d.Client.Delete(ctx, &pb.DeleteRequest{Key: key.String()})
	return GRPCToDSError(err)
}

func (d *Datastore) Sync(ctx context.Context, prefix ds.Key) error {
	_, err := d.Client.Sync(ctx, &pb.SyncRequest{Prefix: prefix.String()})
	return GRPCToDSError(err)
}

func (d *Datastore) Close() error {
	return nil
}

func (d *Datastore) Query(ctx context.Context, q query.Query) (query.Results, error) {
	filters := map[string][]byte{}
	for _, f := range q.Filters {
		typ := reflect.TypeOf(f)
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		codec, ok := d.QueryFilterCodecs[typ]
		if !ok {
			return nil, fmt.Errorf("no codec registered for query filter '%s'", typ)
		}
		b, err := codec.Encode(f)
		if err != nil {
			return nil, fmt.Errorf("encoding query filter: %w", err)
		}
		filters[codec.Name()] = b
	}

	orders := map[string][]byte{}
	for _, o := range q.Orders {
		typ := reflect.TypeOf(o)
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		codec, ok := d.QueryOrderCodecs[typ]
		if !ok {
			return nil, fmt.Errorf("no codec registered for query order '%s'", typ)
		}
		b, err := codec.Encode(o)
		if err != nil {
			return nil, fmt.Errorf("encoding query filter: %w", err)
		}
		orders[codec.Name()] = b
	}

	res, err := d.Client.Query(ctx, &pb.QueryRequest{
		Prefix:            q.Prefix,
		Filters:           filters,
		Orders:            orders,
		Limit:             uint64(q.Limit),
		Offset:            uint64(q.Offset),
		KeysOnly:          q.KeysOnly,
		ReturnExpirations: q.ReturnExpirations,
		ReturnSizes:       q.ReturnsSizes,
	})

	if err != nil {
		return nil, GRPCToDSError(err)
	}

	iter := &queryIterator{stream: res}
	return query.ResultsFromIterator(q, query.Iterator{
		Next:  iter.Next,
		Close: iter.Close,
	}), nil
}

func (d *Datastore) Batch(ctx context.Context) (ds.Batch, error) {
	return &batch{
		ds:  d,
		ops: map[ds.Key]batchOp{},
	}, nil
}

type batchOp struct {
	delete bool
	value  []byte
}

type batch struct {
	ds  *Datastore
	ops map[ds.Key]batchOp
}

func (b *batch) Put(ctx context.Context, key ds.Key, value []byte) error {
	b.ops[key] = batchOp{value: value}
	return nil
}

func (b *batch) Delete(ctx context.Context, key ds.Key) error {
	b.ops[key] = batchOp{delete: true}
	return nil
}

func (b *batch) Commit(ctx context.Context) error {
	batchReq := pb.BatchRequest{}
	var err error
	for k, op := range b.ops {
		if op.delete {
			batchReq.Ops = append(batchReq.Ops, &pb.BatchRequest_Op{
				OpCode: pb.BatchRequest_DELETE,
				Key:    k.String(),
			})
		} else {
			batchReq.Ops = append(batchReq.Ops, &pb.BatchRequest_Op{
				OpCode: pb.BatchRequest_PUT,
				Key:    k.String(),
				Value:  op.value,
			})
		}
	}
	_, err = b.ds.Client.Batch(ctx, &batchReq)
	return GRPCToDSError(err)
}

func (d *Datastore) Check(ctx context.Context) error {
	_, err := d.Client.Check(ctx, &pb.CheckRequest{})
	return GRPCToDSError(err)
}

func (d *Datastore) Scrub(ctx context.Context) error {
	_, err := d.Client.Scrub(ctx, &pb.ScrubRequest{})
	return GRPCToDSError(err)
}

func (d *Datastore) CollectGarbage(ctx context.Context) error {
	_, err := d.Client.CollectGarbage(ctx, &pb.CollectGarbageRequest{})
	return GRPCToDSError(err)
}

func (d *Datastore) DiskUsage(ctx context.Context) (uint64, error) {
	res, err := d.Client.DiskUsage(ctx, &pb.DiskUsageRequest{})
	return res.Size, GRPCToDSError(err)
}

func (d *Datastore) PutWithTTL(ctx context.Context, key ds.Key, value []byte, ttl time.Duration) error {
	_, err := d.Client.PutWithTTL(ctx, &pb.PutWithTTLRequest{
		Key:   key.String(),
		Value: value,
		TTL:   durationpb.New(ttl),
	})
	return GRPCToDSError(err)
}

func (d *Datastore) SetTTL(ctx context.Context, key ds.Key, ttl time.Duration) error {
	_, err := d.Client.SetTTL(ctx, &pb.SetTTLRequest{
		Key: key.String(),
		TTL: durationpb.New(ttl),
	})
	return GRPCToDSError(err)
}

func (d *Datastore) GetExpiration(ctx context.Context, key ds.Key) (time.Time, error) {
	res, err := d.Client.GetExpiration(ctx, &pb.GetExpirationRequest{
		Key: key.String(),
	})
	return res.Expiration.AsTime(), GRPCToDSError(err)
}

func (d *Datastore) NewTransaction(ctx context.Context, readOnly bool) (ds.Txn, error) {
	// TODO: implement
	return nil, nil
}
