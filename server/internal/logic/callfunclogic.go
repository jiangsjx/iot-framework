package logic

import (
	"context"
	"fmt"
	"time"

	"iot/server/internal/svc"
	"iot/server/internal/types"

	"github.com/go-resty/resty/v2"
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
	device := l.svcCtx.Manager.GetDevice(req.UUID)
	if device == nil {
		return &types.CallFuncResp{BaseResp: types.BaseResp{
			Code:    0,
			Message: "Success",
			Data:    nil,
		}}, nil
	}

	url := fmt.Sprintf("http://%s/api/v1/func/%s", device.Endpoint, req.Func)
	_, err = resty.New().SetTimeout(3 * time.Second).R().Post(url)
	if err != nil {
		return &types.CallFuncResp{BaseResp: types.BaseResp{
			Code:    0,
			Message: "Error",
			Data:    nil,
		}}, err
	}

	return &types.CallFuncResp{BaseResp: types.BaseResp{
		Code:    0,
		Message: "Success",
		Data:    nil,
	}}, nil
}
