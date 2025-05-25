package service

import (
	"context"
	"log"
	corev1 "saastack/gen/core/v1"
	"saastack/interfaces"
	"saastack/interfaces/payment/types"
	"saastack/plugins/payment"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var pluginClient map[interfaces.PluginID]types.PaymentPlugin = make(map[interfaces.PluginID]types.PaymentPlugin)

func init() {
	// Razor Pay Client
	razorpayClient := payment.NewRazorPayClient()
	pluginClient[types.RAZORPAY] = razorpayClient

	// Stripe Client
	stripeClient := payment.NewStripeClient()
	pluginClient[types.STRIPE] = stripeClient
}

type PaymentService struct {
	corev1.UnimplementedPaymentServiceServer
}

func (payment *PaymentService) Charge(_ context.Context, req *corev1.ChargePaymentRequest) (*corev1.Response, error) {
	log.Println("Payment Charge req :", req)

	client, ok := pluginClient[interfaces.PluginID(req.PluginId)]
	if !ok {
		return nil, status.Errorf(codes.Unimplemented, "invalid plugin id")
	}

	response, err := client.Charge(req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	return response, nil
}

func (payment *PaymentService) Refund(_ context.Context, req *corev1.RefundPaymentRequest) (*corev1.Response, error) {
	log.Println("Payment Refund req :", req)

	client, ok := pluginClient[interfaces.PluginID(req.PluginId)]
	if !ok {
		return nil, status.Errorf(codes.Unimplemented, "invalid plugin id")
	}

	response, err := client.Refund(req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	return response, nil
}
