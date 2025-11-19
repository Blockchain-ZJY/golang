// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"my-micro-service/app/comment/api/internal/config"
	"my-micro-service/app/comment/rpc/commentclient"
	"my-micro-service/app/user/rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	CommentRpc commentclient.Comment
	UserRpc    userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		CommentRpc: commentclient.NewComment(zrpc.MustNewClient(c.CommentRpc)),
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
