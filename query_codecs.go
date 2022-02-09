package grpcds

import (
	"reflect"

	pb "github.com/guseggert/go-ds-grpc/proto"
	"github.com/ipfs/go-datastore/query"
	"google.golang.org/protobuf/proto"
)

// QueryFilterCodec provides a codec definition for RPC transfers of a query filter.
type QueryFilterCodec interface {
	Name() string
	Type() reflect.Type
	Encode(query.Filter) ([]byte, error)
	Decode([]byte) (query.Filter, error)
}

// QueryOrderCodec provides a codec definition for RPC transfers of a query order.
type QueryOrderCodec interface {
	Name() string
	Type() reflect.Type
	Encode(query.Order) ([]byte, error)
	Decode([]byte) (query.Order, error)
}

type FilterKeyCompareCodec struct{}

func (f *FilterKeyCompareCodec) Name() string       { return "FilterKeyCompare" }
func (f *FilterKeyCompareCodec) Type() reflect.Type { return reflect.TypeOf(query.FilterKeyCompare{}) }
func (f *FilterKeyCompareCodec) Encode(filter query.Filter) ([]byte, error) {
	fkc := filter.(*query.FilterKeyCompare)
	fkcPB := pb.QueryFilterKeyCompare{
		Op:  string(fkc.Op),
		Key: fkc.Key,
	}
	return proto.Marshal(&fkcPB)
}
func (f *FilterKeyCompareCodec) Decode(b []byte) (query.Filter, error) {
	fkc := pb.QueryFilterKeyCompare{}
	err := proto.Unmarshal(b, &fkc)
	if err != nil {
		return nil, err
	}
	return &query.FilterKeyCompare{
		Op:  query.Op(fkc.Op),
		Key: fkc.Key,
	}, nil
}

type FilterValueCompareCodec struct{}

func (f *FilterValueCompareCodec) Name() string { return "FilterValueCompare" }
func (f *FilterValueCompareCodec) Type() reflect.Type {
	return reflect.TypeOf(query.FilterValueCompare{})
}
func (f *FilterValueCompareCodec) Encode(filter query.Filter) ([]byte, error) {
	fvc := filter.(*query.FilterValueCompare)
	fvcPB := pb.QueryFilterValueCompare{
		Op:    string(fvc.Op),
		Value: fvc.Value,
	}
	return proto.Marshal(&fvcPB)
}

func (f *FilterValueCompareCodec) Decode(b []byte) (query.Filter, error) {
	fvc := pb.QueryFilterValueCompare{}
	err := proto.Unmarshal(b, &fvc)
	if err != nil {
		return nil, err
	}
	return &query.FilterValueCompare{
		Op:    query.Op(fvc.Op),
		Value: fvc.Value,
	}, nil
}

type FilterKeyPrefixCodec struct{}

func (f *FilterKeyPrefixCodec) Name() string       { return "FilterKeyPrefix" }
func (f *FilterKeyPrefixCodec) Type() reflect.Type { return reflect.TypeOf(query.FilterKeyPrefix{}) }
func (f *FilterKeyPrefixCodec) Encode(filter query.Filter) ([]byte, error) {
	fkp := filter.(*query.FilterKeyPrefix)
	fkpPB := pb.QueryFilterKeyPrefix{
		Prefix: fkp.Prefix,
	}
	return proto.Marshal(&fkpPB)
}

func (f *FilterKeyPrefixCodec) Decode(b []byte) (query.Filter, error) {
	fkp := pb.QueryFilterKeyPrefix{}
	err := proto.Unmarshal(b, &fkp)
	if err != nil {
		return nil, err
	}
	return &query.FilterKeyPrefix{
		Prefix: fkp.Prefix,
	}, nil
}

type OrderByValueCodec struct{}

func (o *OrderByValueCodec) Name() string                         { return "OrderByValue" }
func (o *OrderByValueCodec) Type() reflect.Type                   { return reflect.TypeOf(query.OrderByValue{}) }
func (o *OrderByValueCodec) Encode(query.Order) ([]byte, error)   { return []byte{}, nil }
func (o *OrderByValueCodec) Decode(b []byte) (query.Order, error) { return &query.OrderByValue{}, nil }

type OrderByValueDescCodec struct{}

func (o *OrderByValueDescCodec) Name() string { return "OrderByValueDesc" }
func (o *OrderByValueDescCodec) Type() reflect.Type {
	return reflect.TypeOf(query.OrderByValueDescending{})
}
func (o *OrderByValueDescCodec) Encode(query.Order) ([]byte, error) { return []byte{}, nil }
func (o *OrderByValueDescCodec) Decode(b []byte) (query.Order, error) {
	return &query.OrderByValueDescending{}, nil
}

type OrderByKeyCodec struct{}

func (o *OrderByKeyCodec) Name() string                         { return "OrderByKey" }
func (o *OrderByKeyCodec) Type() reflect.Type                   { return reflect.TypeOf(query.OrderByKey{}) }
func (o *OrderByKeyCodec) Encode(query.Order) ([]byte, error)   { return []byte{}, nil }
func (o *OrderByKeyCodec) Decode(b []byte) (query.Order, error) { return &query.OrderByKey{}, nil }

type OrderByKeyDescCodec struct{}

func (o *OrderByKeyDescCodec) Name() string { return "OrderByKeyDesc" }
func (o *OrderByKeyDescCodec) Type() reflect.Type {
	return reflect.TypeOf(query.OrderByKeyDescending{})
}
func (o *OrderByKeyDescCodec) Encode(query.Order) ([]byte, error) { return []byte{}, nil }
func (o *OrderByKeyDescCodec) Decode(b []byte) (query.Order, error) {
	return &query.OrderByKeyDescending{}, nil
}
