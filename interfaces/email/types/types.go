package types

import (
	"encoding/json"
)

const (
	SendMailRoute = "sendMail"
)

type EmailInterfaceData struct {
	From string
	To   string
	Body string
}

func (emailData *EmailInterfaceData) Parse(data []byte) {
	json.Unmarshal(data, emailData)
}
