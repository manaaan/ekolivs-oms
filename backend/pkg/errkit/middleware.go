package errkit

import (
	"log/slog"
	"net/http"
)

type LoggerCtx string

const LoggerInCtx LoggerCtx = "logger"

// Store errkit options in context to simplify usage and improve logging.
// If the options are not provided, errkit will be using default values to provide its logging.
func StoreOptsInCtx(opts *Options) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if opts == nil {
				// Build default options, if none are provided
				opts = &Options{
					Instance: "",
					Logger:   slog.Default(),
				}
			} else {
				if opts.Logger == nil {
					// Try to fetch logger from context values, expected from a previous middleware to be provided under `logger`
					loggerInCtx := r.Context().Value(LoggerInCtx)
					if loggerInCtx != nil {
						parsed, ok := loggerInCtx.(*slog.Logger)
						if !ok || parsed == nil {
							opts.Logger = slog.Default()
						} else {
							opts.Logger = parsed
						}
					}
				}
			}
			ctx := ContextWithErrkitOpts(r.Context(), opts)

			// Call the next handler in the chain
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
