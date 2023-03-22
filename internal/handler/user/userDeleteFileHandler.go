package user

import (
	"net/http"

	"minio-server/internal/logic/user"
	"minio-server/internal/svc"
	"minio-server/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户删除文件管理
func UserDeleteFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserDeleteFileRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserDeleteFileLogic(r.Context(), svcCtx)
		req.Identity = r.URL.Query().Get("identity")
		resp, err := l.UserDeleteFile(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
