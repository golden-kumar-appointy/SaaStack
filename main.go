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

func main() {
	srv := core.NewGrpcServer()

	mailGun := email.NewMailGun()
	emailService.RegisterNewEmailPlugin(emailType.PluginMapData{
		Plugin: interfaces.PluginData{
			Name:       "mailgun",
			Deployment: string(interfaces.MONOLITHIC),
		},
		Client: mailGun,
	})

	stripe := payment.NewStripeClient()
	paymentService.RegisterNewPaymentPlugin(paymentType.PluginMapData{
		Plugin: interfaces.PluginData{
			Name:       "stripe",
			Deployment: string(interfaces.MONOLITHIC),
		},
		Client: stripe,
	})

	emailv1.RegisterEmailServiceServer(srv, &emailService.EmailService{})
	paymentv1.RegisterPaymentServiceServer(srv, &paymentService.PaymentService{})

	if err := core.StartServer(srv); err != nil {
		panic(err)
	}
}
