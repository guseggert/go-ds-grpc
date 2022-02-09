package server

import (
	"context"
	"fmt"
	"net"
	"testing"

	grpcds "github.com/guseggert/go-ds-grpc"
	pb "github.com/guseggert/go-ds-grpc/proto"
	ds "github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
	dssync "github.com/ipfs/go-datastore/sync"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestGRPCServer(t *testing.T) {
	ctx := context.Background()
	s := grpc.NewServer()
	pb.RegisterDatastoreServer(s, New(dssync.MutexWrap(ds.NewMapDatastore())))
	l, err := net.Listen("tcp", "127.0.0.1:8383")
	if err != nil {
		panic(err)
	}
	go s.Serve(l)
	defer s.Stop()

	conn, err := grpc.Dial("127.0.0.1:8383", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewDatastoreClient(conn)
	grpcDS := grpcds.New(client)

	for i := 0; i < 10; i++ {
		err := grpcDS.Put(ctx, ds.NewKey(fmt.Sprintf("%d", i)), []byte("foo"))
		if err != nil {
			panic(err)
		}
	}

	qr, err := grpcDS.Query(ctx, query.Query{Orders: []query.Order{&query.OrderByKey{}}})
	if err != nil {
		panic(err)
	}
	entries, err := qr.Rest()
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 10, len(entries))
	for i := 0; i < 10; i++ {
		e := entries[i]
		assert.Equal(t, []byte("foo"), e.Value)
		assert.Equal(t, fmt.Sprintf("/%d", i), e.Key)
	}
}
