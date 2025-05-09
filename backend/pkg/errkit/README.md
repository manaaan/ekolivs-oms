# Overview

Package errkit implements error types to be more declarative when throwing errors and allow a unified way of resolving them.

Example error:

```go
&errkit.ErrBadRequest{Err: err}
```

This line initiates an error for a bad request, saying that there is something wrong with the given request which makes it unusable.
It wraps an (potentially unknown) error `err` within it.

It helps resolving the errors into a format abiding by [RFC 9457](https://datatracker.ietf.org/doc/html/rfc9457#name-the-problem-details-json-ob), to help the consumers to understand the problem and resolve it.

## Creating an error from the errkit

`errkit` provides a list of common error types, that you can use in your application to specify what kind of error you're throwing.

It aims to support you in providing additional details for the error in a common and understandable format.

Example on a failed data fetch:

```go
return nil, &errkit.ErrNotFound{Err: err, DataName: "product", Query: productID}
```

## Implement types

Each error needs to satisfy the interface `ErrWrapper` and `ErrGRPC` to be able to be used for properly resolving the error.

For gRPC APIs the `GRPCStatus()` method needs to be implemented.

```go
type ErrGRPC interface {
	Error() string
	JsonRes(opts *Options) ([]byte, error)
	Unwrap() error
	// Responding with the gRPC status code for this error with its details
	GRPCStatus(opts *Options) *status.Status
}
```

You can implement your domain specific error types with consideration of these interfaces and they will be resolved as any other errkit error type.
