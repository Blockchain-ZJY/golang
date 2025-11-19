package main

import (
	"context"
	"flag"
	"fmt"

	"grpcdemo/grpcdemo"
	orderclient "grpcdemo/moretestfromaccount2order"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
)

var configFile = flag.String("f", "etc/OrderClintconfig.yaml", "the config file")

func main() {
	flag.Parse()
	// ------------------ 1. 定义并加载配置 ------------------
	var config zrpc.RpcClientConf
	// RpcClientConf 是 zrpc 客户端的标准配置结构体
	conf.MustLoad(*configFile, &config)

	// ------------------ 2. 创建 zrpc 客户端 ------------------
	// 使用 zrpc.MustNewClient 创建一个客户端，它会自动处理连接、服务发现等
	clint, err := zrpc.NewClient(config)

	if err != nil {
		return
	}
	// ------------------ 3. 实例化具体服务的客户端 ------------------
	// 使用 goctl 生成的 New 函数，将 zrpc.Client 传入
	// 这样我们就得到了一个可以调用 MoreTestFromAccount2Order 服务方法的客户端
	orderclint := orderclient.NewMoreTestFromAccount2Order(clint)

	userRequest := &grpcdemo.User{
		Id:   "user-123",
		Name: "Tony Stark",
		Age:  50,
	}

	res, _ := orderclint.GetMyOrder(context.Background(), userRequest)
	fmt.Println(res)
	// ------------------ 4. 准备并发送 RPC 请求 ------------------
	// 准备请求数据

	// 调用远程方法，注意这里不再需要自己管理 context 的超时，
	// zrpc 的配置中可以统一设置超时时间

	// ------------------ 5. 处理并打印响应结果 ------------------

	// flag.Parse()

	// // ------------------ 1. 定义并加载配置 ------------------
	// // RpcClientConf 是 zrpc 客户端的标准配置结构体
	// var c zrpc.RpcClientConf
	// conf.MustLoad(*configFile, &c)

	// // ------------------ 2. 创建 zrpc 客户端 ------------------
	// // 使用 zrpc.MustNewClient 创建一个客户端，它会自动处理连接、服务发现等
	// client, err := zrpc.NewClient(c)
	// if err != nil {
	// 	fmt.Printf("创建 zrpc client 失败: %v\n", err)
	// 	return
	// }

	// // ------------------ 3. 实例化具体服务的客户端 ------------------
	// // 使用 goctl 生成的 New 函数，将 zrpc.Client 传入
	// // 这样我们就得到了一个可以调用 MoreTestFromAccount2Order 服务方法的客户端
	// orderClient := moretestfromaccount2orderclient.NewMoreTestFromAccount2Order(client)

	// // ------------------ 4. 准备并发送 RPC 请求 ------------------
	// // 准备请求数据
	// userRequest := &grpcdemo.User{
	// 	Id:   "user-123",
	// 	Name: "Tony Stark",
	// 	Age:  50,
	// }

	// // 调用远程方法，注意这里不再需要自己管理 context 的超时，
	// // zrpc 的配置中可以统一设置超时时间
	// ordersResponse, err := orderClient.GetMyOrder(context.Background(), userRequest)
	// if err != nil {
	// 	fmt.Printf("调用 GetMyOrder 失败: %v\n", err)
	// 	return
	// }

	// // ------------------ 5. 处理并打印响应结果 ------------------
	// fmt.Printf("成功获取到用户 '%s' 的订单信息\n", userRequest.GetName())

	// if len(ordersResponse.Orders) == 0 {
	// 	fmt.Println("该用户没有任何订单。")
	// 	return
	// }

	// for _, order := range ordersResponse.Orders {
	// 	fmt.Println("==================================")
	// 	fmt.Printf("  订单ID: %s\n", order.OrderId)
	// 	fmt.Printf("  订单总价: %.2f\n", order.TotalAmount)
	// 	fmt.Printf("  订单状态: %s\n", order.Status.String())
	// 	fmt.Println("  订单商品列表:")
	// 	for _, item := range order.Items {
	// 		fmt.Printf("    - 商品: %s, 数量: %d, 单价: %.2f\n",
	// 			item.ProductName, item.Quantity, item.Price)
	// 	}
	// 	fmt.Println("==================================")
	// }
}
