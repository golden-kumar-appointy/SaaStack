package service

import (
	"saastack/interfaces"

	paymentv1 "saastack/interfaces/payment/proto/gen/v1"
)

type PaymentPlugin = paymentv1.PaymentServiceServer

type PluginMapData struct {
	Plugin interfaces.PluginData
	Client PaymentPlugin
}
