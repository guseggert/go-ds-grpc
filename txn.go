package grpcds

import (
	"context"

	"github.com/guseggert/go-ds-grpc/proto"
	ds "github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
)

type txn struct {
	id uint64
	ds *Datastore
}

func (t *txn) Get(ctx context.Context, key ds.Key) (value []byte, err error) {
	return t.ds.get(ctx, key, t.id)
}

func (t *txn) Has(ctx context.Context, key ds.Key) (exists bool, err error) {
	return t.ds.has(ctx, key, t.id)
}

func (t *txn) GetSize(ctx context.Context, key ds.Key) (size int, err error) {
	return t.ds.getSize(ctx, key, t.id)
}

func (t *txn) Query(ctx context.Context, q query.Query) (query.Results, error) {
	return t.ds.query(ctx, q, t.id)
}

func (t *txn) Put(ctx context.Context, key ds.Key, value []byte) error {
	return t.ds.put(ctx, key, value, t.id)
}

func (t *txn) Delete(ctx context.Context, key ds.Key) error {
	return t.ds.delete(ctx, key, t.id)
}

func (t *txn) Commit(ctx context.Context) error {
	_, err := t.ds.Client.CommitTransaction(ctx, &proto.CommitTransactionRequest{TransactionID: t.id})
	return GRPCToDSError(err)
}
func (t *txn) Discard(ctx context.Context) {
	_, err := t.ds.Client.DiscardTransaction(ctx, &proto.DiscardTransactionRequest{TransactionID: t.id})
	// TODO: the Discard() method on the datastore Txn interface should return an error,
	//       to accomodate these kinds of dispatching use cases
	if err != nil {
		log.Error("error discarding transaction", "TransactionID", t.id, "Error", err)
	}
}
