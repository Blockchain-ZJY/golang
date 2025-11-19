package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	// 导入生成的 inventory 和 order 包
	inventorypb "grpc/ComplexCase/proto/inventory"
	orderpb "grpc/ComplexCase/proto/order"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

// 定义库存服务的地址
const (
	inventoryServiceAddr = "localhost:50052"
)

// 模拟订单数据库，key是order_id, value是status
var orders = make(map[string]string)

// server 结构体既要实现 OrderServer 接口，
// 也要持有一个到 Inventory 服务的客户端连接。
type server struct {
	orderpb.UnimplementedOrderServer
	inventoryClient inventorypb.InventoryClient
}

// CreateOrder 是一元 RPC。它接收创建订单的请求。
func (s *server) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderReply, error) {
	log.Printf("[订单服务] 收到创建订单请求: 用户 %s, 商品 %s, 数量 %d", req.GetUserId(), req.GetProductId(), req.GetQuantity())

	// ✨✨ 关键步骤: 作为 gRPC 客户端调用库存服务 ✨✨
	stockReq := &inventorypb.CheckStockRequest{
		ProductId:        req.GetProductId(),
		QuantityRequired: req.GetQuantity(),
	}

	// 发起 RPC 调用到库存服务
	stockReply, err := s.inventoryClient.CheckStock(ctx, stockReq)
	if err != nil {
		log.Printf("[订单服务] 调用库存服务失败: %v", err)
		// 返回一个标准的 gRPC 内部错误
		return nil, status.Errorf(codes.Internal, "检查库存时发生内部错误: %v", err)
	}

	// 检查库存服务的返回结果
	if !stockReply.IsInStock {
		log.Printf("[订单服务] 库存不足，拒绝订单")
		// 返回一个标准的 gRPC 前置条件失败错误
		return nil, status.Errorf(codes.FailedPrecondition, "商品库存不足")
	}

	// 如果库存充足，则创建订单
	orderID := fmt.Sprintf("order-%d", time.Now().UnixNano())
	orders[orderID] = "PROCESSING" // 设置初始状态
	log.Printf("[订单服务] 订单 %s 创建成功", orderID)

	return &orderpb.CreateOrderReply{OrderId: orderID, Message: "订单创建成功"}, nil
}

// TrackOrder 是服务端流式 RPC。它会持续向客户端推送订单状态。
func (s *server) TrackOrder(req *orderpb.TrackOrderRequest, stream orderpb.Order_TrackOrderServer) error {
	orderID := req.GetOrderId()
	log.Printf("[订单服务] 收到订单 %s 的状态追踪请求", orderID)

	if _, ok := orders[orderID]; !ok {
		return status.Errorf(codes.NotFound, "订单 %s 不存在", orderID)
	}

	// 模拟订单状态随时间变化
	statusUpdates := []string{"PROCESSING", "SHIPPED", "IN_TRANSIT", "DELIVERED"}
	for _, st := range statusUpdates {
		// 检查客户端是否已断开连接
		if stream.Context().Err() != nil {
			log.Printf("[订单服务] 客户端断开连接，停止追踪订单 %s", orderID)
			return nil
		}

		time.Sleep(2 * time.Second) // 模拟处理延迟

		reply := &orderpb.TrackOrderReply{OrderStatus: st}
		// 使用 stream.Send() 发送数据流到客户端
		if err := stream.Send(reply); err != nil {
			log.Printf("[订单服务] 发送状态更新失败: %v", err)
			return err
		}
		log.Printf("[订单服务] 已发送订单 %s 的状态更新: %s", orderID, st)
	}

	log.Printf("[订单服务] 订单 %s 追踪完成", orderID)
	return nil
}

func main() {
	// --- Part 1: 作为 gRPC 客户端连接到库存服务 ---
	conn, err := grpc.Dial(inventoryServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("无法连接到库存服务: %v", err)
	}
	defer conn.Close()
	// 创建库存服务的客户端存根
	inventoryClient := inventorypb.NewInventoryClient(conn)
	log.Println("[订单服务] 已成功连接到库存服务")

	// --- Part 2: 作为 gRPC 服务端启动，监听自己的端口 ---
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("监听端口 :50051 失败: %v", err)
	}

	s := grpc.NewServer()
	// 注册服务实现，注意这里我们将 inventoryClient 注入到了 server 结构体中
	orderpb.RegisterOrderServer(s, &server{inventoryClient: inventoryClient})

	log.Println("[订单服务] 正在监听端口 :50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
