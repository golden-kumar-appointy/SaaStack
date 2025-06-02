package main

import (
	"context"
	"fmt"
	"log"
	"net"
	paymentv1 "saastack/gen/payment/v1"
	"saastack/interfaces"
	service "saastack/interfaces/payment"

	"google.golang.org/grpc"
)

type CustomPayment struct {
	paymentv1.UnimplementedPaymentServiceServer
}

const (
	CUSTOM_ID      interfaces.PluginID = "custom"
	PLUGIN_ADDRESS string              = "localhost:9003"
)

func (provider *CustomPayment) Charge(_ context.Context, req *paymentv1.ChargePaymentRequest) (*paymentv1.Response, error) {
	fmt.Println("CustomPayment.Refund request:", req)

	response := paymentv1.Response{
		Msg: "Custom.Charge: Payment is charge",
	}
	return &response, nil
}

func (provider *CustomPayment) Refund(_ context.Context, req *paymentv1.RefundPaymentRequest) (*paymentv1.Response, error) {
	fmt.Println("CustomPayment.Refund request:", req)

	response := paymentv1.Response{
		Msg: "CustomPayment: Refund Made",
	}
	return &response, nil
}

func NewCustomPayment() service.PaymentPlugin {
	return &CustomPayment{}
}

func main() {
	lis, err := net.Listen("tcp", PLUGIN_ADDRESS)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	paymentv1.RegisterPaymentServiceServer(grpcServer, &CustomPayment{})

	log.Printf("custom payment plugin server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
