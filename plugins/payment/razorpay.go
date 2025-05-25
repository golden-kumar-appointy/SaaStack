package payment

import (
	"fmt"
	corev1 "saastack/gen/core/v1"
)

type Razorpay struct{}

func (provider *Razorpay) Charge(req *corev1.ChargePaymentRequest_ChargeData) (*corev1.Response, error) {
	fmt.Println("Razorpay.Charge request:", req)

	response := corev1.Response{
		Msg: "Razorpay: payment Made",
	}
	return &response, nil
}

func (provider *Razorpay) Refund(req *corev1.RefundPaymentRequest_RefundData) (*corev1.Response, error) {
	fmt.Println("Razorpay.Refund request:", req)

	response := corev1.Response{
		Msg: "Razorpay: Refund Made",
	}
	return &response, nil
}

func NewRazorPayClient() *Razorpay {
	return &Razorpay{}
}
