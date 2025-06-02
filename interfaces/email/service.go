package service

import (
	"context"
	"fmt"
	"log"
	"saastack/interfaces"
	emailv1 "saastack/interfaces/email/proto/gen/v1"

	"google.golang.org/grpc/codes"
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

	reqData := emailv1.SendEmailRequest{
		PluginId: plugin.Plugin.Name,
		Data:     req.Data,
	}

	client := plugin.Client
	response, err := client.SendEmail(context.Background(), &reqData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	return response, nil
}
