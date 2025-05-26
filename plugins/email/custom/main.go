package main

import (
	"context"
	"fmt"
	"log"
	"net"
	corev1 "saastack/gen/core/v1"
	"saastack/interfaces"

	"google.golang.org/grpc"
)

type CustomEmail struct {
	corev1.UnimplementedEmailServiceServer
}

const (
	CUSTOM_ID      interfaces.PluginID = "custom"
	PLUGIN_ADDRESS string              = "localhost:9002"
)

func (provider *CustomEmail) SendEmail(_ context.Context, req *corev1.SendEmailRequest) (*corev1.Response, error) {
	fmt.Println("Custom.sendEmail request:", req)

	response := corev1.Response{
		Msg: "Custom: sent Email",
	}
	return &response, nil
}

func NewCustomEmail() *CustomEmail {
	return &CustomEmail{}
}

func main() {
	lis, err := net.Listen("tcp", PLUGIN_ADDRESS)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	// Register service handler
	corev1.RegisterEmailServiceServer(grpcServer, &CustomEmail{})

	log.Printf("core server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
