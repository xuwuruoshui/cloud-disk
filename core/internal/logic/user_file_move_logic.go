package logic

import (
	"cloud-disk/core/models"
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveReply, err error) {

	// 1.判断是否存在该parent文件夹
	var up models.UserRepository
	found, err := l.svcCtx.Engine.Where("identity = ? and user_identity=?", req.ParentIdentity, userIdentity).Get(&up)
	if !found{
		return nil,errors.New("不存在该文件夹")
	}

	// 2.更新parentId

	_, err = l.svcCtx.Engine.Where("identity=?", req.Identity).Update(&models.UserRepository{
		ParentId: int64(up.Id),
	})
	return
}
