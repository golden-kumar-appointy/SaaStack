package types

import (
	corev1 "saastack/gen/core/v1"
	"saastack/interfaces"
)

type EmailPlugin interface {
	SendEmail(req *corev1.SendEmailRequest_SendEmailData) (*corev1.Response, error)
}

type PluginMapData struct {
	Plugin interfaces.PluginData
	Client EmailPlugin
}
