package service

import (
	"context"
	"fmt"
	"log"
	corev1 "saastack/gen/core/v1"
	"saastack/interfaces"
	"saastack/interfaces/email/types"
	"saastack/plugins/email"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var configPath string = "interfaces/email/plugins.yaml"

var pluginMap map[interfaces.PluginID]types.PluginMapData = make(map[interfaces.PluginID]types.PluginMapData)

func init() {
	config := interfaces.ParsePluginYaml(configPath)

	for _, plugin := range config.Plugins {
		if plugin.Deployment == string(interfaces.MICROSERVICE) {
			log.Println("Plugin deploy via microservice", plugin)
		} else {
			log.Println("Plugin deploy via monolithic", plugin)
			// AWS Client
			awsSESClient := email.NewAmazonSES()
			// Mailgun Client
			mailgunClient := email.NewMailGun()

			switch plugin.Name {
			case string(email.AWSSES_ID):
				pluginMap[email.AWSSES_ID] = types.PluginMapData{
					Plugin: plugin,
					Client: awsSESClient,
				}

			case string(email.MAILGUN_ID):
				pluginMap[email.MAILGUN_ID] = types.PluginMapData{
					Plugin: plugin,
					Client: mailgunClient,
				}
			default:
				log.Println("plugin is invalid", plugin)
			}
		}
	}
}

type EmailService struct {
	corev1.UnimplementedEmailServiceServer
}

func (email *EmailService) SendEmail(_ context.Context, req *corev1.SendEmailRequest) (*corev1.Response, error) {
	fmt.Println("Email Service Req: ", req)

	plugin, ok := pluginMap[interfaces.PluginID(req.PluginId)]
	if !ok {
		return nil, status.Errorf(codes.Unimplemented, "invalid plugin id")
	}

	client := plugin.Client
	response, err := client.SendEmail(req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	return response, nil
}
