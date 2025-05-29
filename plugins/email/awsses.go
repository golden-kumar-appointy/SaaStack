package email

import (
	"context"
	"fmt"
	emailv1 "saastack/gen/email/v1"
	"saastack/interfaces"
)

const AWSSES_ID interfaces.PluginID = "awsses"

type AmazonSES struct{}

func (provider *AmazonSES) SendEmail(_ context.Context, req *emailv1.SendEmailRequest) (*emailv1.Response, error) {
	fmt.Println("AmazonSES.sendEmail request:", req)

	response := emailv1.Response{
		Msg: "AmazonSES: sent Email",
	}
	return &response, nil
}

func NewAmazonSES() *AmazonSES {
	return &AmazonSES{}
}
