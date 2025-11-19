package logic

import (
	"context"

	"my-micro-service/app/comment/api/internal/svc"
	"my-micro-service/app/comment/api/internal/types"

	// 引入 user rpc 包
	"my-micro-service/app/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	// 直接调用 User RPC
	rpcResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &userclient.GetUserInfoReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	// 将 RPC 的返回结果 转换成 API 的返回结果
	return &types.GetUserInfoResp{
		Id:       rpcResp.Id,
		Nickname: rpcResp.Nickname,
		Avatar:   rpcResp.Avatar,
	}, nil
}
