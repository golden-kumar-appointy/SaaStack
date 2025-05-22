package email

import (
	"fmt"
	"saastack/core/types"
	emailtypes "saastack/interfaces/email/types"
)

type MailGun struct{}

func (provider *MailGun) sendEmail(request emailtypes.EmailInterfaceData) types.ResponseData {
	fmt.Println("MailGun.sendEmail request:", request)

	response := types.ResponseData{
		Msg: "Mailgun: sent Email",
	}
	return response
}

func (p *MailGun) Run(request types.InterfaceRequestData) types.ResponseData {
	var data emailtypes.EmailInterfaceData
	data.Parse(request.Data)

	fmt.Println("PluginId :", request.PluginId)
	fmt.Println("Route :", request.Route)

	var response types.ResponseData

	switch request.Route {
	case emailtypes.SendMailRoute:
		response = p.sendEmail(data)

	default:
		response.Msg = "Route not present"
	}

	return response
}

func NewMailGun() *MailGun {
	return &MailGun{}
}
