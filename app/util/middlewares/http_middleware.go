package middlewares

import (
	"github.com/hibbannn/grpc-http-boilerplate/app/util/constants"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

func (m *Middleware) RateLimiterMiddleware() func(http.Handler) http.Handler {
	if !m.Config.Enabled {
		return func(next http.Handler) http.Handler {
			return next
		}
	}

	limiter := rate.NewLimiter(rate.Every(time.Minute/time.Duration(m.Config.RequestsPerMinute)), m.Config.RequestsPerMinute)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			if limiter.Allow() {
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			}
		})
	}
}

func (m *Middleware) SecurityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Content Security Policy
		// Content Security Policy
		w.Header().Set(constants.ContentSecurityPolicy, "default-src 'self'; script-src 'self' 'unsafe-inline'")
		// X-Content-Type-Options
		w.Header().Set(constants.XContentTypeOptions, "nosniff")
		// X-Frame-Options
		w.Header().Set(constants.XFrameOptions, "SAMEORIGIN")
		// X-XSS-Protection
		w.Header().Set(constants.XXSSProtection, "1; mode=block")
		// Strict-Transport-Security
		w.Header().Set(constants.StrictTransport, "max-age=63072000; includeSubDomains")

		// Lanjutkan dengan chain handler
		next.ServeHTTP(w, r)
	})
}
