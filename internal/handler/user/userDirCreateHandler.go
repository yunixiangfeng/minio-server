package user

import (
	"net/http"

	"minio-server/internal/logic/user"
	"minio-server/internal/svc"
	"minio-server/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 创建文件夹接口
func UserDirCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserDirCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserDirCreateLogic(r.Context(), svcCtx)
		resp, err := l.UserDirCreate(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
