package types

import (
	emailv1 "saastack/gen/email/v1"
	"saastack/interfaces"
)

type EmailPlugin interface {
	SendEmail(req *emailv1.SendEmailRequest_SendEmailData) (*emailv1.Response, error)
}

type PluginMapData struct {
	Plugin interfaces.PluginData
	Client EmailPlugin
}
