package errkit

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Function to resolve an error (from errkit) on a gRPC service to provide a well-structured grpc/status error response to the consumer
func BuildGRPCStatusErr(ctx context.Context, err error) error {
	if err != nil {
		opts := optionsFromCtx(ctx)
		var errGRPC ErrGRPC
		if errors.As(err, &errGRPC) {
			opts.Logger.Error("errkit: grpc error", "error", errGRPC, "instance", opts.Instance)
			return errGRPC
		} else {
			opts.Logger.Error("errkit: unknown grpc error type in response", "error", err)
		}
	}
	return err
}

// Convert a gRPC error - to errkit errors to transition the severity of the error between internal systems.
// I.e. if we get 5 NOT_FOUND from gRPC, we can send this through to the final consumer, instead of providing it as internal server error.
func MapGRPCErr(ctx context.Context, err error) (error, *status.Status) {
	if err != nil {
		opts := optionsFromCtx(ctx)
		st, ok := status.FromError(err)
		if ok {
			opts.Logger.Warn("errkit: error on grpc call", "error", err, "status", st.Code())
			switch st.Code() {
			case codes.Unknown:
				return &errUnknown{}, st
			case codes.InvalidArgument:
				for _, detail := range st.Details() {
					switch t := detail.(type) {
					case *VFDets:
						return &ErrValidationFailed{Validations: t.Validations}, st
					case *UnmarshalDetails:
						return &ErrUnmarshal{DataName: t.DataName}, st
					}
				}
				return &ErrBadRequest{Err: err}, st
			case codes.DeadlineExceeded:
				for _, detail := range st.Details() {
					switch t := detail.(type) {
					case *TimeoutDetails:
						return &ErrTimeout{Timeout: t.Timeout}, st
					}
				}
				return &ErrTimeout{Err: err}, st
			case codes.NotFound:
				for _, detail := range st.Details() {
					switch t := detail.(type) {
					case *NotFoundDetails:
						return &ErrNotFound{DataName: t.DataName, Query: t.Query}, st
					}
				}
				return &ErrNotFound{Err: err}, st
			case codes.PermissionDenied:
				for _, detail := range st.Details() {
					switch t := detail.(type) {
					case *PermissionDeniedDetails:
						return &ErrPermissionDenied{Username: t.Username, Msg: t.Msg}, st
					}
				}
				return &ErrPermissionDenied{Err: err}, st
			case codes.Internal:
				for _, detail := range st.Details() {
					switch t := detail.(type) {
					case *ConversionDetails:
						return &ErrConversion{Struct: t.Struct}, st
					case *ReadFileDetails:
						return &ErrReadFile{FilePath: t.FilePath}, st
					case *MarshalDetails:
						return &ErrMarshal{}, st
					case *InternalDetails:
						return &ErrInternal{Msg: t.Msg}, st
					}
				}
				return &ErrInternal{Err: err}, st
			case codes.Unavailable:
				for _, detail := range st.Details() {
					switch t := detail.(type) {
					case *UnavailableDetails:
						return &ErrUnavailable{Msg: t.Msg}, st
					}
				}
				return &ErrUnavailable{Err: err}, st
			case codes.Unauthenticated:
				return &ErrUnauthenticated{}, st
			case codes.Canceled:
				return context.Canceled, st
			}
			opts.Logger.Warn("errkit: no details in error to convert to http error", "error", err)
			return &errUnknown{Err: err}, st
		}
		opts.Logger.Debug("errkit: no grpc status error to convert", "error", err)
	}
	return err, nil
}
