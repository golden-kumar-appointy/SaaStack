package core

import (
	"saastack/core/types"
	"saastack/interfaces"
	"saastack/interfaces/email"
	"saastack/interfaces/payment"
)

func RunInterface(request types.RequestData) types.ResponseData {
	var client types.InterfaceHandler

	switch request.InterfaceType {
	case types.EmailInterfaceType:
		client = *email.NewEmailInterfaceHandler(request.Params)

	case types.PaymentInterfaceType:
		client = *payment.NewPaymentInterfaceHandler(request.Params)

	default:
		client = interfaces.NewUnimplementedInterfaceHandlerClient(request.Params)
	}

	response := client.Run(request.Params)
	return response
}
