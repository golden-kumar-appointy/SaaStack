package service

import (
	"context"
	"fmt"
	"log"
	emailv1 "saastack/gen/email/v1"
	"saastack/interfaces"
	"saastack/interfaces/email/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var configPath string = "interfaces/email/plugins.yaml"

var pluginMap map[interfaces.PluginID]types.PluginMapData = make(map[interfaces.PluginID]types.PluginMapData)

func RegisterNewEmailPlugin(pluginData types.PluginMapData) {
	pluginMap[interfaces.PluginID(pluginData.Plugin.Name)] = pluginData

	log.Println("Added Plugin to Email interface", pluginData.Plugin.Name)
}

//	func init() {
//		config := interfaces.ParsePluginYaml(configPath)
//
//		for _, plugin := range config.Plugins {
//			if plugin.Deployment == string(interfaces.MICROSERVICE) {
//				log.Println("Plugin deploy via microservice", plugin)
//				pluginMap[interfaces.PluginID(plugin.Name)] = types.PluginMapData{
//					Plugin: plugin,
//				}
//			} else {
//				log.Println("Plugin deploy via monolithic", plugin)
//				// AWS Client
//				awsSESClient := email.NewAmazonSES()
//				// Mailgun Client
//				mailgunClient := email.NewMailGun()
//
//				switch plugin.Name {
//				case string(email.AWSSES_ID):
//					pluginMap[interfaces.PluginID(plugin.Name)] = types.PluginMapData{
//						Plugin: plugin,
//						Client: awsSESClient,
//					}
//
//				case string(email.MAILGUN_ID):
//					pluginMap[interfaces.PluginID(plugin.Name)] = types.PluginMapData{
//						Plugin: plugin,
//						Client: mailgunClient,
//					}
//				default:
//					log.Println("plugin is invalid", plugin)
//				}
//			}
//		}
//
//		log.Println("Email PluginMap:", pluginMap)
//	}

type EmailService struct {
	emailv1.UnimplementedEmailServiceServer
}

func (email *EmailService) SendEmail(_ context.Context, req *emailv1.SendEmailRequest) (*emailv1.Response, error) {
	fmt.Println("Email Service Req: ", req)

	plugin, ok := pluginMap[interfaces.PluginID(req.PluginId)]
	if !ok {
		return nil, status.Errorf(codes.Unimplemented, "invalid plugin id")
	}

	var response *emailv1.Response
	reqData := emailv1.SendEmailRequest{
		PluginId: plugin.Plugin.Name,
		Data:     req.Data,
	}

	if plugin.Plugin.Deployment == string(interfaces.MONOLITHIC) {
		client := plugin.Client
		result, err := client.SendEmail(context.Background(), &reqData)
		response = result
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Server Error")
		}
	} else if plugin.Plugin.Deployment == string(interfaces.MICROSERVICE) {
		log.Println("microservice called")
		conn, err := grpc.NewClient(plugin.Plugin.Source, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		client := emailv1.NewEmailServiceClient(conn)
		result, err := client.SendEmail(context.Background(), &reqData)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		response = result
	}

	return response, nil
}
