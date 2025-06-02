package payment

import (
	"context"
	"fmt"
	"saastack/interfaces"
	emailservice "saastack/interfaces/email"
	emailv1 "saastack/interfaces/email/proto/gen/v1"
	service "saastack/interfaces/payment"
	paymentv1 "saastack/interfaces/payment/proto/gen/v1"
)

const RAZORPAY_ID interfaces.PluginID = "razorpay"

type Razorpay struct {
	paymentv1.UnimplementedPaymentServiceServer
}

func (provider *Razorpay) Charge(_ context.Context, req *paymentv1.ChargePaymentRequest) (*paymentv1.Response, error) {
	fmt.Println("Razorpay.Charge request:", req)
	res1, err := emailservice.PluginMap[interfaces.PluginID(emailservice.DefaultPlugin)].Client.SendEmail(context.Background(), &emailv1.SendEmailRequest{
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
		Msg: "Razorpay: payment Made",
	}
	return &response, nil
}

func (provider *Razorpay) Refund(_ context.Context, req *paymentv1.RefundPaymentRequest) (*paymentv1.Response, error) {
	fmt.Println("Razorpay.Refund request:", req)
	res1, err := emailservice.PluginMap[interfaces.PluginID(emailservice.DefaultPlugin)].Client.SendEmail(context.Background(), &emailv1.SendEmailRequest{
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
		Msg: "Razorpay: Refund Made",
	}
	return &response, nil
}

func NewRazorPayClient() service.PaymentPlugin {
	return &Razorpay{}
}
