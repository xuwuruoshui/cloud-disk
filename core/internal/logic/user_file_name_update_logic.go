package logic

import (
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateReply, err error) {

	// 判断当前名称在该层级下是否存在
	var pid int64
	_, err = l.svcCtx.Engine.
		Table("user_repository").
		Select("parent_id").
		Where("identity = ?", req.Identity).
		Limit(1).Get(&pid)
	if err!=nil{
		return nil,err
	}

	// 存在则直接返回
	var up models.UserRepository
	exist, err := l.svcCtx.Engine.Where("name = ? and parent_id=?", req.Name, pid).Limit(1).Exist(&up)
	if err!=nil{
		return nil,err
	}
	if exist{
		return nil,errors.New("该名称已存在")
	}

	// 文件名称修改
	data := &models.UserRepository{
		Name: req.Name,
	}
	_, err = l.svcCtx.Engine.Where("identity = ? and user_identity=?", req.Identity, userIdentity).Update(data)


	return
}
