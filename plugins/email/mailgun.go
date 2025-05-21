package email

import (
	"fmt"
	"saastack/core/types"
)

type MailGun struct{}

func (p *MailGun) Run(request types.InterfaceRequestData) types.ResponseData {
	fmt.Println("Data :", request)
	response := types.ResponseData{
		Msg: "MailGun: sent Email",
	}
	return response
}

func NewMailGun() *MailGun {
	return &MailGun{}
}
