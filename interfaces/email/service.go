package service

import (
	"context"
	"fmt"
	"log"
	"saastack/core"
	"saastack/interfaces"
	emailpb "saastack/interfaces/email/proto/gen/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var (
	PluginMap     map[interfaces.PluginID]PluginMapData = make(map[interfaces.PluginID]PluginMapData)
	DefaultPlugin string
)

type Service struct {
	emailpb.UnimplementedEmailServiceServer
}

func (service *Service) SendEmail(_ context.Context, req *emailpb.SendEmailRequest) (*emailpb.Response, error) {
	fmt.Println("Email Service Req: ", req)

	if len(req.PluginId) == 0 {
		req.PluginId = DefaultPlugin
	}

	plugin, ok := PluginMap[interfaces.PluginID(req.PluginId)]
	if !ok {
		return nil, status.Errorf(codes.Unimplemented, "invalid plugin id")
	}

	reqData := emailpb.SendEmailRequest{
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

func NewService() *Service {
	return &Service{}
}

func RegisterNewPlugin(pluginData PluginMapData) {
	PluginMap[interfaces.PluginID(pluginData.Plugin.Name)] = pluginData

	log.Println("Added Plugin to Email interface", pluginData.Plugin.Name)
}

func RegisterDefaultPlugin(name string) {
	DefaultPlugin = name
}

func RegisterGrpcHandler(srv *grpc.Server) {
	emailpb.RegisterEmailServiceServer(srv, NewService())
}

func RegisterHTTPHandler(srv *grpc.Server, mux *runtime.ServeMux, ctx context.Context) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := emailpb.RegisterEmailServiceHandlerFromEndpoint(ctx, mux, core.CORE_ADDRESS, opts); err != nil {
		panic(err)
	}
}
