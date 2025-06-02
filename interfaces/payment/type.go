package service

import (
	"context"
	"saastack/interfaces"

	paymentv1 "saastack/interfaces/payment/proto/gen/v1"
)

type PaymentPlugin interface {
	Charge(context.Context, *paymentv1.ChargePaymentRequest) (*paymentv1.Response, error)
	Refund(context.Context, *paymentv1.RefundPaymentRequest) (*paymentv1.Response, error)
}

type PluginMapData struct {
	Plugin interfaces.PluginData
	Client PaymentPlugin
}
