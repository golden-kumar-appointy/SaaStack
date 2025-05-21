package payment

import (
	"fmt"
	"saastack/core/types"
	"saastack/interfaces/datatypes"
)

type Stripe struct{}

func (p *Stripe) Run(request types.InterfaceRequestData) types.ResponseData {
	var data datatypes.PaymentInterfaceData
	data.Parse(request.Data)

	fmt.Println("PluginId :", request.PluginId)
	fmt.Println("Data :", data)

	response := types.ResponseData{
		Msg: "Stripe: Payment Sent",
	}

	return response
}

func NewStripeClient() *Stripe {
	return &Stripe{}
}
