package logic

import (
	"cloud-disk/core/helper"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadChunkCompleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadChunkCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadChunkCompleteLogic {
	return &FileUploadChunkCompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadChunkCompleteLogic) FileUploadChunkComplete(req *types.FileUploadChunkCompleteRequest) (resp *types.FileUploadChunkCompleteReply, err error) {

	cosObjects := make([]cos.Object,0,len(req.CosObjects))
	for _,v := range req.CosObjects {
		cosObjects = append(cosObjects,cos.Object{
			PartNumber: v.PartNum,ETag: v.Etag,
		})
	}

	err = helper.CosPartUploadComplete(req.Key, req.UploadId, cosObjects)

	return
}
