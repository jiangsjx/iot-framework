package logic

import (
	"context"
	"time"

	"iot/server/internal/svc"
	"iot/server/internal/types"

	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type ScanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScanLogic {
	return &ScanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScanLogic) Scan(req *types.ScanReq) (resp *types.ScanResp, err error) {
	resp2, err := resty.New().SetTimeout(3 * time.Second).R().SetResult(&types.GetDeviceResp{}).
		Get("http://127.0.0.1:8889/api/v1/device")
	if err != nil {
		return &types.ScanResp{BaseResp: types.BaseResp{
			Code:    0,
			Message: "Error",
			Data:    nil,
		}}, err
	}

	device := resp2.Result().(*types.GetDeviceResp)
	l.Infof("get device: %v", device.Device)

	l.svcCtx.Manager.AddDevice(device.Device)

	return &types.ScanResp{BaseResp: types.BaseResp{
		Code:    0,
		Message: "Success",
		Data:    nil,
	}}, nil
}
