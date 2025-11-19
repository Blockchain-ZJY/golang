package logic

import (
	"context"

	"my-micro-service/app/user/rpc/internal/svc"
	"my-micro-service/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	// 模拟数据库查询
	return &user.GetUserInfoResp{
		Id:       in.Id,
		Nickname: "go-zero高手",
		Avatar:   "https://example.com/avatar.png",
	}, nil

}
