package service

import (
	"context"
	emailv1 "saastack/gen/email/v1"
	"saastack/interfaces"
)

type EmailPlugin interface {
	SendEmail(context.Context, *emailv1.SendEmailRequest) (*emailv1.Response, error)
}

type PluginMapData struct {
	Plugin interfaces.PluginData
	Client EmailPlugin
}
