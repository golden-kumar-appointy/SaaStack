package service

import (
	"context"
	"fmt"
	"log"
	emailv1 "saastack/gen/email/v1"
	"saastack/interfaces"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var (
	pluginMap     map[interfaces.PluginID]PluginMapData = make(map[interfaces.PluginID]PluginMapData)
	defaultPlugin string
)

func RegisterNewEmailPlugin(pluginData PluginMapData) {
	pluginMap[interfaces.PluginID(pluginData.Plugin.Name)] = pluginData

	log.Println("Added Plugin to Email interface", pluginData.Plugin.Name)
}

func RegisterDefaultPlugin(name string) {
	defaultPlugin = name
}

type EmailService struct {
	emailv1.UnimplementedEmailServiceServer
}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (email *EmailService) SendEmail(_ context.Context, req *emailv1.SendEmailRequest) (*emailv1.Response, error) {
	fmt.Println("Email Service Req: ", req)

	if len(req.PluginId) == 0 {
		req.PluginId = defaultPlugin
	}

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
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Server Error")
		}
		response = result
	} else if plugin.Plugin.Deployment == string(interfaces.MICROSERVICE) {
		log.Println("microservice email called")
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
