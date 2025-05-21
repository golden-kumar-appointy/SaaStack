package payment

import (
	"fmt"
	"SaaStack/core"
	"SaaStack/interfaces/payment"
)

type Stripe struct {
}

func NewStripe() Stripe {
	return Stripe{}
}
func (p Stripe) MakePayment() {
	fmt.Println("Razorpay: Payment Made")
}
