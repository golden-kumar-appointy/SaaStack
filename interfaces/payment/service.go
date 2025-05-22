package payment

import (
	"saastack/core/types"
	paymenttypes "saastack/interfaces/payment/types"
	"saastack/plugins"
	"saastack/plugins/payment"
)

const (
	RAZORPAY             = "razorpay"
	STRIPE               = "stripe"
	UNIMPLEMENTEDPAYMENT = "unimplementedPayment"
)

func NewPaymentInterfaceHandler(request types.InterfaceRequestData) types.InterfaceHandler {
	var client paymenttypes.PaymentInterfaceHandler

	switch request.PluginId {
	case RAZORPAY:
		client = payment.NewRazorPayClient()

	case STRIPE:
		client = payment.NewStripeClient()

	default:
		client = plugins.NewUnimplementedPlugin()
	}

	return client
}
