package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"iot/server/internal/logic"
	"iot/server/internal/svc"
	"iot/server/internal/types"
)

func getDeviceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetDeviceReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetDeviceLogic(r.Context(), svcCtx)
		resp, err := l.GetDevice(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
