package grpcds

import (
	"context"
	"fmt"
	"reflect"

	pb "github.com/guseggert/go-ds-grpc/proto"
	ds "github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
)

// datastore forwards all requests to a gRPC server.
type datastore struct {
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

func New(client pb.DatastoreClient, optFns ...func(o *Options)) *datastore {
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

	return &datastore{
		Client:            client,
		QueryFilterCodecs: opts.QueryFilterCodecs,
		QueryOrderCodecs:  opts.QueryOrderCodecs,
	}
}

func (d *datastore) Get(ctx context.Context, key ds.Key) (value []byte, err error) {
	res, err := d.Client.Get(ctx, &pb.GetRequest{Key: key.String()})
	if err != nil {
		return nil, err
	}
	return res.Value, nil
}

func (d *datastore) Has(ctx context.Context, key ds.Key) (exists bool, err error) {
	res, err := d.Client.Has(ctx, &pb.HasRequest{Key: key.String()})
	if err != nil {
		return false, err
	}
	return res.Has, nil
}

func (d *datastore) GetSize(ctx context.Context, key ds.Key) (size int, err error) {
	res, err := d.Client.GetSize(ctx, &pb.GetSizeRequest{Key: key.String()})
	if err != nil {
		return 0, err
	}
	return int(res.Size), nil
}
func (d *datastore) Put(ctx context.Context, key ds.Key, value []byte) error {
	_, err := d.Client.Put(ctx, &pb.PutRequest{Key: key.String(), Value: value})
	return err
}

func (d *datastore) Delete(ctx context.Context, key ds.Key) error {
	_, err := d.Client.Delete(ctx, &pb.DeleteRequest{Key: key.String()})
	return err
}

func (d *datastore) Sync(ctx context.Context, prefix ds.Key) error {
	_, err := d.Client.Sync(ctx, &pb.SyncRequest{Prefix: prefix.String()})
	return err
}

func (d *datastore) Close() error {
	return nil
}

func (d *datastore) Query(ctx context.Context, q query.Query) (query.Results, error) {
	filters := map[string][]byte{}
	for _, f := range q.Filters {
		typ := reflect.TypeOf(f).Elem()
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
		typ := reflect.TypeOf(o).Elem()
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
		return nil, err
	}

	iter := &queryIterator{stream: res}
	return query.ResultsFromIterator(q, query.Iterator{
		Next:  iter.Next,
		Close: iter.Close,
	}), nil
}
