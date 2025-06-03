package service

import (
	"context"
	"log"
	"saastack/core"
	"saastack/interfaces"
	paymentpb "saastack/interfaces/payment/proto/gen/v1"

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
	paymentpb.UnimplementedPaymentServiceServer
}

func (service *Service) Charge(_ context.Context, req *paymentpb.ChargePaymentRequest) (*paymentpb.Response, error) {
	log.Println("Payment Charge req :", req)

	if len(req.PluginId) == 0 {
		req.PluginId = DefaultPlugin
	}

	plugin, ok := PluginMap[interfaces.PluginID(req.PluginId)]
	if !ok {
		return nil, status.Errorf(codes.Unimplemented, "invalid plugin id")
	}

	reqData := paymentpb.ChargePaymentRequest{
		PluginId: plugin.Plugin.Name,
		Data:     req.Data,
	}

	client := plugin.Client
	response, err := client.Charge(context.Background(), &reqData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	return response, nil
}

func (service *Service) Refund(_ context.Context, req *paymentpb.RefundPaymentRequest) (*paymentpb.Response, error) {
	log.Println("Payment Refund req :", req)

	if len(req.PluginId) == 0 {
		req.PluginId = DefaultPlugin
	}

	plugin, ok := PluginMap[interfaces.PluginID(req.PluginId)]
	if !ok {
		return nil, status.Errorf(codes.Unimplemented, "invalid plugin id")
	}

	reqData := paymentpb.RefundPaymentRequest{
		PluginId: plugin.Plugin.Name,
		Data:     req.Data,
	}

	client := plugin.Client
	response, err := client.Refund(context.Background(), &reqData)
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

	log.Println("Added Plugin to Payment interface", pluginData.Plugin.Name)
}

func RegisterDefaultPlugin(name string) {
	DefaultPlugin = name
}

func RegisterGrpcHandler(srv *grpc.Server) {
	paymentpb.RegisterPaymentServiceServer(srv, NewService())
}

func RegisterHTTPHandler(srv *grpc.Server, mux *runtime.ServeMux, ctx context.Context) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := paymentpb.RegisterPaymentServiceHandlerFromEndpoint(ctx, mux, core.CORE_ADDRESS, opts); err != nil {
		panic(err)
	}
}
