package demo

import (
	"encoding/json"
	"saastack/core"
	"saastack/core/types"
	"saastack/interfaces/datatypes"
)

func UnimplementedInterfaceHandler() types.ResponseData {
	emailData := datatypes.EmailInterfaceData{
		From: "abc@def.ghi",
		To:   "jkl@mno.pqr",
		Body: "stuvwxyz",
	}
	data, _ := json.Marshal(emailData)

	request := types.RequestData{
		InterfaceType: types.UnimplementedInterfaceType,
		Params: types.InterfaceRequestData{
			PluginId: types.EmailInterfaceType,
			Data:     data,
		},
	}

	response := core.RunInterface(request)
	return response
}

func UnimplementedInterfacePlugin() types.ResponseData {
	emailData := datatypes.EmailInterfaceData{
		From: "abc@def.ghi",
		To:   "jkl@mno.pqr",
		Body: "stuvwxyz",
	}
	data, _ := json.Marshal(emailData)

	request := types.RequestData{
		InterfaceType: types.EmailInterfaceType,
		Params: types.InterfaceRequestData{
			PluginId: types.UnimplementedInterfaceType,
			Data:     data,
		},
	}

	response := core.RunInterface(request)
	return response
}
