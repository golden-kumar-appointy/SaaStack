package email

import (
	"fmt"
	emailv1 "saastack/gen/email/v1"
	"saastack/interfaces"
)

type AmazonSES struct{}

const AWSSES_ID interfaces.PluginID = "awsses"

func (provider *AmazonSES) SendEmail(req *emailv1.SendEmailRequest_SendEmailData) (*emailv1.Response, error) {
	fmt.Println("AmazonSES.sendEmail request:", req)

	response := emailv1.Response{
		Msg: "AmazonSES: sent Email",
	}
	return &response, nil
}

func NewAmazonSES() *AmazonSES {
	return &AmazonSES{}
}
