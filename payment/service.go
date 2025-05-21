package payment

type Payment interface {
	MakePayment(amount int) error
}

type NewPaymentClient struct {
	Id string
}

func PaymentClient(id string) *NewPaymentClient {
	return &NewPaymentClient{
		Id: id,
	}
}
