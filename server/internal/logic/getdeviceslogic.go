package logic

import (
	"context"

	"iot/server/internal/svc"
	"iot/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDevicesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDevicesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDevicesLogic {
	return &GetDevicesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDevicesLogic) GetDevices(req *types.GetDevicesReq) (resp *types.GetDevicesResp, err error) {
	devices := l.svcCtx.Manager.GetDevices()
	return &types.GetDevicesResp{Devices: devices}, nil
}
