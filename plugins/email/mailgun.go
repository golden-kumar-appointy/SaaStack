package email

import (
	"context"
	"fmt"
	"saastack/interfaces"
	service "saastack/interfaces/email"
	emailv1 "saastack/interfaces/email/proto/gen/v1"
)

const MAILGUN_ID interfaces.PluginID = "mailgun"

type MailGun struct{}

func (provider *MailGun) SendEmail(_ context.Context, req *emailv1.SendEmailRequest) (*emailv1.Response, error) {
	fmt.Println("MailGun.sendEmail request:", req)

	response := emailv1.Response{
		Msg: "Mailgun: sent Email",
	}
	return &response, nil
}

func NewMailGun() service.EmailPlugin {
	return &MailGun{}
}
