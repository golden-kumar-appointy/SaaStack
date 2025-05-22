package email

import (
	"fmt"
	"saastack/core/types"
	emailtypes "saastack/interfaces/email/types"
)

type AmazonSES struct{}

func (provider *AmazonSES) SendEmail(request emailtypes.EmailInterfaceData) types.ResponseData {
	fmt.Println("AmazonSES.sendEmail request:", request)

	response := types.ResponseData{
		Msg: "AmazonSES: sent Email",
	}
	return response
}

func (p *AmazonSES) Run(request types.InterfaceRequestData) types.ResponseData {
	var data emailtypes.EmailInterfaceData
	data.Parse(request.Data)

	fmt.Println("PluginId :", request.PluginId)
	fmt.Println("Route :", request.Route)

	var response types.ResponseData

	switch request.Route {
	case emailtypes.SendMailRoute:
		response = p.SendEmail(data)

	default:
		response.Msg = "Route not present"
	}

	return response
}

func NewAmazonSES() *AmazonSES {
	return &AmazonSES{}
}
