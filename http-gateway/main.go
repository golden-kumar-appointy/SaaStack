package main

import (
	"context"
	"log"
	"net/http"
	emailv1 "saastack/gen/email/v1"
	paymentv1 "saastack/gen/payment/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	GRPC_SERVER_ADDRESS string = "localhost:9000"
	HTTP_SERVER_ADDRESS string = "localhost:9001"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := emailv1.RegisterEmailServiceHandlerFromEndpoint(ctx, mux, GRPC_SERVER_ADDRESS, opts); err != nil {
		panic(err)
	}
	if err := paymentv1.RegisterPaymentServiceHandlerFromEndpoint(ctx, mux, GRPC_SERVER_ADDRESS, opts); err != nil {
		panic(err)
	}

	log.Println("HTTP PROXY Server started on ", HTTP_SERVER_ADDRESS)

	if err := http.ListenAndServe(HTTP_SERVER_ADDRESS, mux); err != nil {
		panic(err)
	}
}
