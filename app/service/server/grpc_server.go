package server

import (
	"fmt"
	"github.com/hibbannn/grpc-http-boilerplate/app/core/logic"
	"github.com/hibbannn/grpc-http-boilerplate/stub/serviceUser"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func (s *Server) initializeGRPCServer() {
	var opts []grpc.ServerOption
	if s.config.SSL.Enabled {
		creds, err := credentials.NewServerTLSFromFile(s.config.SSL.CertFile, s.config.SSL.KeyFile)
		if err != nil {
			fmt.Printf("failed to load credentials: %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	opts = append(opts, grpc.UnaryInterceptor(s.middleware.UnaryInterceptor))

	s.grpcServer = grpc.NewServer(opts...)
	if s.config.Server.GRPC.Reflection {
		reflection.Register(s.grpcServer)
	}

	logics := logic.NewLogic(s.repo, s.cache, s.mail, s.client, s.rule, s.security, s.log)
	serviceUser.RegisterUserServiceServer(s.grpcServer, logics)

}
