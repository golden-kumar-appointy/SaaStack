package types

import corev1 "saastack/gen/core/v1"

const (
	RAZORPAY = "razorpay"
	STRIPE   = "stripe"
)

type PaymentPlugin interface {
	Charge(req *corev1.ChargePaymentRequest_ChargeData) (*corev1.Response, error)
	Refund(req *corev1.RefundPaymentRequest_RefundData) (*corev1.Response, error)
}
