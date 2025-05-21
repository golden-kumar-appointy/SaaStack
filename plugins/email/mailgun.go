package email

import (
	"fmt"
	"saastack/core/types"
	"saastack/interfaces/datatypes"
)

type MailGun struct{}

func (p *MailGun) Run(request types.InterfaceRequestData) types.ResponseData {
	var data datatypes.EmailInterfaceData
	data.Parse(request.Data)

	fmt.Println("PluginId :", request.PluginId)
	fmt.Println("Data :", data)

	response := types.ResponseData{
		Msg: "MailGun: sent Email",
	}
	return response
}

func NewMailGun() *MailGun {
	return &MailGun{}
}
