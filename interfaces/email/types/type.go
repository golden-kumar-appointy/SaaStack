package types

import (
	corev1 "saastack/gen/core/v1"
	"saastack/interfaces"
)

const (
	AWSSES  interfaces.PluginID = "awsses"
	MAILGUN interfaces.PluginID = "mailgun"
)

type EmailPlugin interface {
	SendEmail(req *corev1.SendEmailRequest_SendEmailData) (*corev1.Response, error)
}
