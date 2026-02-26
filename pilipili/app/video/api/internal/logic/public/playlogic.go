// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package public

import (
	"context"

	"pilipili/app/video/api/internal/svc"
	"pilipili/app/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取真实的视频流媒体播放地址
func NewPlayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayLogic {
	return &PlayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlayLogic) Play(req *types.PlayReq) (resp *types.PlayResp, err error) {
	// todo: add your logic here and delete this line

	return
}
