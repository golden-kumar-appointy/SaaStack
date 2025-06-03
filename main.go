package main

import (
	"context"
	"log"
	"os"
	"path"
	"path/filepath"
	"saastack/core"
	"saastack/interfaces"
	"saastack/plugins/email"
	"saastack/plugins/payment"

	emailService "saastack/interfaces/email"
	paymentService "saastack/interfaces/payment"

	"gopkg.in/yaml.v3"
)

var Services = make(map[string]bool)

func main() {
	// Register plugin to services(interfaces)
	ReadConfigFile()

	// gRPC Server
	srv := core.NewGrpcServer()

	// HTTP Gateway
	mux := core.NewMuxServer()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register Service to core
	for key := range Services {
		switch key {
		case "email":
			emailService.RegisterGrpcHandler(srv)
			emailService.RegisterHTTPHandler(srv, mux, ctx)
		case "payment":
			paymentService.RegisterGrpcHandler(srv)
			paymentService.RegisterHTTPHandler(srv, mux, ctx)
		default:
			log.Println("Interface Handler Not Implemented", key)
		}
	}

	go core.StartHttpGateway(mux)

	if err := core.StartServer(srv); err != nil {
		panic(err)
	}
}

func ReadConfigFile() {
	src := "config.yaml"
	currDir, _ := filepath.Abs(".")
	src = path.Join(currDir, src)

	res := ParsePluginYaml(src)

	for _, val := range res.Services {
		Services[val] = false
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
			if !Services[config.Interface] {
				emailService.RegisterDefaultPlugin(config.Name)
				Services[config.Interface] = true
				log.Println("Default plugin for email: ", config.Name)
			}
		case "payment":
			RegisterPaymentPlugin(config)
			if !Services[config.Interface] {
				paymentService.RegisterDefaultPlugin(config.Name)
				Services[config.Interface] = true
				log.Println("Default plugin for payment: ", config.Name)
			}
		default:
			log.Println("Interface not implemented", config.Interface)
		}
	}
	log.Println(res)
}

func RegisterEmailPlugin(config interfaces.PluginData) {
	var data emailService.PluginMapData

	var client emailService.EmailPlugin

	switch config.Name {
	case "mailgun":
		client = email.NewMailGun()
	case "awsses":
		client = email.NewAmazonSES()
	default:
		log.Println("Plugin Instance Not Implemented")
		return
	}

	data = emailService.PluginMapData{
		Plugin: interfaces.PluginData{
			Name: config.Name,
		},
		Client: client,
	}

	emailService.RegisterNewPlugin(data)
}

func RegisterPaymentPlugin(config interfaces.PluginData) {
	var data paymentService.PluginMapData

	var client paymentService.PaymentPlugin

	switch config.Name {
	case "stripe":
		client = payment.NewStripeClient()
	case "razorpay":
		client = payment.NewRazorPayClient()
	default:
		log.Println("Plugin Instance Not Implemented")
		return
	}

	data = paymentService.PluginMapData{
		Plugin: interfaces.PluginData{
			Name: config.Name,
		},
		Client: client,
	}
	paymentService.RegisterNewPlugin(data)
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
