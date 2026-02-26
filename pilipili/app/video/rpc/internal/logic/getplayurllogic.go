package logic

import (
	"context"

	"pilipili/app/video/rpc/internal/svc"
	"pilipili/app/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPlayUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPlayUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPlayUrlLogic {
	return &GetPlayUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 4. 获取实际播放流 (通常前端拿到视频详情后，再单独请求播放地址)
func (l *GetPlayUrlLogic) GetPlayUrl(in *video.GetPlayUrlReq) (*video.GetPlayUrlResp, error) {
	// todo: add your logic here and delete this line

	return &video.GetPlayUrlResp{}, nil
}
