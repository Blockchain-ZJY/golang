package main

import (
	"context"
	"fmt"
	hello_grpc "grpc_study/grpc_proto"
)

// HelloServer1 得有一个结构体，需要实现这个服务的全部方法,叫什么名字不重要
type HelloServer1 struct {
	hello_grpc.UnimplementedHelloServiceServer
}

// HelloServer1 得有一个结构体，需要实现这个服务的全部方法,叫什么名字不重要
type UserRegister struct {
	hello_grpc.UnimplementedUserServiceServer
}

func (UserRegister) UserRegisterFun(ctx context.Context, request *hello_grpc.UserRegister) (*hello_grpc.UserRegisterResponse, error) {
	fmt.Println("用户注册传参：", request.Name, request.Password)
	// 返回 HelloResponse 类型，而非 HelloRequest
	return &hello_grpc.UserRegisterResponse{
		Status: "你好", // 字段需与 HelloResponse 的定义匹配
		Code:   "Good",
	}, nil
}

func (HelloServer1) SayHello(ctx context.Context, request *hello_grpc.HelloRequest) (*hello_grpc.HelloResponse, error) {
	fmt.Println("入参：", request.Name, request.Message)
	// 返回 HelloResponse 类型，而非 HelloRequest
	return &hello_grpc.HelloResponse{
		Name:    "你好", // 字段需与 HelloResponse 的定义匹配
		Message: "Good",
	}, nil
}

// func main() {
// 	// 监听端口
// 	listen, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		grpclog.Fatalf("Failed to listen: %v", err)
// 	}

// 	// 创建一个gRPC服务器实例。
// 	s := grpc.NewServer()
// 	HelloServer := HelloServer1{}

// 	UserServer := UserRegister{}
// 	// 将server结构体注册为gRPC服务。
// 	hello_grpc.RegisterHelloServiceServer(s, HelloServer)
// 	hello_grpc.RegisterUserServiceServer(s, UserServer)

// 	fmt.Println("grpc server running :8080")
// 	// 开始处理客户端请求。
// 	s.Serve(listen)
// }
