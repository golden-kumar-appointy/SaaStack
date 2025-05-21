package email

import (
	"fmt"
	"saastack/core/types"
)

type AmazonSES struct{}

func (p *AmazonSES) Run(request types.InterfaceRequestData) types.ResponseData {
	fmt.Println("Data :", request)
	response := types.ResponseData{
		Msg: "AmazonSES: sent Email",
	}
	return response
}

func NewAmazonSES() *AmazonSES {
	return &AmazonSES{}
}
