package interfaces

import (
	"fmt"
	"saastack/core/types"
)

type UnimplementedInterfaceHandler struct{}

func (u *UnimplementedInterfaceHandler) Run(request types.InterfaceRequestData) types.ResponseData {
	fmt.Println("Data :", request)
	response := types.ResponseData{
		Msg: "This Interface handler not implemented" + request.PluginId,
	}
	return response
}

func NewUnimplementedInterfaceHandlerClient() *UnimplementedInterfaceHandler {
	return &UnimplementedInterfaceHandler{}
}
