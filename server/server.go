package server

import (
	"context"
	"fmt"
	"io/fs"
	"io/ioutil"
	"strings"

	grpcds "github.com/guseggert/go-ds-grpc"
	pb "github.com/guseggert/go-ds-grpc/proto"
	ds "github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var featuresForDS = ds.FeaturesForDatastore

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

func New(dstore ds.Datastore, optFns ...func(o *Options)) *grpcServer {
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
		DS:                dstore,
		QueryFilterCodecs: opts.QueryFilterCodecs,
		QueryOrderCodecs:  opts.QueryOrderCodecs,
	}
}

func (s *grpcServer) Features(context.Context, *pb.FeaturesRequest) (*pb.FeaturesResponse, error) {
	resp := pb.FeaturesResponse{}
	for _, f := range featuresForDS(s.DS) {
		pbFeat, ok := pb.FeaturesResponse_Feature_value[strings.ToUpper(f.Name)]
		if !ok {
			return nil, fmt.Errorf("unknown feature '%s'", f.Name)
		}
		resp.Features = append(resp.Features, pb.FeaturesResponse_Feature(pbFeat))
	}

	return &resp, nil
}

func (s *grpcServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	v, err := s.DS.Get(ctx, ds.NewKey(req.Key))
	if err != nil {
		return nil, grpcds.DSToGRPCError(err).Err()
	}
	return &pb.GetResponse{Value: v}, nil
}

func (s *grpcServer) Has(ctx context.Context, req *pb.HasRequest) (*pb.HasResponse, error) {
	ok, err := s.DS.Has(ctx, ds.NewKey(req.Key))
	if err != nil {
		return nil, grpcds.DSToGRPCError(err).Err()
	}
	return &pb.HasResponse{Has: ok}, nil
}

func (s *grpcServer) GetSize(ctx context.Context, req *pb.GetSizeRequest) (*pb.GetSizeResponse, error) {
	size, err := s.DS.GetSize(ctx, ds.NewKey(req.Key))
	if err != nil {
		return nil, grpcds.DSToGRPCError(err).Err()
	}
	return &pb.GetSizeResponse{Size: uint64(size)}, nil
}

func (s *grpcServer) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	err := s.DS.Put(ctx, ds.NewKey(req.Key), req.Value)
	if err != nil {
		return nil, grpcds.DSToGRPCError(err).Err()
	}
	return &pb.PutResponse{}, nil
}

func (s *grpcServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := s.DS.Delete(ctx, ds.NewKey(req.Key))
	if err != nil {
		return nil, grpcds.DSToGRPCError(err).Err()
	}
	return &pb.DeleteResponse{}, nil
}

func (s *grpcServer) Sync(ctx context.Context, req *pb.SyncRequest) (*pb.SyncResponse, error) {
	err := s.DS.Sync(ctx, ds.NewKey(req.Prefix))
	if err != nil {
		return nil, grpcds.DSToGRPCError(err).Err()
	}
	return &pb.SyncResponse{}, nil
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
		return grpcds.DSToGRPCError(err).Err()
	}
	defer results.Close()
	for res := range results.Next() {
		msg := &pb.QueryResult{
			Key:        res.Key,
			Value:      res.Value,
			Expiration: uint64(res.Expiration.Unix()),
			Size:       uint64(res.Size),
		}
		// pack errors as status protobufs, reusing the same logic for general API errors
		if res.Error != nil {
			writeErr := ioutil.WriteFile("/tmp/queryerr-server", []byte(fmt.Sprintf("got a query result error: %s\n\n", res.Error.Error())), fs.ModeAppend)
			if writeErr != nil {
				panic(writeErr)
			}

			grpcErr := grpcds.DSToGRPCError(res.Error)
			packedErr, err := anypb.New(grpcErr.Proto())
			if err != nil {
				return err
			}
			msg.Error = packedErr
		}

		err = stream.SendMsg(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *grpcServer) Batch(ctx context.Context, req *pb.BatchRequest) (*pb.BatchResponse, error) {
	batching, ok := s.DS.(ds.Batching)
	if !ok {
		fmt.Printf("uhhh batching is not supported asshole\n")
		return nil, grpcds.DSToGRPCError(ds.ErrBatchUnsupported).Err()
	}
	batch, err := batching.Batch(ctx)
	if err != nil {
		return nil, grpcds.DSToGRPCError(err).Err()
	}
	for _, op := range req.Ops {
		switch op.OpCode {
		case pb.BatchRequest_DELETE:
			err := batch.Delete(ctx, ds.NewKey(op.Key))
			if err != nil {
				return nil, grpcds.DSToGRPCError(err).Err()
			}
		case pb.BatchRequest_PUT:
			err := batch.Put(ctx, ds.NewKey(op.Key), op.Value)
			if err != nil {
				return nil, grpcds.DSToGRPCError(err).Err()
			}
		default:
			return nil, grpcds.DSToGRPCError(fmt.Errorf("unknown op '%s'\n", op.OpCode.String())).Err()
		}
	}
	err = batch.Commit(ctx)
	if err != nil {
		return nil, grpcds.DSToGRPCError(err).Err()
	}
	return &pb.BatchResponse{}, nil
}

func (s *grpcServer) Check(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	checked, ok := s.DS.(ds.CheckedDatastore)
	if !ok {
		return nil, status.New(codes.Unimplemented, "datastore is not a checked datastore").Err()
	}
	return &pb.CheckResponse{}, grpcds.DSToGRPCError(checked.Check(ctx)).Err()
}

func (s *grpcServer) Scrub(ctx context.Context, req *pb.ScrubRequest) (*pb.ScrubResponse, error) {
	scrubbed, ok := s.DS.(ds.ScrubbedDatastore)
	if !ok {
		return nil, status.New(codes.Unimplemented, "datastore is not a scrubbed datastore").Err()
	}
	return &pb.ScrubResponse{}, grpcds.DSToGRPCError(scrubbed.Scrub(ctx)).Err()
}

func (s *grpcServer) CollectGarbage(ctx context.Context, req *pb.CollectGarbageRequest) (*pb.CollectGarbageResponse, error) {
	gc, ok := s.DS.(ds.GCDatastore)
	if !ok {
		return nil, status.New(codes.Unimplemented, "datastore is not a GC datastore").Err()
	}
	return &pb.CollectGarbageResponse{}, grpcds.DSToGRPCError(gc.CollectGarbage(ctx)).Err()
}

func (s *grpcServer) DiskUsage(ctx context.Context, req *pb.DiskUsageRequest) (*pb.DiskUsageResponse, error) {
	persistent, ok := s.DS.(ds.PersistentDatastore)
	if !ok {
		return nil, status.New(codes.Unimplemented, "datastore is not a persistent datastore").Err()
	}
	size, err := persistent.DiskUsage(ctx)
	return &pb.DiskUsageResponse{Size: size}, grpcds.DSToGRPCError(err).Err()
}

func (s *grpcServer) PutWithTTL(ctx context.Context, req *pb.PutWithTTLRequest) (*pb.PutWithTTLResponse, error) {
	ttl, ok := s.DS.(ds.TTL)
	if !ok {
		return nil, status.New(codes.Unimplemented, "datastore is not a TTL datastore").Err()
	}
	err := ttl.PutWithTTL(ctx, ds.NewKey(req.Key), req.Value, req.TTL.AsDuration())
	return &pb.PutWithTTLResponse{}, grpcds.DSToGRPCError(err).Err()
}

func (s *grpcServer) SetTTL(ctx context.Context, req *pb.SetTTLRequest) (*pb.SetTTLResponse, error) {
	ttl, ok := s.DS.(ds.TTL)
	if !ok {
		return nil, status.New(codes.Unimplemented, "datastore is not a TTL datastore").Err()
	}
	err := ttl.SetTTL(ctx, ds.NewKey(req.Key), req.TTL.AsDuration())
	return &pb.SetTTLResponse{}, grpcds.DSToGRPCError(err).Err()
}

func (s *grpcServer) GetExpiration(ctx context.Context, req *pb.GetExpirationRequest) (*pb.GetExpirationResponse, error) {
	ttl, ok := s.DS.(ds.TTL)
	if !ok {
		return nil, status.New(codes.Unimplemented, "datastore is not a TTL datastore").Err()
	}
	expiration, err := ttl.GetExpiration(ctx, ds.NewKey(req.Key))
	return &pb.GetExpirationResponse{Expiration: timestamppb.New(expiration)}, grpcds.DSToGRPCError(err).Err()
}
