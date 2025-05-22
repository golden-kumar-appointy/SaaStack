package payment

import (
	"fmt"
	"saastack/core/types"
	paymenttypes "saastack/interfaces/payment/types"
)

type Razorpay struct{}

func (provider *Razorpay) MakePayment(request paymenttypes.PaymentInterfaceData) types.ResponseData {
	fmt.Println("Razorpay.MakePayment request:", request)

	response := types.ResponseData{
		Msg: "Razorpay: payment Made",
	}
	return response
}

func (p *Razorpay) Run(request types.InterfaceRequestData) types.ResponseData {
	var data paymenttypes.PaymentInterfaceData
	data.Parse(request.Data)

	fmt.Println("PluginId :", request.PluginId)
	fmt.Println("Route :", request.Route)

	var response types.ResponseData

	switch request.Route {
	case paymenttypes.MakePaymentRoute:
		response = p.MakePayment(data)

	default:
		response.Msg = "Route not present"
	}

	return response
}

func NewRazorPayClient() *Razorpay {
	return &Razorpay{}
}
