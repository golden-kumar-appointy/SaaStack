package demo

import (
	"fmt"
	"saastack/core"
	"saastack/core/types"
	"saastack/interfaces/email"
	"saastack/interfaces/payment"
)

func SendEmailViaAWSSES() types.ResponseData {
	request := types.RequestData{
		InterfaceType: types.EmailInterfaceType,
		Params: types.InterfaceRequestData{
			PluginId: email.AWSSES,
			Data:     "Data to send email via AWSSES",
		},
	}

	response := core.RunInterface(request)

	return response
}

func SendEmailAndNotification() types.ResponseData {
	request := types.RequestData{
		InterfaceType: types.PaymentInterfaceType,
		Params: types.InterfaceRequestData{
			PluginId: payment.RAZORPAY,
			Data:     "payment made via razorpay",
		},
	}

	response := core.RunInterface(request)
	fmt.Println("Payment Request response:", response)

	request = types.RequestData{
		InterfaceType: types.EmailInterfaceType,
		Params: types.InterfaceRequestData{
			PluginId: email.MAILGUN,
			Data:     "Mail send via Mailgun",
		},
	}

	response = core.RunInterface(request)
	fmt.Println("Notification Send via MailGun")

	return response
}
