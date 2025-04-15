## Custom GCP logger

Package logger offers structured logging that complies with the W3C Trace Context standard, integrating
distributed tracing via the `traceparent` header.

:exclamation: `traceparent` has a special format that consists of the following:

```
version-traceid-parentid-flags
```

`traceid` allows correlation of all operations under a single transaction across multiple systems.

### Logger Package for gRPC Calls

In the case of gRPC calls, our logger package (pkg/log) is designed to handle `traceparent` IDs and will attempt to extract it from gRPC metadata.

The context when creating a new logger instance contains the logger itself and prepares it for future gRPC calls to chain through the traceparent for correlating the logs for troubleshooting.

### Usage

#### Log messages

To log messages, initiate the logger with a context containing traceparent.

```go
log, ctx := tlog.New(context.Background())
log.Info("Processing request", "method", req.Method, "path", req.URL.Path)
```

#### Making gRPC calls

Continue using the context enriched with gRPC metadata. However, ensure that it now includes the traceparent value.

```go
ctx := tlog.GetGRPCContext(req)
```
