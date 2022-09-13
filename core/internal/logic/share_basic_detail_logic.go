package logic

import (
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (resp *types.ShareBasicDetailReply, err error) {

	// 分享点击次数加1
	_, err = l.svcCtx.Engine.Exec("update share_basic set click_num=click_num+1 where identity = ?", req.Identity)
	if err!=nil{
		return nil,err
	}

	// 获取资源信息
	resp = new(types.ShareBasicDetailReply)
	_, err = l.svcCtx.Engine.Table("share_basic sb").
	Select("ur.name,rp.ext,rp.size,rp.path,ur.repository_identity").
		Join("left","user_repository ur","sb.user_repository_identity=ur.identity").
		Join("left", "repository_pool rp", "ur.repository_identity=rp.identity").
		Where("sb.identity=?", req.Identity).Get(resp)

	return
}
