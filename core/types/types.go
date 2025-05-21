package types

const (
	EmailInterfaceType   = "email"
	PaymentInterfaceType = "payment"
)

type InterfaceType string

type RequestData struct {
	InterfaceType InterfaceType
	Params        InterfaceRequestData
}

type ResponseData struct {
	Msg string
}

type InterfaceRequestData struct {
	PluginId string
	Data     any
}

type InterfaceHandler interface {
	Run(InterfaceRequestData) ResponseData
}
