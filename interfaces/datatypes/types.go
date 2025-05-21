package datatypes

import "encoding/json"

type EmailInterfaceData struct {
	From string
	To   string
	Body string
}

func (emailData *EmailInterfaceData) Parse(data []byte) {
	json.Unmarshal(data, emailData)
}

type PaymentInterfaceData struct {
	Amount   int
	ClientId string
}

func (paymentData *PaymentInterfaceData) Parse(data []byte) {
	json.Unmarshal(data, paymentData)
}
