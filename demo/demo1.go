package demo

import (
	"encoding/json"
	"fmt"
	"saastack/core"
	"saastack/core/types"
	"saastack/interfaces/datatypes"
	"saastack/interfaces/email"
	"saastack/interfaces/payment"
)

func SendEmailViaAWSSES() types.ResponseData {
	emailData := datatypes.EmailInterfaceData{
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
		},
	}

	response := core.RunInterface(request)

	return response
}

func SendEmailAndNotification() types.ResponseData {
	paymentData := datatypes.PaymentInterfaceData{
		Amount:   100,
		ClientId: "client123",
	}
	data, _ := json.Marshal(paymentData)

	request := types.RequestData{
		InterfaceType: types.PaymentInterfaceType,
		Params: types.InterfaceRequestData{
			PluginId: payment.RAZORPAY,
			Data:     data,
		},
	}

	response := core.RunInterface(request)
	fmt.Println("Payment Request response:", response)

	emailData := datatypes.EmailInterfaceData{
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
		},
	}

	response = core.RunInterface(request)
	fmt.Println("Notification Send via MailGun")

	return response
}
