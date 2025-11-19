package logic

import (
	"context"

	"grpcdemo/grpcdemo"
	"grpcdemo/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMyOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyOrderLogic {
	return &GetMyOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMyOrderLogic) GetMyOrder(in *grpcdemo.User) (*grpcdemo.Orders, error) {
	// todo: add your logic here and delete this line
	logx.Infof("Received message: %s", in.Name)

	return &grpcdemo.Orders{
		Orders: []*grpcdemo.Order{
			{
				OrderId: "1",
				Items: []*grpcdemo.OrderItem{
					{
						ProductId:   "1",
						ProductName: "Item 1",
						Price:       10,
					},
				},
			},
			{
				OrderId: "2",
				Items: []*grpcdemo.OrderItem{
					{
						ProductId:   "2",
						ProductName: "Item 2",
						Price:       20,
					},
				},
			},
		},
	}, nil
}
