package email

import (
	"SaaStack/core"
)

type EmailInterface interface {
	SendEmail()
}

func NewEmailInterface() EmailInterface {

}
