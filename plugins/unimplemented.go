package plugins

import (
	"fmt"
	"saastack/core/types"
)

type UnimplementedPlugin struct{}

func (u *UnimplementedPlugin) Run(request types.InterfaceRequestData) types.ResponseData {
	fmt.Println("PluginId :", request.PluginId)
	fmt.Println("Data :", string(request.Data))

	response := types.ResponseData{
		Msg: "Plugin not implemented",
	}
	return response
}

func NewUnimplementedEmail() *UnimplementedPlugin {
	return &UnimplementedPlugin{}
}
