package service

import (
	"context"
	"log"
	"saastack/interfaces"
	paymentv1 "saastack/interfaces/payment/proto/gen/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var (
	pluginMap     map[interfaces.PluginID]PluginMapData = make(map[interfaces.PluginID]PluginMapData)
	defaultPlugin string
)

func RegisterNewPaymentPlugin(pluginData PluginMapData) {
	pluginMap[interfaces.PluginID(pluginData.Plugin.Name)] = pluginData

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

	plugin, ok := pluginMap[interfaces.PluginID(req.PluginId)]
	if !ok {
		return nil, status.Errorf(codes.Unimplemented, "invalid plugin id")
	}

	var response *paymentv1.Response
	reqData := paymentv1.ChargePaymentRequest{
		PluginId: plugin.Plugin.Name,
		Data:     req.Data,
	}

	if plugin.Plugin.Deployment == string(interfaces.MONOLITHIC) {
		client := plugin.Client
		result, err := client.Charge(context.Background(), &reqData)
		response = result
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Server Error")
		}

	} else if plugin.Plugin.Deployment == string(interfaces.MICROSERVICE) {
		log.Println("microservice payment called")

		conn, err := grpc.NewClient(plugin.Plugin.Source, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		client := paymentv1.NewPaymentServiceClient(conn)
		result, err := client.Charge(context.Background(), &reqData)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		response = result
	}
	return response, nil
}

func (payment *PaymentService) Refund(_ context.Context, req *paymentv1.RefundPaymentRequest) (*paymentv1.Response, error) {
	log.Println("Payment Refund req :", req)

	if len(req.PluginId) == 0 {
		req.PluginId = defaultPlugin
	}

	plugin, ok := pluginMap[interfaces.PluginID(req.PluginId)]
	if !ok {
		return nil, status.Errorf(codes.Unimplemented, "invalid plugin id")
	}

	var response *paymentv1.Response
	reqData := paymentv1.RefundPaymentRequest{
		PluginId: plugin.Plugin.Name,
		Data:     req.Data,
	}

	if plugin.Plugin.Deployment == string(interfaces.MONOLITHIC) {
		client := plugin.Client
		result, err := client.Refund(context.Background(), &reqData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Server Error")
		}
		response = result
	} else if plugin.Plugin.Deployment == string(interfaces.MICROSERVICE) {
		log.Println("microservice payment called")
		conn, err := grpc.NewClient(plugin.Plugin.Source, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		client := paymentv1.NewPaymentServiceClient(conn)

		result, err := client.Refund(context.Background(), &reqData)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		response = result
	}

	return response, nil
}
