#!/bin/sh

set -e

protoc --go_out=. --go_opt=Mds.proto=./proto --go-grpc_out=. ds.proto
