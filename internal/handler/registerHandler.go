package handler

import (
	"net/http"

	"minio-sercer/internal/types"
	"minio-server/internal/logic"
	"minio-server/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 注册方法
func RegisterHandler(svcCtx *svc.ServiceContext) http.Handlerfunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resq, err := l.Register(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
