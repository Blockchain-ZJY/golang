package logic

import (
	"context"

	"pilipili/app/video/rpc/internal/svc"
	"pilipili/app/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoDetailLogic {
	return &GetVideoDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 2. 核心读操作
func (l *GetVideoDetailLogic) GetVideoDetail(in *video.GetVideoDetailReq) (*video.GetVideoDetailResp, error) {
	// todo: add your logic here and delete this line

	return &video.GetVideoDetailResp{}, nil
}
