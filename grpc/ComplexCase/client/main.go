package main

import (
	"context"
	"io"
	"log"
	"time"

	orderpb "grpc/ComplexCase/proto/order"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	orderServiceAddr = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("无法连接到订单服务: %v", err)
	}
	defer conn.Close()

	client := orderpb.NewOrderClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// --- 1. 尝试下一个成功的订单 ---
	log.Println("--- 场景1: 下一个成功的订单 (product-123) ---")
	createReqSuccess := &orderpb.CreateOrderRequest{
		ProductId: "product-123",
		Quantity:  2,
		UserId:    "user-abc",
	}
	createRes, err := client.CreateOrder(ctx, createReqSuccess)
	if err != nil {
		log.Fatalf("创建订单失败: %v", err)
	}
	log.Printf("成功创建订单，ID: %s", createRes.GetOrderId())
	successfulOrderID := createRes.GetOrderId()

	// --- 2. 追踪这个成功的订单 (演示流式调用) ---
	log.Printf("\n--- 开始追踪订单 %s 的状态 ---", successfulOrderID)
	trackReq := &orderpb.TrackOrderRequest{OrderId: successfulOrderID}
	stream, err := client.TrackOrder(ctx, trackReq)
	if err != nil {
		log.Fatalf("无法开始追踪订单: %v", err)
	}

	// 循环接收来自服务端的流数据
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			// 流结束
			log.Println("--- 订单追踪结束 ---")
			break
		}
		if err != nil {
			log.Fatalf("接收状态更新时发生错误: %v", err)
		}
		log.Printf("收到订单状态更新: [%s]", res.GetOrderStatus())
	}

	// --- 3. 尝试下一个失败的订单 (库存不足) ---
	log.Println("\n--- 场景2: 下一个失败的订单 (product-456 库存不足) ---")
	createReqFail := &orderpb.CreateOrderRequest{
		ProductId: "product-456",
		Quantity:  1,
		UserId:    "user-xyz",
	}
	_, err = client.CreateOrder(ctx, createReqFail)
	if err != nil {
		log.Printf("订单创建失败 (符合预期): %v", err)
	}
}
