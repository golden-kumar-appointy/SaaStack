package email

import (
	"saastack/core/types"
	"saastack/plugins"
	"saastack/plugins/email"
)

const (
	AWSSES             = "awsses"
	MAILGUN            = "mailgun"
	UNIMPLEMENTEDEMAIL = "unimplementedEmail"
)

func NewEmailInterfaceHandler(request types.InterfaceRequestData) *types.InterfaceHandler {
	var client types.InterfaceHandler

	switch request.PluginId {
	case AWSSES:
		client = email.NewAmazonSES()

	case MAILGUN:
		client = email.NewMailGun()

	default:
		client = plugins.NewUnimplementedEmail()
	}

	return &client
}
