// Code generated by goctl. DO NOT EDIT.
package types

type BaseDevice struct {
	UUID     string   `json:"uuid"`
	Name     string   `json:"name"`
	Endpoint string   `json:"endpoint"`
	Funcs    []string `json:"funcs"`
}

type BaseResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ScanReq struct {
}

type ScanResp struct {
	BaseResp
}

type GetDevicesReq struct {
}

type GetDevicesResp struct {
	Devices []*BaseDevice `json:"devices"`
}

type GetDeviceReq struct {
	UUID string `path:"uuid"`
}

type GetDeviceResp struct {
	Device *BaseDevice `json:"device"`
}

type CallFuncReq struct {
	UUID string `path:"uuid"`
	Func string `path:"func"`
}

type CallFuncResp struct {
	BaseResp
}
