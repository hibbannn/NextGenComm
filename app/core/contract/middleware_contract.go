package contract

import (
	"context"
	"google.golang.org/grpc"
	"net/http"
)

type MiddlewareContract interface {
	UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
	RateLimiterMiddleware() func(http.Handler) http.Handler
	SecurityHeadersMiddleware(next http.Handler) http.Handler
}
