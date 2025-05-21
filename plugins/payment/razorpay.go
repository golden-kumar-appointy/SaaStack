package payment

import (
	"fmt"
	"SaaStack/core"
	"SaaStack/interfaces/payment"
)

type Razorpay struct {
}

func NewRazorpay() Razorpay {
	return Razorpay{}
}
func (p Razorpay) MakePayment() {
	fmt.Println("Razorpay: Payment Made")
}
