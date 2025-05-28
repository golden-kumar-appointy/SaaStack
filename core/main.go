package core

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const CORE_ADDRESS = "localhost:9000"

func NewGrpcServer() *grpc.Server {
	return grpc.NewServer()
}

func StartServer(srv *grpc.Server) error {
	lis, err := net.Listen("tcp", CORE_ADDRESS)
	if err != nil {
		return err
	}

	log.Printf("core server listening at %v", lis.Addr())

	if err := srv.Serve(lis); err != nil {
		return err
	}

	return nil
}
