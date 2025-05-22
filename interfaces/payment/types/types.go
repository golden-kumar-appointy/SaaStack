package payment

import (
	"encoding/json"
	"saastack/core/types"
)

const (
	MakePaymentRoute = "makePayment"
)

type PaymentInterfaceHandler interface {
	types.InterfaceHandler
	MakePayment(request PaymentInterfaceData) types.ResponseData
}

type PaymentInterfaceData struct {
	Amount   int
	ClientId string
}

func (paymentData *PaymentInterfaceData) Parse(data []byte) {
	json.Unmarshal(data, paymentData)
}
