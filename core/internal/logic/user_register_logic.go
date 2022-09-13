package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {

	session := l.svcCtx.Engine.NewSession()
	defer session.Close()

	// add Begin() before any action
	if err := session.Begin(); err != nil {
		return nil,err
	}

	// 1.判断验证码是否正确
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err!=nil{
		return nil,errors.New("该邮箱验证码不存在")
	}

	if code!=req.Code{
		err = errors.New("验证码错误")
		return
	}

	// 2.判断用户名是否已存在
	found, err := l.svcCtx.Engine.Where("name = ?", req.Name).Limit(1).Get(new(models.UserBasic))
	if err!=nil{
		return nil,err
	}
	if found{
		return nil,errors.New("用户名已存在")
	}

	passwd, err := helper.Encrypt(req.Password)
	if err!=nil{
		return nil,errors.New("加密失败")
	}
	user := models.UserBasic{
		Identity: helper.UUID(),
		Name: req.Name,
		Password: passwd,
		Email: req.Email,
	}

	row, err := l.svcCtx.Engine.Insert(user)
	l.Logger.Info("insert user row: ",row)

	// 3.删除邮箱验证码
	_, err = l.svcCtx.RDB.Del(l.ctx, req.Email).Result()
	session.Commit()
	return
}
