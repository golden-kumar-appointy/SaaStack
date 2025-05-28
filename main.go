package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"saastack/core"
	emailv1 "saastack/gen/email/v1"
	paymentv1 "saastack/gen/payment/v1"
	"saastack/interfaces"
	emailService "saastack/interfaces/email"
	emailType "saastack/interfaces/email/types"
	paymentService "saastack/interfaces/payment"
	paymentType "saastack/interfaces/payment/types"
	"saastack/plugins/email"
	"saastack/plugins/payment"

	"gopkg.in/yaml.v2"
)

var Services map[string]bool = make(map[string]bool)

func init() {
	src := "interface.yaml"
	currDir, _ := filepath.Abs(".")
	src = path.Join(currDir, src)

	res := ParsePluginYaml(src)

	for _, val := range res.Services {
		Services[val] = true
	}

	for _, config := range res.Plugins {
		_, ok := Services[config.Interface]
		if !ok {
			log.Println("Service not found")
			continue
		}

		switch config.Interface {
		case "email":
			RegisterEmailPlugin(config)
		case "payment":
			RegisterPaymentPlugin(config)
		default:
			log.Println("Interface not implemented", config.Interface)
		}
	}

	log.Println(res)
}

func main() {
	srv := core.NewGrpcServer()

	// Register Service to core
	for key := range Services {
		switch key {
		case "email":
			emailv1.RegisterEmailServiceServer(srv, &emailService.EmailService{})
		case "payment":
			paymentv1.RegisterPaymentServiceServer(srv, &paymentService.PaymentService{})
		default:
			log.Println("Interface Handler Not Implemented")
		}
	}

	if err := core.StartServer(srv); err != nil {
		panic(err)
	}
}

type PluginConfig struct {
	Plugins  []interfaces.PluginData `yaml:"plugins"`
	Services []string                `yaml:"services"`
}

func ParsePluginYaml(src string) *PluginConfig {
	data, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}

	result := PluginConfig{}
	if err := yaml.Unmarshal(data, &result); err != nil {
		panic(err)
	}

	return &result
}

func RegisterEmailPlugin(config interfaces.PluginData) {
	var data emailType.PluginMapData

	if config.Deployment == string(interfaces.MICROSERVICE) {
		data = emailType.PluginMapData{
			Plugin: interfaces.PluginData{
				Name:       config.Name,
				Deployment: config.Deployment,
				Source:     config.Source,
			},
		}
	} else if config.Deployment == string(interfaces.MONOLITHIC) {
		var client emailType.EmailPlugin

		switch config.Name {
		case "mailgun":
			client = email.NewMailGun()
		case "awsses":
			client = email.NewAmazonSES()
		default:
			log.Println("Plugin Instance Not Implemented")
			return
		}
		data = emailType.PluginMapData{
			Plugin: interfaces.PluginData{
				Name:       config.Name,
				Deployment: config.Deployment,
				Source:     config.Source,
			},
			Client: client,
		}
	}
	emailService.RegisterNewEmailPlugin(data)
}

func RegisterPaymentPlugin(config interfaces.PluginData) {
	var data paymentType.PluginMapData

	if config.Deployment == string(interfaces.MICROSERVICE) {
		data = paymentType.PluginMapData{
			Plugin: interfaces.PluginData{
				Name:       config.Name,
				Deployment: config.Deployment,
				Source:     config.Source,
			},
		}
	} else if config.Deployment == string(interfaces.MONOLITHIC) {
		var client paymentType.PaymentPlugin

		switch config.Name {
		case "stripe":
			client = payment.NewStripeClient()
		case "razorpay":
			client = payment.NewRazorPayClient()
		default:
			log.Println("Plugin Instance Not Implemented")
			return
		}
		data = paymentType.PluginMapData{
			Plugin: interfaces.PluginData{
				Name:       config.Name,
				Deployment: config.Deployment,
				Source:     config.Source,
			},
			Client: client,
		}
	}
	paymentService.RegisterNewPaymentPlugin(data)
}
