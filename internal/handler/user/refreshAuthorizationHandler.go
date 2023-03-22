package user

import (
	"minio-server/internal/logic/user"
	"minio-server/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 刷新授权管理
func RefreshAuthorizationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewRefreshAuthorizationLogic(r.Context(), svcCtx)
		resp, err := l.RefreshAuthorization(r.Header.Get("Authorization"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
