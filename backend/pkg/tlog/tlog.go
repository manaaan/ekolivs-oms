// Package logger offers structured logging for GCP that complies with the W3C Trace Context standard, integrating
// distributed tracing via the `traceparent` header.

// Features:
// - Extracts `traceparent` from HTTP/gRPC for correlated logs.
// - Generates UUID for logs without `traceparent`.
// - Integrates with HTTP and gRPC services.
// - Logs source information where the error was logged from
// - Integrates to GCP severity levels to be categorized correctly

// Usage Example:
//
//	func handler(w http.ResponseWriter, r *http.Request) {
//	    log, ctx := tlog.New(context.Background())
//	    log.Info("Handling request", "path", r.URL.Path)
//	}
package tlog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/manaaan/ekolivs-oms/pkg/tlog/traceparent"
	"google.golang.org/grpc/metadata"
)

// Unique type to avoid conflict on ctx values
type Logger string

const Key Logger = "logger"

// New creates a new logger instance prioritizing data extraction from gRPC metadata,
// falling back to HTTP context or a fallback.
//
// Will firstly attempt to reuse logger from ctx
func New(ctx context.Context) (*slog.Logger, context.Context) {
	if logger := ctx.Value(Key); logger != nil {
		if resLogger, ok := logger.(*slog.Logger); ok {
			return resLogger, ctx
		}
	}

	var traceParent string

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		// Attempt to extract traceparent from gRPC metadata
		if values := md.Get(string(traceparent.Key)); len(values) > 0 {
			traceParent = values[0]
		}
	}

	// Fallback to extracting traceParent from the regular context, falling back to new UUID.
	if traceParent == "" {
		traceParent, _ = traceparent.TraceParentFromContext(ctx)
	}

	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				// Rename the level key from "level" to "severity", to have GCP interpret it correctly
				a.Key = "severity"
			}
			if a.Key == slog.SourceKey {
				// Hijacking the source key with a slimmed string value instead of the json group.
				// If we want all fields from AddSource we should rename the key to "sourceLocation"
				// for better logging in GCP as defined here: https://cloud.google.com/logging/docs/agent/logging/configuration#special-fields
				if src, ok := a.Value.Any().(*slog.Source); ok {
					// Get the two levels above the base file name
					dir, file := filepath.Split(src.File)
					parentDir := filepath.Base(dir)
					grandparentDir := filepath.Base(filepath.Dir(filepath.Dir(dir)))
					a.Value = slog.StringValue(fmt.Sprintf("%s/%s/%s:%d", grandparentDir, parentDir, file, src.Line))
				}
			}
			return a
		}}))

	logger := log.With(slog.String("traceparent", traceParent))

	ctx = context.WithValue(ctx, Key, logger)
	grpcCtx := traceparent.GetGRPCContext(ctx, traceParent)

	return logger, grpcCtx
}
