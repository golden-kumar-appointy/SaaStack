package types

import (
	corev1 "saastack/gen/core/v1"
	"saastack/interfaces"
)

type PaymentPlugin interface {
	Charge(req *corev1.ChargePaymentRequest_ChargeData) (*corev1.Response, error)
	Refund(req *corev1.RefundPaymentRequest_RefundData) (*corev1.Response, error)
}

type PluginMapData struct {
	Plugin interfaces.PluginData
	Client PaymentPlugin
}
