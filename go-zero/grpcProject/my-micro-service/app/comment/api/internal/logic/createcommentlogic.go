// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"my-micro-service/app/comment/api/internal/svc"
	"my-micro-service/app/comment/api/internal/types"
	"my-micro-service/app/comment/rpc/commentclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CreateCommentReq) (resp *types.CreateCommentResp, err error) {
	// 调用 Comment RPC
	rpcResp, err := l.svcCtx.CommentRpc.CreateComment(l.ctx, &commentclient.CreateCommentReq{
		UserId:    req.UserId,
		ArticleId: req.ArticleId,
		Content:   req.Content,
	})

	if err != nil {
		return nil, err
	}

	return &types.CreateCommentResp{
		CommentId: rpcResp.CommentId,
	}, nil
}
