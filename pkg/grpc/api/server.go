package api

import (
	"net"

	"google.golang.org/grpc"
)

// NewGrpcServer launches new grpc server on a specified address
func NewGrpcServer(address string, service ExistanceServer) (server *grpc.Server, errChan <-chan error) {
	errBuf := make(chan error)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		errBuf <- err
		return nil, errBuf
	}
	server = grpc.NewServer()
	RegisterExistanceServer(server, service)

	go func() {
		err = server.Serve(listener)
		if err != nil {
			errBuf <- err
		}
	}()
	return server, errBuf
}
