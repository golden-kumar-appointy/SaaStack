package email

import (
	"fmt"
	corev1 "saastack/gen/core/v1"
)

type AmazonSES struct{}

func (provider *AmazonSES) SendEmail(req *corev1.SendEmailRequest_SendEmailData) (*corev1.Response, error) {
	fmt.Println("AmazonSES.sendEmail request:", req)

	response := corev1.Response{
		Msg: "AmazonSES: sent Email",
	}
	return &response, nil
}

func NewAmazonSES() *AmazonSES {
	return &AmazonSES{}
}
