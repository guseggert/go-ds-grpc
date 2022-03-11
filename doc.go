package grpcds

//go:generate protoc --go_out=. --go_opt=Mds.proto=./proto --go-grpc_out=. ds.proto
//go:generate go run gen/gen.go
