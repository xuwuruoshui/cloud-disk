package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepostitorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepostitorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepostitorySaveLogic {
	return &UserRepostitorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepostitorySaveLogic) UserRepostitorySave(req *types.UserRepositorySaveRequest, userIdentity string) (resp *types.UserRepositorySaveReply, err error) {
	ur := &models.UserRepository{
		Identity: helper.UUID(),
		UserIdentity: userIdentity,
		ParentId: req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext: req.Ext,
		Name: req.Name,
	}

	_, err = l.svcCtx.Engine.Insert(ur)
	if err!=nil{
		return nil,err
	}

	return &types.UserRepositorySaveReply{
		Identity:ur.Identity,
	},nil
}
