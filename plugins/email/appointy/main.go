package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	emailv1 "saastack/gen/email/v1"
	"saastack/interfaces"

	"google.golang.org/grpc"
)

type AppointyEmail struct {
	emailv1.UnimplementedEmailServiceServer
}

const (
	CUSTOM_ID           interfaces.PluginID = "appointy"
	PLUGIN_ADDRESS      string              = "localhost:9005"
	NOTIFICATION_SERVER string              = "http://localhost:3001"
)

type EmailDataType struct {
	SenderId   string `json:"senderId"`
	RecieverId string `json:"recieverId"`
	Subject    string `json:"subject"`
}

type payloadType struct {
	Type       string            `json:"type"`
	TemplateId int               `json:"templateId"`
	Params     map[string]string `json:"params"`
	Data       EmailDataType     `json:"data"`
}

func sendToEmailService(from string, to string) error {
	source, _ := url.Parse(NOTIFICATION_SERVER)
	source.Path = "api/v1/send"

	name := "Golden Kumar"

	emailData := EmailDataType{
		SenderId:   from,
		RecieverId: to,
		Subject:    "SaaStack Demo Final Test",
	}

	paylaodData := payloadType{
		Type:       "email",
		TemplateId: 20,
		Params: map[string]string{
			"name": name,
		},
		Data: emailData,
	}

	data, err := json.Marshal(paylaodData)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, source.String(), bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiaWF0IjoxNzQ1OTk1MzY2LCJleHAiOjE3NTQ2MzUzNjZ9.W1Am1qPJ_UePkhzNc6sw0HoITwaunQqIsBMiVSLc7Kk")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	log.Println(string(body))
	return nil
}

func (provider *AppointyEmail) SendEmail(_ context.Context, req *emailv1.SendEmailRequest) (*emailv1.Response, error) {
	log.Println("Appointy.sendEmail request:", req)

	sendToEmailService(req.Data.From, req.Data.To)

	response := emailv1.Response{
		Msg: "Appointy: sent Email",
	}
	return &response, nil
}

func NewAppointyEmail() *AppointyEmail {
	return &AppointyEmail{}
}

func main() {
	lis, err := net.Listen("tcp", PLUGIN_ADDRESS)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	emailv1.RegisterEmailServiceServer(grpcServer, &AppointyEmail{})

	log.Printf("custom email plugin server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
