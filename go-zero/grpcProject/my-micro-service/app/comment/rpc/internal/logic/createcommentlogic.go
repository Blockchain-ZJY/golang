package logic

import (
	"context"
	"fmt"

	"my-micro-service/app/comment/rpc/comment"
	"my-micro-service/app/comment/rpc/internal/svc"
	"my-micro-service/app/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCommentLogic) CreateComment(in *comment.CreateCommentReq) (*comment.CreateCommentResp, error) {
	// 1. 跨微服务调用：去 User 服务验证用户是否存在或获取昵称
	userResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &userclient.GetUserInfoReq{
		Id: in.UserId,
	})
	if err != nil {
		return nil, err
	}

	// 2. 执行评论逻辑 (这里打印模拟入库)
	fmt.Printf("用户 [%s] 在文章 %d 发表了评论: %s\n", userResp.Nickname, in.ArticleId, in.Content)

	return &comment.CreateCommentResp{
		CommentId: 888, // 模拟生成的ID
	}, nil
}
