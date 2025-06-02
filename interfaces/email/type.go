package service

import (
	"context"
	"saastack/interfaces"

	emailv1 "saastack/interfaces/email/proto/gen/v1"
)

type EmailPlugin interface {
	SendEmail(context.Context, *emailv1.SendEmailRequest) (*emailv1.Response, error)
}

type PluginMapData struct {
	Plugin interfaces.PluginData
	Client EmailPlugin
}
