package payment

import (
	"context"
	"fmt"
	"log"
	"saastack/core"
	"saastack/interfaces"
	emailv1 "saastack/interfaces/email/proto/gen/v1"
	service "saastack/interfaces/payment"
	paymentv1 "saastack/interfaces/payment/proto/gen/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const STRIPE_ID interfaces.PluginID = "stripe"

type Stripe struct{}

func (provider *Stripe) Charge(_ context.Context, req *paymentv1.ChargePaymentRequest) (*paymentv1.Response, error) {
	log.Println("Stripe.Charge request:", req)

	conn, err := grpc.NewClient(core.CORE_ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer conn.Close()

	client := emailv1.NewEmailServiceClient(conn)
	reqData := &emailv1.SendEmailRequest{
		Data: &emailv1.SendEmailRequest_SendEmailData{
			From: "paymentCharge@auenkr.com",
			To:   req.Data.ClientId,
			Body: "Amount charge: " + string(req.Data.Amount),
		},
	}

	result, err := client.SendEmail(context.Background(), reqData)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Notification charge call: ", result)

	response := paymentv1.Response{
		Msg: "Stripe: payment Made",
	}
	return &response, nil
}

func (provider *Stripe) Refund(_ context.Context, req *paymentv1.RefundPaymentRequest) (*paymentv1.Response, error) {
	fmt.Println("Stripe.Refund request:", req)

	conn, err := grpc.NewClient(core.CORE_ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer conn.Close()

	client := emailv1.NewEmailServiceClient(conn)
	reqData := &emailv1.SendEmailRequest{
		Data: &emailv1.SendEmailRequest_SendEmailData{
			From: "paymentCharge@auenkr.com",
			To:   req.Data.PaymentId,
			Body: "Amount charge: " + string(req.Data.Amount),
		},
	}

	result, err := client.SendEmail(context.Background(), reqData)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Notification for Refund call: ", result)

	response := paymentv1.Response{
		Msg: "Stripe: Refund Made",
	}
	return &response, nil
}

func NewStripeClient() service.PaymentPlugin {
	return &Stripe{}
}
