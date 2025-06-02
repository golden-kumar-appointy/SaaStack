package email

import (
	"context"
	"fmt"
	"saastack/interfaces"
	service "saastack/interfaces/email"
	emailv1 "saastack/interfaces/email/proto/gen/v1"
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

func NewAmazonSES() service.EmailPlugin {
	return &AmazonSES{}
}
