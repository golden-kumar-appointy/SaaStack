package payment

import (
	"fmt"
	corev1 "saastack/gen/core/v1"
)

type Stripe struct{}

func (provider *Stripe) Charge(req *corev1.ChargePaymentRequest_ChargeData) (*corev1.Response, error) {
	fmt.Println("Stripe.Charge request:", req)

	response := corev1.Response{
		Msg: "Stripe: payment Made",
	}
	return &response, nil
}

func (provider *Stripe) Refund(req *corev1.RefundPaymentRequest_RefundData) (*corev1.Response, error) {
	fmt.Println("Stripe.Refund request:", req)

	response := corev1.Response{
		Msg: "Stripe: Refund Made",
	}
	return &response, nil
}

func NewStripeClient() *Stripe {
	return &Stripe{}
}
