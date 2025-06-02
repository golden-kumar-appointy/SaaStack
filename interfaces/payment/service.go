package service

import (
	"context"
	"log"
	"saastack/interfaces"
	paymentv1 "saastack/interfaces/payment/proto/gen/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

)

var (
	PluginMap map[interfaces.PluginID]PluginMapData = make(map[interfaces.PluginID]PluginMapData)
	defaultPlugin string
)

func RegisterNewPaymentPlugin(pluginData PluginMapData) {
	PluginMap[interfaces.PluginID(pluginData.Plugin.Name)] = pluginData

	log.Println("Added Plugin to Payment interface", pluginData.Plugin.Name)
}

func RegisterDefaultPlugin(name string) {
	defaultPlugin = name
}

type PaymentService struct {
	paymentv1.UnimplementedPaymentServiceServer
}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (payment *PaymentService) Charge(_ context.Context, req *paymentv1.ChargePaymentRequest) (*paymentv1.Response, error) {
	log.Println("Payment Charge req :", req)

	if len(req.PluginId) == 0 {
		req.PluginId = defaultPlugin
	}

	plugin, ok := PluginMap[interfaces.PluginID(req.PluginId)]
	if !ok {
		return nil, status.Errorf(codes.Unimplemented, "invalid plugin id")
	}

	reqData := paymentv1.ChargePaymentRequest{
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

func (payment *PaymentService) Refund(_ context.Context, req *paymentv1.RefundPaymentRequest) (*paymentv1.Response, error) {
	log.Println("Payment Refund req :", req)

	if len(req.PluginId) == 0 {
		req.PluginId = defaultPlugin
	}
	
	plugin, ok := PluginMap[interfaces.PluginID(req.PluginId)]
	if !ok {
		return nil, status.Errorf(codes.Unimplemented, "invalid plugin id")
	}

	reqData := paymentv1.RefundPaymentRequest{
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
