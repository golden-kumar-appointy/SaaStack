package email

import (
	"fmt"
	"SaaStack/core"
	"SaaStack/interfaces/email"
)

type MailGun struct{}

func NewMailGun() MailGun {
	return MailGun{}
}

func (p MailGun) SendEmail() {
	fmt.Println("MailGun: sent EMail")
}
