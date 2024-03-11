package middlewares

import (
	"github.com/hibbannn/grpc-http-boilerplate/app/core/contract"
	"github.com/hibbannn/grpc-http-boilerplate/app/core/domain"
	"golang.org/x/time/rate"
	"time"
)

// Middleware menyimpan dependensi untuk middleware HTTP dan gRPC.
type Middleware struct {
	RateLimiter *rate.Limiter
	Config      domain.RateLimiterConfig
	Log         contract.LogContract
}

// NewMiddleware membuat dan mengembalikan instance baru dari Middleware.
func NewMiddleware(rateLimitCfg domain.RateLimiterConfig) (*Middleware, error) {
	limiter := rate.NewLimiter(rate.Every(time.Minute/time.Duration(rateLimitCfg.RequestsPerMinute)), rateLimitCfg.RequestsPerMinute)

	return &Middleware{
		RateLimiter: limiter,
		Config:      rateLimitCfg,
	}, nil
}
