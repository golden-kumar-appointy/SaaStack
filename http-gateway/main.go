package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	corev1 "saastack/gen/core/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

const (
	GRPC_SERVER_ADDRESS string = "localhost:9000"
	HTTP_SERVER_ADDRESS string = "localhost:9001"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := corev1.RegisterEmailServiceHandlerFromEndpoint(ctx, mux, GRPC_SERVER_ADDRESS, opts); err != nil {
		panic(err)
	}
	if err := corev1.RegisterPaymentServiceHandlerFromEndpoint(ctx, mux, GRPC_SERVER_ADDRESS, opts); err != nil {
		panic(err)
	}

	log.Println("HTTP PROXY Server started on ", HTTP_SERVER_ADDRESS)
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(HTTP_SERVER_ADDRESS, mux)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
