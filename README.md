This is a gRPC shim for [go-datastore](https://github.com/ipfs/go-datastore). 

The server wraps a datastore in a gRPC server, so that datastores can run out-of-process. If you have a special datastore implementation, simply wrap your datastore in this server and run it, and your datastore can then be called via gRPC from another process.

The client implements the datastore interface, and calls out to a remote gRPC server.

## Why not HTTP?
Theoretically this could be implemented with HTTP, but streaming results of unknown size with proper error messaging is complicated with HTTP, requiring a bespoke serialization format on top of chunked encoding, and each client language would need to implement that non-trivial logic. This is much easier with gRPC, which has streaming support and client codegen built-in.

## Filters and Orders
Filters and orders require special treatment, since they do not have codecs by default.

To support this, each filter and order must be named, and the name mapped to a codec implementation. Both client and server must do this for the desired filters and orders.

This library includes support for the default datastore filters and orders, except OrderByFunction.
