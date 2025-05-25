package service

import (
	"context"
	"fmt"
	corev1 "saastack/gen/core/v1"
	"saastack/interfaces"
	"saastack/interfaces/email/types"
	"saastack/plugins/email"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var pluginClient map[interfaces.PluginID]types.EmailPlugin = make(map[interfaces.PluginID]types.EmailPlugin)

func init() {
	// AWS Client
	awsSESClient := email.NewAmazonSES()
	pluginClient[types.AWSSES] = awsSESClient

	// Mailgun Client
	mailgunClient := email.NewMailGun()
	pluginClient[types.MAILGUN] = mailgunClient
}

type EmailService struct {
	corev1.UnimplementedEmailServiceServer
}

func (email *EmailService) SendEmail(_ context.Context, req *corev1.SendEmailRequest) (*corev1.Response, error) {
	fmt.Println("Email Service Req: ", req)

	client, ok := pluginClient[interfaces.PluginID(req.PluginId)]
	if !ok {
		return nil, status.Errorf(codes.Unimplemented, "invalid plugin id")
	}

	response, err := client.SendEmail(req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	return response, nil
}
