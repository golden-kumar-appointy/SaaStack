package service

import (
	"saastack/interfaces"
	emailv1 "saastack/interfaces/email/proto/gen/v1"
)

type EmailPlugin = emailv1.EmailServiceServer

type PluginMapData struct {
	Plugin interfaces.PluginData
	Client EmailPlugin
}
