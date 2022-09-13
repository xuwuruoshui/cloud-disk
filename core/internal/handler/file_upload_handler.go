package handler

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"cloud-disk/core/internal/logic"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		file, header, err := r.FormFile("file")
		if err!=nil{
			return
		}
		// 判断文件是否存在
		b := make([]byte,header.Size)
		_, err = file.Read(b)
		if err!=nil{
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))
		rp := new(models.RepositoryPool)
		has, err := svcCtx.Engine.Where("hash=?", hash).Get(rp)
		if err!=nil{
			return
		}
		if has{
			httpx.OkJson(w,&types.FileUploadReply{Identity: rp.Identity,Ext: rp.Ext,Name: rp.Name})
			return
		}

		// Cos上传
		cosPath, err := helper.CosUpload(r)
		if err!=nil{
			return
		}

		// 往logic中传递request
		req.Name = header.Filename
		req.Ext = path.Ext(header.Filename)
		req.Size = header.Size
		req.Hash = hash
		req.Path = cosPath

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
