package email

import (
	"fmt"
	emailv1 "saastack/gen/email/v1"
	"saastack/interfaces"
)

type MailGun struct{}

const MAILGUN_ID interfaces.PluginID = "mailgun"

func (provider *MailGun) SendEmail(req *emailv1.SendEmailRequest_SendEmailData) (*emailv1.Response, error) {
	fmt.Println("MailGun.sendEmail request:", req)

	response := emailv1.Response{
		Msg: "Mailgun: sent Email",
	}
	return &response, nil
}

func NewMailGun() *MailGun {
	return &MailGun{}
}
