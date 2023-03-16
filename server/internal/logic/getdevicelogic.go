package logic

import (
	"context"

	"iot/server/internal/svc"
	"iot/server/internal/types"

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
	device := l.svcCtx.Manager.GetDevice(req.UUID)
	return &types.GetDeviceResp{Device: device}, nil
}
