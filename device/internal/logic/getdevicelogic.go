package logic

import (
	"context"

	"iot/device/internal/svc"
	"iot/device/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeviceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDeviceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeviceLogic {
	return &GetDeviceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDeviceLogic) GetDevice(req *types.GetDeviceReq) (resp *types.GetDeviceResp, err error) {
	return &types.GetDeviceResp{
		Device: &types.BaseDevice{
			UUID:     "ac249bd5-49a1-4dd4-ba7d-7b3367318173",
			Name:     "ABC",
			Endpoint: "127.0.0.1:8889",
			Funcs:    []string{"open", "close"},
		}}, nil
}
