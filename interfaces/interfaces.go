package interfaces

import (
	"fmt"
	"saastack/core/types"
)

type UnimplementedInterfaceHandler struct{}

func (u *UnimplementedInterfaceHandler) Run(request types.InterfaceRequestData) types.ResponseData {
	fmt.Println("PluginId :", request.PluginId)
	fmt.Println("Data :", string(request.Data))

	response := types.ResponseData{
		Msg: "This Interface handler not implemented : " + request.PluginId,
	}
	return response
}

func NewUnimplementedInterfaceHandlerClient(request types.InterfaceRequestData) *UnimplementedInterfaceHandler {
	return &UnimplementedInterfaceHandler{}
}
