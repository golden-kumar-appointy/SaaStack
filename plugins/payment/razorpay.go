package payment

import "fmt"

type Razorpay struct {
}

func NewRazorpay() Razorpay {
	return Razorpay{}
}
func (p Razorpay) MakePayment() {
	fmt.Println("Razorpay: Payment Made")
}
