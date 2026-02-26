package logic

import (
	"context"

	"pilipili/app/video/rpc/internal/svc"
	"pilipili/app/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 1. 核心写操作
func (l *PublishVideoLogic) PublishVideo(in *video.PublishReq) (*video.PublishResp, error) {
	// todo: add your logic here and delete this line

	return &video.PublishResp{}, nil
}
