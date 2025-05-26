package service

import (
	"context"
	"log"
	corev1 "saastack/gen/core/v1"
	"saastack/interfaces"
	"saastack/interfaces/payment/types"
	"saastack/plugins/payment"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var configPath string = "interfaces/payment/plugins.yaml"

var pluginMap map[interfaces.PluginID]types.PluginMapData = make(map[interfaces.PluginID]types.PluginMapData)

func init() {
	config := interfaces.ParsePluginYaml(configPath)

	for _, plugin := range config.Plugins {
		if plugin.Deployment == string(interfaces.MICROSERVICE) {
			log.Println("Plugin deploy via microservice", plugin)
		} else {
			log.Println("Plugin deploy via monolithic", plugin)
			// Razor Pay Client
			razorpayClient := payment.NewRazorPayClient()
			// Stripe Client
			stripeClient := payment.NewStripeClient()

			switch plugin.Name {
			case string(payment.RAZORPAY_ID):
				pluginMap[payment.RAZORPAY_ID] = types.PluginMapData{
					Plugin: plugin,
					Client: razorpayClient,
				}

			case string(payment.STRIPE_ID):
				pluginMap[payment.STRIPE_ID] = types.PluginMapData{
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
	corev1.UnimplementedPaymentServiceServer
}

func (payment *PaymentService) Charge(_ context.Context, req *corev1.ChargePaymentRequest) (*corev1.Response, error) {
	log.Println("Payment Charge req :", req)

	plugin, ok := pluginMap[interfaces.PluginID(req.PluginId)]
	if !ok {
		return nil, status.Errorf(codes.Unimplemented, "invalid plugin id")
	}

	client := plugin.Client
	response, err := client.Charge(req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	return response, nil
}

func (payment *PaymentService) Refund(_ context.Context, req *corev1.RefundPaymentRequest) (*corev1.Response, error) {
	log.Println("Payment Refund req :", req)

	plugin, ok := pluginMap[interfaces.PluginID(req.PluginId)]
	if !ok {
		return nil, status.Errorf(codes.Unimplemented, "invalid plugin id")
	}

	client := plugin.Client
	response, err := client.Refund(req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	return response, nil
}
