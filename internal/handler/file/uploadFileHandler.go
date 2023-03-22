package file

import (
	"crypto/md5"
	"fmt"
	"minio-server/internal/logic/file"
	"minio-server/models"
	"minio-server/result"
	"minio-server/utils"
	"net/http"
	"path"

	"minio-server/internal/svc"
	"minio-server/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadFileRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		files, header, err := r.FormFile("file")
		if err != nil {
			return
		}
		// 计算hash
		bytes := make([]byte, header.Size)
		_, err = files.Read(bytes)
		if err != nil {
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(bytes))
		// 定义对象
		// 查询是否存在hash值，如果存在直接返回
		get, err := models.RepositoryPool{}.GetHashRepositoryPool(hash, svcCtx.engine)
		if err != nil {
			return
		}
		if get.Id > 0 {
			m := make(map[string]string)
			m["identity"] = get.identity
			m["ext"] = get.ext
			m["name"] = get.name
			m["type"] = get.Type
			httpx.OkJson(w, &types.UploadFileResponse{
				Result: result.OK("", m),
			})
			return
		}
		// 不存在进行存储
		upload, err := utils.UploadFileToMinio(r)
		if err != nil {
			return
		}
		req.Name = header.Filename
		req.Ext = path.Ext(header.Filename)
		req.Size = header.Size
		req.Hash = hash
		req.Path = upload
		l := file.NewUploadFileLogic(r.Context(), svcCtx)
		resq, err := l.UploadFile(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
