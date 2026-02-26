package logic

import (
	"context"

	"pilipili/app/video/rpc/internal/svc"
	"pilipili/app/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchGetVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchGetVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchGetVideoLogic {
	return &BatchGetVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 3. 高并发批量读操作 (供推荐系统、API网关调用)
func (l *BatchGetVideoLogic) BatchGetVideo(in *video.BatchGetVideoReq) (*video.BatchGetVideoResp, error) {
	// todo: add your logic here and delete this line

	return &video.BatchGetVideoResp{}, nil
}
