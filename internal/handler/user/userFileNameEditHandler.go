package user

import (
	"minio-server/internal/logic/user"
	"minio-server/internal/svc"
	"minio-server/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 修改文件名称的接口
func UserFileNameEditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFileNameEditRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserFileNameEditLogic(r.Context(), svcCtx)
		resp, err := l.UserFileNameEdit(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
