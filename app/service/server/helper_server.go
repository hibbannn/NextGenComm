package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/hibbannn/grpc-http-boilerplate/app/util/constants"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (s *Server) waitForShutdown() {
	// Wait for interrupt signal to gracefully shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		fmt.Printf("Failed to shutdown HTTP server: %v\n", err)
	}

	s.grpcServer.GracefulStop()
	fmt.Printf("Server gracefully stopped\n")
}

func (s *Server) startGRPCServer() {
	lis, err := net.Listen(constants.TCP, fmt.Sprintf(":%d", s.config.Server.GRPC.Port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	if err := s.grpcServer.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}

func (s *Server) startHTTPServer() {
	if err := s.httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("failed to listen and serve: %v", err)
	}
}

// StartServer runs both the GRPC and HTTP servers.
func (s *Server) StartServer(timeStart time.Time) {
	go s.startGRPCServer()
	fmt.Printf("GRPC server started on port %d\n", s.config.Server.GRPC.Port)
	go s.startHTTPServer()
	fmt.Printf("HTTP server started on port %d\n", s.config.Server.HTTP.Port)
	fmt.Printf("Server started in %s\n", time.Since(timeStart))
	s.waitForShutdown()
}
