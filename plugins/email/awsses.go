package email

import (
	"fmt"
	"saastack/core/types"
	"saastack/interfaces/datatypes"
)

type AmazonSES struct{}

func (p *AmazonSES) Run(request types.InterfaceRequestData) types.ResponseData {
	var data datatypes.EmailInterfaceData
	data.Parse(request.Data)

	fmt.Println("PluginId :", request.PluginId)
	fmt.Println("Data :", data)

	response := types.ResponseData{
		Msg: "AmazonSES: sent Email",
	}
	return response
}

func NewAmazonSES() *AmazonSES {
	return &AmazonSES{}
}
