package email

import (
	"fmt"
	corev1 "saastack/gen/core/v1"
)

type MailGun struct{}

func (provider *MailGun) SendEmail(req *corev1.SendEmailRequest_SendEmailData) (*corev1.Response, error) {
	fmt.Println("MailGun.sendEmail request:", req)

	response := corev1.Response{
		Msg: "Mailgun: sent Email",
	}
	return &response, nil
}

func NewMailGun() *MailGun {
	return &MailGun{}
}
