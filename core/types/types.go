package types

const (
	EmailInterfaceType         = "email"
	PaymentInterfaceType       = "payment"
	UnimplementedInterfaceType = "unimplemented"
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
	Route    string
	Data     []byte
}

type InterfaceHandler interface {
	Run(request InterfaceRequestData) ResponseData
}
