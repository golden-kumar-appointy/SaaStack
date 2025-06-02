package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"saastack/interfaces"
	service "saastack/interfaces/email"
	emailv1 "saastack/interfaces/email/proto/gen/v1"

	"google.golang.org/grpc"
)

const (
	CUSTOM_ID      interfaces.PluginID = "custom"
	PLUGIN_ADDRESS string              = "localhost:9002"
)

type CustomEmail struct {
	emailv1.UnimplementedEmailServiceServer
}

func (provider *CustomEmail) SendEmail(_ context.Context, req *emailv1.SendEmailRequest) (*emailv1.Response, error) {
	fmt.Println("Custom.sendEmail request:", req)

	response := emailv1.Response{
		Msg: "Custom: sent Email",
	}
	return &response, nil
}

func NewCustomEmail() service.EmailPlugin {
	return &CustomEmail{}
}

func main() {
	lis, err := net.Listen("tcp", PLUGIN_ADDRESS)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	emailv1.RegisterEmailServiceServer(grpcServer, &CustomEmail{})

	log.Printf("custom email plugin server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
