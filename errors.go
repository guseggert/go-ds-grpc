package grpcds

import (
	"errors"

	ds "github.com/ipfs/go-datastore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	errMsgBatchUnsupported = "BatchUnsupported"
)

// DSToGRPCError converts a datastore error into a well-formed gRPC error.
func DSToGRPCError(err error) *status.Status {
	if errors.Is(err, ds.ErrNotFound) {
		return status.New(codes.NotFound, err.Error())
	}
	if errors.Is(err, ds.ErrBatchUnsupported) {
		return status.New(codes.Unimplemented, errMsgBatchUnsupported)
	}
	return status.New(codes.Unknown, err.Error())
}

// GRPCToDSError converts a well-formed gRPC error into a Datastore error.
func GRPCToDSError(err error) error {
	if err == nil {
		return nil
	}
	errStatus, ok := status.FromError(err)
	if ok {
		if errStatus.Code() == codes.NotFound {
			return ds.ErrNotFound
		}
		if errStatus.Message() == errMsgBatchUnsupported {
			return ds.ErrBatchUnsupported
		}

		// swallow context cancellations and deadline timeouts, which is what datastores are expected to do
		if errStatus.Code() == codes.Canceled {
			return nil
		}
		if errStatus.Code() == codes.DeadlineExceeded {
			return nil
		}
	}
	return err
}
