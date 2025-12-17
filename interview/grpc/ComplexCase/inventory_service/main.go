package main

import (
	"context"
	"log"
	"net"

	inventorypb "grpc/ComplexCase/proto/inventory"

	"google.golang.org/grpc"
)

// 模拟一个库存数据库
var stock = map[string]int32{
	"product-123": 10,
	"product-456": 0,
}

type server struct {
	inventorypb.UnimplementedInventoryServer
}

func (s *server) CheckStock(ctx context.Context, req *inventorypb.CheckStockRequest) (*inventorypb.CheckStockReply, error) {
	productID := req.GetProductId()
	qtyRequired := req.GetQuantityRequired()

	log.Printf("[库存服务] 收到库存检查请求: 商品ID %s, 需要数量 %d", productID, qtyRequired)

	currentStock, ok := stock[productID]
	if !ok {
		log.Printf("[库存服务] 商品 %s 不存在", productID)
		return &inventorypb.CheckStockReply{IsInStock: false}, nil
	}

	if currentStock >= qtyRequired {
		log.Printf("[库存服务] 商品 %s 库存充足 (%d >= %d)", productID, currentStock, qtyRequired)
		return &inventorypb.CheckStockReply{IsInStock: true}, nil
	}

	log.Printf("[库存服务] 商品 %s 库存不足 (%d < %d)", productID, currentStock, qtyRequired)
	return &inventorypb.CheckStockReply{IsInStock: false}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052") // 注意端口是 50052
	if err != nil {
		log.Fatalf("监听端口失败: %v", err)
	}

	s := grpc.NewServer()
	inventorypb.RegisterInventoryServer(s, &server{})

	log.Println("[库存服务] 正在监听端口 :50052...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
