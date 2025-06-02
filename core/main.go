package core

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	CORE_ADDRESS       = "localhost:9000"
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidReq      = status.Errorf(codes.Unimplemented, "invalid request")
)

func NewGrpcServer() *grpc.Server {
	return grpc.NewServer(
		grpc.UnaryInterceptor(Logger),
	)
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

func Logger(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	start := time.Now()
	m, err := handler(ctx, req)
	if err != nil {
		return nil, errInvalidReq
	}

	log.Println("Processing Time:", time.Since(start))
	log.Println("Metadata:", md)
	log.Println("Request:", req)
	log.Println("Response:", m)
	return m, err
}
