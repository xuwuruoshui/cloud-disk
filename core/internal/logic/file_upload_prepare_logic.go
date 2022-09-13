package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrepareRequest) (resp *types.FileUploadPrepareReply, err error) {

	var rp models.RepositoryPool
	found, err := l.svcCtx.Engine.Where("hash = ?", req.Md5).Get(&rp)
	if err!=nil{
		return
	}

	resp = new(types.FileUploadPrepareReply)
	if found{
		// 秒传
		resp.Identity = rp.Identity

	}else{
		// TODO 获取文件的UploadId、Key,进行文件的分片上传
		// 1.初始化
		key, uploadId, err := helper.CosInitPartUpload(req.Ext)
		if err!=nil{
			return nil,err
		}
		resp.Key = key
		resp.UploadId = uploadId
	}

	return
}
