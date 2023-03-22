package handler

import (
	"net/http"

	"minio-server/internal/logic"
	"minio-server/internal/svc"
	"minio-server/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 注册码
func RegisterCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCodeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterCodeLogic(r.Context(), svcCtx)
		resp, err := l.RegisterCode(&req, r.URL.Query().Get("email"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
