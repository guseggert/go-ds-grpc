package server

import (
	"context"
	"fmt"

	grpcds "github.com/guseggert/go-ds-grpc"
	pb "github.com/guseggert/go-ds-grpc/proto"
	ds "github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
)

type grpcServer struct {
	pb.UnimplementedDatastoreServer
	QueryFilterCodecs map[string]grpcds.QueryFilterCodec
	QueryOrderCodecs  map[string]grpcds.QueryOrderCodec
	DS                ds.Datastore
}

type Options struct {
	QueryFilterCodecs map[string]grpcds.QueryFilterCodec
	QueryOrderCodecs  map[string]grpcds.QueryOrderCodec
}

func WithQueryFilterCodec(codec grpcds.QueryFilterCodec) func(o *Options) {
	return func(o *Options) {
		o.QueryFilterCodecs[codec.Name()] = codec
	}
}
func WithQueryOrderCodec(codec grpcds.QueryOrderCodec) func(o *Options) {
	return func(o *Options) {
		o.QueryOrderCodecs[codec.Name()] = codec
	}
}

func New(ds ds.Datastore, optFns ...func(o *Options)) *grpcServer {
	opts := &Options{
		QueryFilterCodecs: map[string]grpcds.QueryFilterCodec{},
		QueryOrderCodecs:  map[string]grpcds.QueryOrderCodec{},
	}

	WithQueryFilterCodec(&grpcds.FilterKeyCompareCodec{})(opts)
	WithQueryFilterCodec(&grpcds.FilterValueCompareCodec{})(opts)
	WithQueryFilterCodec(&grpcds.FilterKeyPrefixCodec{})(opts)
	WithQueryOrderCodec(&grpcds.OrderByKeyCodec{})(opts)
	WithQueryOrderCodec(&grpcds.OrderByKeyDescCodec{})(opts)
	WithQueryOrderCodec(&grpcds.OrderByValueCodec{})(opts)
	WithQueryOrderCodec(&grpcds.OrderByValueDescCodec{})(opts)

	for _, optFn := range optFns {
		optFn(opts)
	}

	return &grpcServer{
		DS:                ds,
		QueryFilterCodecs: opts.QueryFilterCodecs,
		QueryOrderCodecs:  opts.QueryOrderCodecs,
	}
}

func (s *grpcServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	v, err := s.DS.Get(ctx, ds.NewKey(req.Key))
	if err != nil {
		return nil, err
	}
	return &pb.GetResponse{Value: v}, nil
}
func (s *grpcServer) Has(ctx context.Context, req *pb.HasRequest) (*pb.HasResponse, error) {
	ok, err := s.DS.Has(ctx, ds.NewKey(req.Key))
	if err != nil {
		return nil, err
	}
	return &pb.HasResponse{Has: ok}, nil
}
func (s *grpcServer) GetSize(ctx context.Context, req *pb.GetSizeRequest) (*pb.GetSizeResponse, error) {
	size, err := s.DS.GetSize(ctx, ds.NewKey(req.Key))
	if err != nil {
		return nil, err
	}
	return &pb.GetSizeResponse{Size: uint64(size)}, nil
}
func (s *grpcServer) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	return &pb.PutResponse{}, s.DS.Put(ctx, ds.NewKey(req.Key), req.Value)
}
func (s *grpcServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return &pb.DeleteResponse{}, s.DS.Delete(ctx, ds.NewKey(req.Key))
}
func (s *grpcServer) Sync(ctx context.Context, req *pb.SyncRequest) (*pb.SyncResponse, error) {
	return &pb.SyncResponse{}, s.DS.Sync(ctx, ds.NewKey(req.Prefix))
}
func (s *grpcServer) Query(req *pb.QueryRequest, stream pb.Datastore_QueryServer) error {
	filters := []query.Filter{}
	for name, value := range req.Filters {
		codec, ok := s.QueryFilterCodecs[name]
		if !ok {
			return fmt.Errorf("unknown filter '%s'", name)
		}
		filter, err := codec.Decode(value)
		if err != nil {
			return fmt.Errorf("decoding filter '%s': %s", name, err.Error())
		}
		filters = append(filters, filter)
	}

	orders := []query.Order{}
	for name, value := range req.Orders {
		codec, ok := s.QueryOrderCodecs[name]
		if !ok {
			return fmt.Errorf("unknown order '%s'", name)
		}
		order, err := codec.Decode(value)
		if err != nil {
			return fmt.Errorf("decoding order: '%s': %s", name, err.Error())
		}
		orders = append(orders, order)
	}

	results, err := s.DS.Query(stream.Context(), query.Query{
		Prefix:            req.Prefix,
		Filters:           filters,
		Orders:            orders,
		Limit:             int(req.Limit),
		Offset:            int(req.Offset),
		KeysOnly:          req.KeysOnly,
		ReturnExpirations: req.ReturnExpirations,
		ReturnsSizes:      req.ReturnSizes,
	})
	if err != nil {
		return err
	}
	defer results.Close()
	for res := range results.Next() {
		err = stream.SendMsg(&pb.QueryResult{
			Key:        res.Key,
			Value:      res.Value,
			Expiration: uint64(res.Expiration.Unix()),
			Size:       uint64(res.Size),
		})
		if err != nil {
			return err
		}
	}
	return nil
}
