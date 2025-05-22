package external

import (
	"saastack/core/types"
	"saastack/plugins"
)

const (
	PAYANDNOTIFY = "payAndNotify"
)

func NewPaymentInterfaceHandler(request types.InterfaceRequestData) *types.InterfaceHandler {
	var client types.InterfaceHandler

	switch request.PluginId {
	default:
		client = plugins.NewUnimplementedPlugin()
	}

	return &client
}
