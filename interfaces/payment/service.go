package payment

import (
	"SaaStack/core"
)
type PaymentInterface interface {
	MakePayment()
}

func NewPaymentInterface() PaymentInterface {

}
