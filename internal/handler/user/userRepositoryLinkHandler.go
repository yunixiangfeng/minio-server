package user

import (
	"minio-server/internal/logic/user"
	"minio-server/internal/svc"
	"minio-server/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户池连接管理
func UserRepositoryLinkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRepositoryLinkRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserRepositoryLinkLogic(r.Context(), svcCtx)
		resp, err := l.UerRepositoryLink(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
