// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package private

import (
	"context"

	"pilipili/app/video/api/internal/svc"
	"pilipili/app/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// UP主投稿视频
func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishLogic) Publish(req *types.PublishReq) (resp *types.PublishResp, err error) {
	// todo: add your logic here and delete this line

	return
}
