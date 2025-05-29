package email

import (
	"context"
	"fmt"
	emailv1 "saastack/gen/email/v1"
	"saastack/interfaces"
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

func NewMailGun() *MailGun {
	return &MailGun{}
}
