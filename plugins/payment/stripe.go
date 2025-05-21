package payment

import (
	"fmt"
	"saastack/core/types"
)

type Stripe struct{}

func (p *Stripe) Run(request types.InterfaceRequestData) types.ResponseData {
	fmt.Println("Data :", request)
	response := types.ResponseData{
		Msg: "Stripe: Payment Sent",
	}

	return response
}

func NewStripeClient() *Stripe {
	return &Stripe{}
}
