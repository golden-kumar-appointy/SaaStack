package payment

import (
	"context"
	"fmt"
	paymentv1 "saastack/gen/payment/v1"
	"saastack/interfaces"
)

const STRIPE_ID interfaces.PluginID = "stripe"

type Stripe struct{}

func (provider *Stripe) Charge(_ context.Context, req *paymentv1.ChargePaymentRequest) (*paymentv1.Response, error) {
	fmt.Println("Stripe.Charge request:", req)

	response := paymentv1.Response{
		Msg: "Stripe: payment Made",
	}
	return &response, nil
}

func (provider *Stripe) Refund(_ context.Context, req *paymentv1.RefundPaymentRequest) (*paymentv1.Response, error) {
	fmt.Println("Stripe.Refund request:", req)

	response := paymentv1.Response{
		Msg: "Stripe: Refund Made",
	}
	return &response, nil
}

func NewStripeClient() *Stripe {
	return &Stripe{}
}
