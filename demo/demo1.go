package demo

import (
	"encoding/json"
	"fmt"
	"saastack/core"
	"saastack/core/types"
	"saastack/interfaces/email"
	emailtypes "saastack/interfaces/email/types"
	"saastack/interfaces/payment"
	paymenttypes "saastack/interfaces/payment/types"
)

func SendEmailViaAWSSES() types.ResponseData {
	emailData := emailtypes.EmailInterfaceData{
		From: "abc@def.ghi",
		To:   "jkl@mno.pqr",
		Body: "stuvwxyz",
	}
	data, _ := json.Marshal(emailData)

	request := types.RequestData{
		InterfaceType: types.EmailInterfaceType,
		Params: types.InterfaceRequestData{
			PluginId: email.AWSSES,
			Data:     data,
			Route:    emailtypes.SendMailRoute,
		},
	}

	response := core.RunInterface(request)

	return response
}

func SendEmailAndNotification() types.ResponseData {
	paymentData := paymenttypes.PaymentInterfaceData{
		Amount:   100,
		ClientId: "client123",
	}
	data, _ := json.Marshal(paymentData)

	request := types.RequestData{
		InterfaceType: types.PaymentInterfaceType,
		Params: types.InterfaceRequestData{
			PluginId: payment.RAZORPAY,
			Data:     data,
			Route:    paymenttypes.MakePaymentRoute,
		},
	}

	response := core.RunInterface(request)
	fmt.Println("Payment Request response:", response)

	emailData := emailtypes.EmailInterfaceData{
		From: "abc@def.ghi",
		To:   "jkl@mno.pqr",
		Body: "stuvwxyz",
	}
	data, _ = json.Marshal(emailData)

	request = types.RequestData{
		InterfaceType: types.EmailInterfaceType,
		Params: types.InterfaceRequestData{
			PluginId: email.MAILGUN,
			Data:     data,
			Route:    emailtypes.SendMailRoute,
		},
	}

	response = core.RunInterface(request)
	fmt.Println("Notification Send via MailGun")

	return response
}
