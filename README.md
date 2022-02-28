This is a gRPC shim for [go-datastore](https://github.com/ipfs/go-datastore). 

The server wraps a datastore in a gRPC server, so that datastores can run out-of-process. If you have a special datastore implementation, simply wrap your datastore in this server and run it, and your datastore can then be called via gRPC from another process.

The client implements the datastore interface, and calls out to a remote gRPC server.

## Why not HTTP?
This could also be implemented with HTTP, but streaming results with chunked encoding and good error messaging is complicated with HTTP, and each client language would need to implement that non-trivial logic. This is much easier with gRPC, which has streaming support and client codegen built-in. The downside is that gRPC is not available for as many programming languages as HTTP.

## Queries
Queries support a number of challenging features for RPC:

* Results are streamed
* Filters and orders are code, not data, so we have to do extra work to use them as data
* Structured errors are embedded in results

### Filters and Orders
Filters and orders require special treatment, since they do not have codecs by default.

To support this, each filter and order must be named, and the name mapped to a codec implementation. Both client and server must do this for the desired filters and orders.

This library includes support for the default datastore filters and orders, except OrderByFunction.

### Errors
Each result can also contain a structured Go error. Some of these are meaningful (for example, when a result is not found in must return datastore.ErrNotFound). When sending an error in a result, we encode the error as a standard gRPC status protobuf in the same way we do API calls, and apply the same translation logic on the client side, so that datastore-specific errors remain consistent.

## Datastore Features
Due to limitations in the way the Datastore interfaces are defined, if you need extra Datastore interfaces such as the ones listed below, then on the client-side you need to define your own struct that embeds the desired implementations, as shown in the following example:

```go
type MegaGRPCDatastore struct {
	grpcds.Datastore
	grpcds.Batching
	grpcds.Scrubbed
	grpcds.Checked
	grpcds.Persistent
	grpcds.GC
	grpcds.TTL
}

func makeMega() *MegaGRPCDatastore {
	var client pb.DatastoreClient
	return &MegaGRPCDatastore{
		Datastore:  grpcds.Datastore{Client: client},
		Batching:   grpcds.Batching{Client: client},
		Scrubbed:   grpcds.Scrubbed{Client: client},
		Checked:    grpcds.Checked{Client: client},
		Persistent: grpcds.Persistent{Client: client},
		GC:         grpcds.GC{Client: client},
		TTL:        grpcds.TTL{Client: client},
	}
}
```

This way, your client-side datastore will match the same interfaces as the underlying server-side datastore. The interfaces that are (and aren't) implemented are important as many datastore consumers switch on interface types to determine behavior.
