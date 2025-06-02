package payment

import (
	"context"
	"fmt"
	"saastack/interfaces"
	service "saastack/interfaces/payment"
	paymentv1 "saastack/interfaces/payment/proto/gen/v1"
)

const RAZORPAY_ID interfaces.PluginID = "razorpay"

type Razorpay struct{}

func (provider *Razorpay) Charge(_ context.Context, req *paymentv1.ChargePaymentRequest) (*paymentv1.Response, error) {
	fmt.Println("Razorpay.Charge request:", req)

	response := paymentv1.Response{
		Msg: "Razorpay: payment Made",
	}
	return &response, nil
}

func (provider *Razorpay) Refund(_ context.Context, req *paymentv1.RefundPaymentRequest) (*paymentv1.Response, error) {
	fmt.Println("Razorpay.Refund request:", req)

	response := paymentv1.Response{
		Msg: "Razorpay: Refund Made",
	}
	return &response, nil
}

func NewRazorPayClient() service.PaymentPlugin {
	return &Razorpay{}
}
