package types

import (
	paymentv1 "saastack/gen/payment/v1"
	"saastack/interfaces"
)

type PaymentPlugin interface {
	Charge(req *paymentv1.ChargePaymentRequest_ChargeData) (*paymentv1.Response, error)
	Refund(req *paymentv1.RefundPaymentRequest_RefundData) (*paymentv1.Response, error)
}

type PluginMapData struct {
	Plugin interfaces.PluginData
	Client PaymentPlugin
}
