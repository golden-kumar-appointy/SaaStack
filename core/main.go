package main

import (
	"log"
	"net"
	corev1 "saastack/gen/core/v1"
	emailService "saastack/interfaces/email"
	paymentService "saastack/interfaces/payment"

	"google.golang.org/grpc"
)

const CORE_ADDRESS = "localhost:9000"

func main() {
	lis, err := net.Listen("tcp", CORE_ADDRESS)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	// Register service handler
	corev1.RegisterEmailServiceServer(grpcServer, &emailService.EmailService{})
	corev1.RegisterPaymentServiceServer(grpcServer, &paymentService.PaymentService{})

	log.Printf("core server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
