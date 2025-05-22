package types

import (
	"encoding/json"
	"saastack/core/types"
)

const (
	SendMailRoute = "sendMail"
)

type EmailInterfaceHandler interface {
	types.InterfaceHandler
	SendEmail(request EmailInterfaceData) types.ResponseData
}

type EmailInterfaceData struct {
	From string
	To   string
	Body string
}

func (emailData *EmailInterfaceData) Parse(data []byte) {
	json.Unmarshal(data, emailData)
}
