package main

import (
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
)

func RegisterEmailPlugin() {
	mailGun := email.NewMailGun()
	emailData := emailType.PluginMapData{
		Plugin: interfaces.PluginData{
			Name:       "mailgun",
			Deployment: string(interfaces.MONOLITHIC),
		},
		Client: mailGun,
	}
	emailService.RegisterNewEmailPlugin(emailData)
	awsses := email.NewAmazonSES()
	emailData = emailType.PluginMapData{
		Plugin: interfaces.PluginData{
			Name:       "awsses",
			Deployment: string(interfaces.MONOLITHIC),
		},
		Client: awsses,
	}
	emailService.RegisterNewEmailPlugin(emailData)

	emailData = emailType.PluginMapData{
		Plugin: interfaces.PluginData{
			Name:       "custom",
			Deployment: string(interfaces.MICROSERVICE),
			Source:     "localhost:9002",
		},
	}
	emailService.RegisterNewEmailPlugin(emailData)
}

func RegisterPaymentPlugin() {
	stripe := payment.NewStripeClient()
	payData := paymentType.PluginMapData{
		Plugin: interfaces.PluginData{
			Name:       "stripe",
			Deployment: string(interfaces.MONOLITHIC),
		},
		Client: stripe,
	}
	paymentService.RegisterNewPaymentPlugin(payData)

	razorpay := payment.NewRazorPayClient()
	payData = paymentType.PluginMapData{
		Plugin: interfaces.PluginData{
			Name:       "razorpay",
			Deployment: string(interfaces.MONOLITHIC),
		},
		Client: razorpay,
	}
	paymentService.RegisterNewPaymentPlugin(payData)

	payData = paymentType.PluginMapData{
		Plugin: interfaces.PluginData{
			Name:       "custom",
			Deployment: string(interfaces.MICROSERVICE),
			Source:     "localhost:9003",
		},
	}
	paymentService.RegisterNewPaymentPlugin(payData)
}

func main() {
	srv := core.NewGrpcServer()

	RegisterPaymentPlugin()
	RegisterEmailPlugin()

	emailv1.RegisterEmailServiceServer(srv, &emailService.EmailService{})
	paymentv1.RegisterPaymentServiceServer(srv, &paymentService.PaymentService{})

	if err := core.StartServer(srv); err != nil {
		panic(err)
	}
}
