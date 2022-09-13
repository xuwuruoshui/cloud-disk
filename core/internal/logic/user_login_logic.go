package logic

import (
	"cloud-disk/core/define"
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	// 1. 从数据库中查询用户
	var user models.UserBasic
	found, err := l.svcCtx.Engine.Where("name=?", req.Name).Get(&user)
	if err!=nil{
		return nil,err
	}
	if !found{
		return nil,errors.New("不存在该用户！！！")
	}

	success := helper.Decrypt(req.Password, user.Password)
	if !success{
		return nil,errors.New("密码错误！！！")
	}

	// 2. 生成token
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name,define.TokenExpireTime)
	if err != nil {
		return nil,err
	}

	// 3.生成refresh_token
	refeshToken, err := helper.GenerateToken(user.Id, user.Identity, user.Name,define.RefreshTokenExpireTime)
	if err != nil {
		return nil,err
	}

	resp = new(types.LoginReply)
	resp.Token = token
	resp.RefreshToken = refeshToken
	return resp,nil
}
