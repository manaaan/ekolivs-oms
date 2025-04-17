// Package errkit implements error types to be more declarative when throwing errors.
// These common error types allow a unified way of resolving them for HTTP and gRPC error responses.
//
// Example errors:
//
//	&ErrBadRequest{Err: err}
//	&ErrNotFound{Options: server.ErrKitOptions, DataName: "product", Query: "12345", Err: err}
//
// Errors should always be initialized with pointers.
//
// The error should finally be resolved with the resolver functionality this package offers, which
// will provide the consumer with an extensive error information according to the RFC 9457 standard.
//
// You'll find more details on the usage of each error type on the types directly.
package errkit

import (
	"context"
	"log/slog"
)

type OptionsCtx string

const optionsInCtx OptionsCtx = "errkitOptions"

// Default options on errkit for your application that allows
// to consider this context when resolving the errors.
type Options struct {
	// Name of the instance that is using errkit, consisting of the service and optionally more detailed instance within that service.
	// This will be used as prefix for `instance` on the json error responses.
	//
	// Examples:
	//   "product"
	Instance string

	// Custom Logger to use within errkit
	Logger *slog.Logger
}

func optionsFromCtx(ctx context.Context) *Options {
	var opts *Options
	ctxVal := ctx.Value(optionsInCtx)
	if ctxVal != nil {
		parsed, ok := ctxVal.(*Options)
		if !ok {
			opts = &Options{
				Instance: "",
				Logger:   slog.Default(),
			}
		} else {
			opts = parsed
			if opts.Logger == nil {
				opts.Logger = slog.Default()
			}
		}
	} else {
		opts = &Options{
			Instance: "",
			Logger:   slog.Default(),
		}
	}
	return opts
}

func ContextWithErrkitOpts(ctx context.Context, opts *Options) context.Context {
	return context.WithValue(ctx, optionsInCtx, opts)
}
