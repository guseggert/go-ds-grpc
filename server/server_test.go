package server

import (
	"context"
	"fmt"
	"net"
	"strings"
	"testing"

	grpcds "github.com/guseggert/go-ds-grpc"
	pb "github.com/guseggert/go-ds-grpc/proto"
	ds "github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
	dssync "github.com/ipfs/go-datastore/sync"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestGRPCServer(t *testing.T) {
	ctx := context.Background()
	s := grpc.NewServer()
	pb.RegisterDatastoreServer(s, New(dssync.MutexWrap(ds.NewMapDatastore())))
	l, err := net.Listen("tcp", "127.0.0.1:")
	// l, err := net.Listen("unix", "/home/gus/ipfs.sock")
	if err != nil {
		panic(err)
	}
	go s.Serve(l)
	defer s.Stop()

	conn, err := grpc.Dial(l.Addr().String(), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewDatastoreClient(conn)
	grpcDS, err := grpcds.New(ctx, client)
	if err != nil {
		panic(err)
	}

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

	assert.Implements(t, (*ds.Batching)(nil), grpcDS)
	assert.Implements(t, (*ds.CheckedDatastore)(nil), grpcDS)
	assert.Implements(t, (*ds.ScrubbedDatastore)(nil), grpcDS)
	assert.Implements(t, (*ds.GCDatastore)(nil), grpcDS)
	assert.Implements(t, (*ds.PersistentDatastore)(nil), grpcDS)

	_, ok := grpcDS.(ds.TTLDatastore)
	assert.False(t, ok, "TTLDatastore")
	_, ok = grpcDS.(ds.TxnDatastore)
	assert.False(t, ok, "TxnDatastore")

}

func TestFeatures(t *testing.T) {
	// Ensure that we can map all datastore features to protobuf features,
	// so that if a new feature is added, the build fails until the corresponding protobuf feature is added.

	oldFeaturesForDS := featuresForDS
	defer func() {
		featuresForDS = oldFeaturesForDS
	}()

	allFeatures := ds.Features()

	featuresForDS = func(dstore ds.Datastore) (features []ds.Feature) {
		return allFeatures
	}

	server := &grpcServer{}

	serverFeatures, err := server.Features(context.Background(), &pb.FeaturesRequest{})
	require.NoError(t, err)

	for i := range allFeatures {
		require.Equal(t, serverFeatures.Features[i].String(), strings.ToUpper(allFeatures[i].Name))
	}
}
