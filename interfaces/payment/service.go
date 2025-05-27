package service

import (
	"context"
	"log"
	paymentv1 "saastack/gen/payment/v1"
	"saastack/interfaces"
	"saastack/interfaces/payment/types"
	"saastack/plugins/payment"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var configPath string = "interfaces/payment/plugins.yaml"

var pluginMap map[interfaces.PluginID]types.PluginMapData = make(map[interfaces.PluginID]types.PluginMapData)

func init() {
	config := interfaces.ParsePluginYaml(configPath)

	for _, plugin := range config.Plugins {
		if plugin.Deployment == string(interfaces.MICROSERVICE) {
			log.Println("Plugin deploy via microservice", plugin)
			pluginMap[interfaces.PluginID(plugin.Name)] = types.PluginMapData{
				Plugin: plugin,
			}
		} else {
			log.Println("Plugin deploy via monolithic", plugin)
			// Razor Pay Client
			razorpayClient := payment.NewRazorPayClient()
			// Stripe Client
			stripeClient := payment.NewStripeClient()

			switch plugin.Name {
			case string(payment.RAZORPAY_ID):
				pluginMap[interfaces.PluginID(plugin.Name)] = types.PluginMapData{
					Plugin: plugin,
					Client: razorpayClient,
				}

			case string(payment.STRIPE_ID):
				pluginMap[interfaces.PluginID(plugin.Name)] = types.PluginMapData{
					Plugin: plugin,
					Client: stripeClient,
				}
			default:
				log.Println("plugin is invalid", plugin)
			}
		}
	}
}

type PaymentService struct {
	paymentv1.UnimplementedPaymentServiceServer
}

func (payment *PaymentService) Charge(_ context.Context, req *paymentv1.ChargePaymentRequest) (*paymentv1.Response, error) {
	log.Println("Payment Charge req :", req)

	plugin, ok := pluginMap[interfaces.PluginID(req.PluginId)]
	if !ok {
		return nil, status.Errorf(codes.Unimplemented, "invalid plugin id")
	}

	var response *paymentv1.Response

	if plugin.Plugin.Deployment == string(interfaces.MONOLITHIC) {
		client := plugin.Client
		result, err := client.Charge(req.Data)
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

		client := paymentv1.NewPaymentServiceClient(conn)
		data := paymentv1.ChargePaymentRequest{
			PluginId: plugin.Plugin.Name,
			Data:     req.Data,
		}
		result, err := client.Charge(context.Background(), &data)
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

	plugin, ok := pluginMap[interfaces.PluginID(req.PluginId)]
	if !ok {
		return nil, status.Errorf(codes.Unimplemented, "invalid plugin id")
	}

	var response *paymentv1.Response

	if plugin.Plugin.Deployment == string(interfaces.MONOLITHIC) {
		client := plugin.Client
		result, err := client.Refund(req.Data)
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

		client := paymentv1.NewPaymentServiceClient(conn)
		data := paymentv1.RefundPaymentRequest{
			PluginId: plugin.Plugin.Name,
			Data:     req.Data,
		}
		result, err := client.Refund(context.Background(), &data)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		response = result
	}

	return response, nil
}
