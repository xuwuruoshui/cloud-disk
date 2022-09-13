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

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateReply, err error) {

	var ur models.UserRepository
	found, err := l.svcCtx.Engine.Where("identity=?", req.UserRepositoryIdentity).Get(&ur)
	if err!=nil{
		return nil, err
	}
	if !found{
		return nil, errors.New("该仓库不存在")
	}

	sb := &models.ShareBasic{
		Identity: helper.UUID(),
		UserIdentity: userIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		RepositoryIdentity: ur.RepositoryIdentity,
		ExpireTime: req.ExpireTime,
	}

	_, err = l.svcCtx.Engine.InsertOne(sb)
	if err!=nil{
		return nil,err
	}
	resp = new(types.ShareBasicCreateReply)
	resp.Identity = sb.Identity
	return
}
