package payment

import (
	"fmt"
	"saastack/core/types"
)

type Razorpay struct{}

func (p *Razorpay) Run(request types.InterfaceRequestData) types.ResponseData {
	fmt.Println("Data :", request)
	response := types.ResponseData{
		Msg: "Razorpay: Payment Send",
	}

	return response
}

func NewRazorPayClient() *Razorpay {
	return &Razorpay{}
}
