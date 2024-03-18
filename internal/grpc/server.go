package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	grpc *grpc.Server
	port int
}

func (s *Server) GetServer() *grpc.Server {
	return s.grpc
}

func (s *Server) Serve() error {
	addr := fmt.Sprintf(":%d", s.port)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to create listener: %v", err)
	}
	fmt.Printf("Listening on port: %d\n", s.port)
	return s.grpc.Serve(l)
}

func NewServer() (*Server, error) {
	grpcServer := grpc.NewServer()

	server := &Server{
		port: 8000, // Default port
		grpc: grpcServer,
	}

	return server, nil
}
