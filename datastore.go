package grpcds

import (
	"context"
	"fmt"
	"reflect"
	"time"

	pb "github.com/guseggert/go-ds-grpc/proto"
	ds "github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
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

func New(client pb.DatastoreClient, optFns ...func(o *Options)) *Datastore {
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

	return &Datastore{
		Client:            client,
		QueryFilterCodecs: opts.QueryFilterCodecs,
		QueryOrderCodecs:  opts.QueryOrderCodecs,
	}
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

type Batching struct {
	Client pb.DatastoreClient
}

func (d *Batching) Batch(ctx context.Context) (ds.Batch, error) {
	return &batch{client: d.Client, ops: map[ds.Key]batchOp{}}, nil
}

type batchOp struct {
	delete bool
	value  []byte
}

type batch struct {
	client pb.DatastoreClient
	ops    map[ds.Key]batchOp
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
	_, err = b.client.Batch(ctx, &batchReq)
	return GRPCToDSError(err)
}

type Checked struct {
	Client pb.DatastoreClient
}

func (d *Checked) Check(ctx context.Context) error {
	_, err := d.Client.Check(ctx, &pb.CheckRequest{})
	return GRPCToDSError(err)
}

type Scrubbed struct {
	Client pb.DatastoreClient
}

func (d *Scrubbed) Scrub(ctx context.Context) error {
	_, err := d.Client.Scrub(ctx, &pb.ScrubRequest{})
	return GRPCToDSError(err)
}

type GC struct {
	Client pb.DatastoreClient
}

func (d *GC) CollectGarbage(ctx context.Context) error {
	_, err := d.Client.CollectGarbage(ctx, &pb.CollectGarbageRequest{})
	return GRPCToDSError(err)
}

type Persistent struct {
	Client pb.DatastoreClient
}

func (d *Persistent) DiskUsage(ctx context.Context) (uint64, error) {
	res, err := d.Client.DiskUsage(ctx, &pb.DiskUsageRequest{})
	return res.Size, GRPCToDSError(err)
}

type TTL struct {
	Client pb.DatastoreClient
}

func (d *TTL) PutWithTTL(ctx context.Context, key ds.Key, value []byte, ttl time.Duration) error {
	_, err := d.Client.PutWithTTL(ctx, &pb.PutWithTTLRequest{
		Key:   key.String(),
		Value: value,
		TTL:   durationpb.New(ttl),
	})
	return GRPCToDSError(err)
}

func (d *TTL) SetTTL(ctx context.Context, key ds.Key, ttl time.Duration) error {
	_, err := d.Client.SetTTL(ctx, &pb.SetTTLRequest{
		Key: key.String(),
		TTL: durationpb.New(ttl),
	})
	return GRPCToDSError(err)
}

func (d *TTL) GetExpiration(ctx context.Context, key ds.Key) (time.Time, error) {
	res, err := d.Client.GetExpiration(ctx, &pb.GetExpirationRequest{
		Key: key.String(),
	})
	return res.Expiration.AsTime(), GRPCToDSError(err)
}
