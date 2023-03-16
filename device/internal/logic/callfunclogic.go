package logic

import (
	"context"

	"iot/device/internal/svc"
	"iot/device/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallFuncLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCallFuncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallFuncLogic {
	return &CallFuncLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CallFuncLogic) CallFunc(req *types.CallFuncReq) (resp *types.CallFuncResp, err error) {
	l.Infof("call func: %s", req.Func)

	return &types.CallFuncResp{
		BaseResp: types.BaseResp{
			Code:    0,
			Message: "Success",
			Data:    nil,
		},
	}, nil
}
