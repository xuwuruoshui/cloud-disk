package logic

import (
	"cloud-disk/core/define"
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendRequest) (resp *types.MailCodeSendReply, err error) {

	// 1.判断邮箱是否注册
	found, err := l.svcCtx.Engine.Where("email = ?", req.Email).Limit(1).Get(new(models.UserBasic))
	if err!=nil{
		return nil,err
	}
	if found{
		return nil,errors.New("该邮箱已被注册")
	}

	// 2. 判断是否已经生成过验证码
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err!=nil && err!=redis.Nil{
		return nil,err
	}
	if code==""{
		// 2.1 不存在则创建
		code = helper.RandCode()

		// 2.2 存在则存储
		l.svcCtx.RDB.Set(l.ctx,req.Email,code,time.Second * time.Duration(define.CodeExpire))
	}

	// 3.发送验证码
	err = helper.MailSendCode(req.Email, code)

	return
}
