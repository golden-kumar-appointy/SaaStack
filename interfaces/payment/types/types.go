package payment

import "encoding/json"

const (
	MakePaymentRoute = "makePayment"
)

type PaymentInterfaceData struct {
	Amount   int
	ClientId string
}

func (paymentData *PaymentInterfaceData) Parse(data []byte) {
	json.Unmarshal(data, paymentData)
}
