// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"firstdemo/internal/svc"
	"firstdemo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FirstdemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFirstdemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FirstdemoLogic {
	return &FirstdemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FirstdemoLogic) Firstdemo(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
