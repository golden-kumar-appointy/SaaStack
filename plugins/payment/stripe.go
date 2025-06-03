package payment

import (
	"context"
	"fmt"
	"log"
	"saastack/interfaces"
	emailservice "saastack/interfaces/email"
	emailv1 "saastack/interfaces/email/proto/gen/v1"
	service "saastack/interfaces/payment"
	paymentv1 "saastack/interfaces/payment/proto/gen/v1"
)

const STRIPE_ID interfaces.PluginID = "stripe"

type Stripe struct {
	paymentv1.UnimplementedPaymentServiceServer
}

func (provider *Stripe) Charge(_ context.Context, req *paymentv1.ChargePaymentRequest) (*paymentv1.Response, error) {
	log.Println("Stripe.Charge request:", req)

	client := emailservice.PluginMap[interfaces.PluginID(emailservice.DefaultPlugin)].Client
	res1, err := client.SendEmail(context.Background(), &emailv1.SendEmailRequest{
		Data: &emailv1.SendEmailRequest_SendEmailData{
			From: "razorpay@payment.com",
			To:   "test@test.com",
			Body: "This is a test email from Razorpay payment",
		},
	})
	fmt.Println("Notification response:", res1)
	if err != nil {
		return nil, err
	}

	response := paymentv1.Response{
		Msg: "Stripe: payment Made",
	}
	return &response, nil
}

func (provider *Stripe) Refund(_ context.Context, req *paymentv1.RefundPaymentRequest) (*paymentv1.Response, error) {
	fmt.Println("Stripe.Refund request:", req)

	client := emailservice.PluginMap[interfaces.PluginID(emailservice.DefaultPlugin)].Client
	res1, err := client.SendEmail(context.Background(), &emailv1.SendEmailRequest{
		Data: &emailv1.SendEmailRequest_SendEmailData{
			From: "razorpay@payment.com",
			To:   "test@test.com",
			Body: "This is a test email from Razorpay payment",
		},
	})
	fmt.Println("Notification response:", res1)
	if err != nil {
		return nil, err
	}

	response := paymentv1.Response{
		Msg: "Stripe: Refund Made",
	}
	return &response, nil
}

func NewStripeClient() service.PaymentPlugin {
	return &Stripe{}
}
