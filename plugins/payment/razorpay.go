package payment

import (
	"fmt"
	"saastack/core/types"
	"saastack/interfaces/datatypes"
)

type Razorpay struct{}

func (p *Razorpay) Run(request types.InterfaceRequestData) types.ResponseData {
	var data datatypes.PaymentInterfaceData
	data.Parse(request.Data)

	fmt.Println("PluginId :", request.PluginId)
	fmt.Println("Data :", data)

	response := types.ResponseData{
		Msg: "Razorpay: Payment Send",
	}

	return response
}

func NewRazorPayClient() *Razorpay {
	return &Razorpay{}
}
