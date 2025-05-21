package email

import "fmt"

type MailGun struct{}

func NewMailGun() MailGun {
	return MailGun{}
}

func (p MailGun) SendEmail() {
	fmt.Println("MailGun: sent EMail")
}
