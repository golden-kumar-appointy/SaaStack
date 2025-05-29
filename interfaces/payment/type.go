package service

import (
	"context"
	paymentv1 "saastack/gen/payment/v1"
	"saastack/interfaces"
)

type PaymentPlugin interface {
	Charge(context.Context, *paymentv1.ChargePaymentRequest) (*paymentv1.Response, error)
	Refund(context.Context, *paymentv1.RefundPaymentRequest) (*paymentv1.Response, error)
}

type PluginMapData struct {
	Plugin interfaces.PluginData
	Client PaymentPlugin
}
