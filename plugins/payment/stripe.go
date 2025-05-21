package payment

import "fmt"

type Stripe struct {
}

func NewStripe() Stripe {
	return Stripe{}
}
func (p Stripe) MakePayment() {
	fmt.Println("Razorpay: Payment Made")
}
