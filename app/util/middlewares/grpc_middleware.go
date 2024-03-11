package middlewares

import (
	"context"
	"github.com/hibbannn/grpc-http-boilerplate/app/util/constants"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// UnaryInterceptor adalah sebuah interceptor untuk unary RPC calls.
func (m *Middleware) UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	var reqID string
	if ok {
		reqIDs := md.Get(constants.RequestId)
		if len(reqIDs) > 0 {
			reqID = reqIDs[0]
		}
	}
	if reqID == constants.EmptyString {
		reqID = constants.Unknown
	}

	// Melakukan logging sebelum memproses request
	m.Log.Infof("gRPC Request - ID: %s, Method: %s", reqID, info.FullMethod)

	// Menangani request
	h, err := handler(ctx, req)

	// Melakukan logging setelah request diproses
	if err != nil {
		m.Log.Errorf("gRPC Error - ID: %s, Error: %v", reqID, err)
	} else {
		m.Log.Infof("gRPC Response - ID: %s, Method: %s", reqID, info.FullMethod)
	}

	return h, err
}
