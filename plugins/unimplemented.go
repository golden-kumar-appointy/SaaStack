package plugins

import (
	"fmt"
	"saastack/core/types"
)

type UnimplementedPlugin struct{}

func (u *UnimplementedPlugin) Run(request types.InterfaceRequestData) types.ResponseData {
	fmt.Println("Data :", request)
	response := types.ResponseData{
		Msg: "Plugin not implemented",
	}
	return response
}

func NewUnimplementedEmail() *UnimplementedPlugin {
	return &UnimplementedPlugin{}
}
