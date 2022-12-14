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

type ShareBasicSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicSaveLogic {
	return &ShareBasicSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicSaveLogic) ShareBasicSave(req *types.ShareBasicSaveRequest, userIdentity string) (resp *types.ShareBasicSaveReply, err error) {

	// 1.获取资源的详情
	var rp models.RepositoryPool
	found, err := l.svcCtx.Engine.Where("identity=?", req.RepositoryIdentity).Get(&rp)
	if err!=nil{
		return nil, err
	}
	if !found{
		return nil, errors.New("资源不存在")
	}

	// 2.user_repository保存
	ur := &models.UserRepository{
		Identity: helper.UUID(),
		UserIdentity: userIdentity,
		ParentId: req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext: rp.Ext,
		Name: rp.Name,
	}
	_, err = l.svcCtx.Engine.InsertOne(ur)
	if err!=nil{
		return nil,err
	}
	resp.Identity = ur.Identity

	return
}
