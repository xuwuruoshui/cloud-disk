package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateReply, err error) {

	var up models.UserRepository
	exist, err := l.svcCtx.Engine.Where("name = ? and parent_id=?", req.Name, req.ParentId).Limit(1).Exist(&up)
	if err!=nil{
		return nil,err
	}
	if exist{
		return nil,errors.New("该名称已存在")
	}

	// 创建文件夹
	data := &models.UserRepository{
		Identity: helper.UUID(),
		UserIdentity: userIdentity,
		ParentId: req.ParentId,
		Name: req.Name,
	}
	_, err = l.svcCtx.Engine.Insert(data)

	return
}
