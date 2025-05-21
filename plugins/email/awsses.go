package email

import "fmt"

type AmazonSES struct{}

func NewAmazonSES() AmazonSES {
	return AmazonSES{}
}

func (p AmazonSES) SendEmail() {
	fmt.Println("AmazonSES: sent EMail")
}
