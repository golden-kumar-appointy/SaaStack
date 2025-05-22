package plugins

import (
	"fmt"
	"saastack/core/types"
	emailtypes "saastack/interfaces/email/types"
	paymenttypes "saastack/interfaces/payment/types"
)

type UnimplementedPlugin struct{}

func (provider *UnimplementedPlugin) SendEmail(request emailtypes.EmailInterfaceData) types.ResponseData {
	fmt.Println("UnimplementedPlugin.SendEmail request:", request)

	response := types.ResponseData{
		Msg: "UnimplementedPlugin: Plugin not implemented",
	}
	return response
}

func (provider *UnimplementedPlugin) MakePayment(request paymenttypes.PaymentInterfaceData) types.ResponseData {
	fmt.Println("UnimplementedPlugin.MakePayment request:", request)

	response := types.ResponseData{
		Msg: "UnimplementedPlugin: Plugin not implemented",
	}
	return response
}

func (u *UnimplementedPlugin) Run(request types.InterfaceRequestData) types.ResponseData {
	fmt.Println("PluginId :", request.PluginId)
	fmt.Println("Route :", request.Route)

	var response types.ResponseData

	switch request.Route {
	case emailtypes.SendMailRoute:
		var data emailtypes.EmailInterfaceData
		data.Parse(request.Data)
		response = u.SendEmail(data)

	case paymenttypes.MakePaymentRoute:
		var data paymenttypes.PaymentInterfaceData
		data.Parse(request.Data)
		response = u.MakePayment(data)

	default:
		response.Msg = "Route not present"
	}

	return response
}

func NewUnimplementedPlugin() *UnimplementedPlugin {
	return &UnimplementedPlugin{}
}
