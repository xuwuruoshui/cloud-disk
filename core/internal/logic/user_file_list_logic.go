package logic

import (
	"cloud-disk/core/define"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListReply, err error) {


	resp = new(types.UserFileListReply)

	// 分页
	size:= req.Size
	if size==0{
		size = define.PageSize
	}
	page := req.Page
	if page==0{
		page=1
	}
	offset := (page-1)*size

	uf := make([]*types.UserFile,0)
	err = l.svcCtx.Engine.Table("user_repository").
		Select("user_repository.id,user_repository.identity,user_repository.repository_identity,user_repository.ext,user_repository.name,repository_pool.path,repository_pool.size").
		Join("left","repository_pool","user_repository.repository_identity = repository_pool.identity").
		Where("parent_id=? and user_identity=? and user_repository.delete_at=? or user_repository.delete_at is null", req.Id, userIdentity,time.Time{}.Format(define.Datetime)).
		Limit(size,offset).
		Find(&uf)
	if err!=nil{
		return nil,err
	}

	// 查询用户文件总数
	cnt, err := l.svcCtx.Engine.Table("user_repository").
		Where("parent_id=? and user_identity=?", req.Id, userIdentity).Count(&models.UserRepository{})

	if err!=nil{
		return nil,err
	}

	resp.List = uf
	resp.Count = cnt

	return
}
