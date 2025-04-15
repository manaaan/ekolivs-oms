package traceparent

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

// Unique type to avoid conflict on ctx values
type TraceParent string

const Key TraceParent = "traceparent"

// GenerateCorrelationID generates a new UUID for use as a Correlation ID.
func GenerateTraceParent() string {
	return uuid.New().String()
}

// TraceParentFromContext retrieves the TraceParent from the context
// or generating a UUID if missing.
func TraceParentFromContext(ctx context.Context) (string, error) {
	if ctx == nil {
		return GenerateTraceParent(), errors.New("context is nil in TraceParentFromContext")
	}

	traceParent, ok := ctx.Value(Key).(string)
	if !ok {
		return GenerateTraceParent(), errors.New("traceparent is not set or not a string in TraceParentFromContext")
	}

	return traceParent, nil
}

// GetGRPCContext returns a context with the correlation ID from the request
// Used for gRPC calls
func GetGRPCContext(ctx context.Context, traceparent string) context.Context {
	headers := map[string]string{}
	headers[string(Key)] = traceparent
	md := metadata.New(headers)
	return metadata.NewOutgoingContext(ctx, md)
}
