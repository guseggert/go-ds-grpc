package grpcds

import (
	"io"
	"sync"
	"time"

	pb "github.com/guseggert/go-ds-grpc/proto"
	"github.com/ipfs/go-datastore/query"
)

// queryIterator wraps a gRPC stream of query results with methods that can be used for query.ResultsFromIterator
type queryIterator struct {
	stream pb.Datastore_QueryClient

	closedMut sync.RWMutex
	closed    bool
}

func (q *queryIterator) Next() (query.Result, bool) {
	// mimic channel behavior, which is what consumers expect
	if q.isClosed() {
		return query.Result{}, false
	}

	select {
	case <-q.stream.Context().Done():
		q.Close()
		// note that it's possible to race here and end up sending this twice
		// but that's okay, reading from closed channel yields default value repeatedly
		return query.Result{}, false
	default:
	}

	queryResult, err := q.stream.Recv()
	if err == io.EOF {
		q.Close()
		return query.Result{}, false
	}
	if err != nil {
		q.Close()
		return query.Result{Error: err}, true
	}
	return query.Result{
		Entry: query.Entry{
			Key:        queryResult.Key,
			Value:      queryResult.Value,
			Expiration: time.Unix(int64(queryResult.Expiration), 0),
			Size:       int(queryResult.Size),
		},
	}, true
}

func (q *queryIterator) isClosed() bool {
	q.closedMut.RLock()
	defer q.closedMut.RUnlock()
	return q.closed
}

func (q *queryIterator) close() {
}

func (q *queryIterator) Close() error {
	q.closedMut.Lock()
	q.closed = true
	q.closedMut.Unlock()
	return nil
}
