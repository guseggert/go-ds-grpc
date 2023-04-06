This is a gRPC shim for [go-datastore](https://github.com/ipfs/go-datastore). 

The server wraps a datastore in a gRPC server, so that datastores can run out-of-process. If you have a special datastore implementation, simply wrap your datastore in this server and run it, and your datastore can then be called via gRPC from another process.

The client implements the datastore interface, and calls out to a remote gRPC server.

## Why not HTTP?
This could also be implemented with HTTP, but streaming results with chunked encoding and good error messaging is complicated with HTTP, and each client language would need to implement that non-trivial logic. This is much easier with gRPC, which has streaming support and client codegen built-in. The downside is that gRPC is not available for as many programming languages as HTTP. Another downside is that you can't reuse ubiquitous HTTP caching infrastructure and load balancing. The primary use case for this is to run as a sidecar, and gRPC fits that bill well. If there's a significant need for remote datastores, an HTTP version of this might be more appropriate.

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
When the client-side datastore is constructed, it negotiates the features supported by the server-side datastore, and then returns an underlying implementation that only implements those features.

"Features" are the extra interfaces that Datastores may implement, such as `datastore.Batching`, `datastore.CheckedDatastore`, etc.

This means that client-side interface assertions will behave exactly the same as if you were using the remote datastore directly.

### Transactions
