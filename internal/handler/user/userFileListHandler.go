package user

import (
	"minio-server/internal/logic/user"
	"minio-server/internal/svc"
	"minio-server/internal/types"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户文件列表管理
func UserFileListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFileListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserFileListLogic(r.Context(), svcCtx)
		resp, err := l.UserFileList(&req, r.Header.Get("UserIdentity"))
		// 获取参数
		query := r.URL.Query()
		id, _ := strconv.Atoi(query.Get("id"))
		req.Id = int64(id)
		page, _ := strconv.Atoi(query.Get("page"))
		req.Page = page
		size, _ := strconv.Atoi(query.Get("size"))
		req.Size = size
		req.Type = query.Get("type")
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
