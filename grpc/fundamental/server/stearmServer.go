// package main

// import (
// 	"fmt"
// 	common "grpc_study/common"
// 	"net"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/grpclog"
// )

// type Common struct {
// 	common.UnimplementedCommonServiceServer
// }

// func main() {
// 	// 监听端口
// 	listen, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		grpclog.Fatalf("Failed to listen: %v", err)
// 	}

// 	// 创建一个gRPC服务器实例。
// 	s := grpc.NewServer()
// 	CommonServer := Common{}

// 	// 将server结构体注册为gRPC服务。
// 	common.RegisterCommonServiceServer(s, CommonServer)

// 	fmt.Println("grpc server running :8080")
// 	// 开始处理客户端请求。
// 	s.Serve(listen)
// }

package main

import (
	"fmt"
	common "grpc_study/common"
	"io"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type Common struct {
	common.UnimplementedCommonServiceServer
}

type ClientStreamType struct {
	common.UnimplementedUploadFileServiceStreamingServer
}

func (c Common) GetUserInfo(req *common.UserId, stream common.CommonService_GetUserInfoServer) error {
	fmt.Println("入参：", req.Id)

	// 发送多条用户信息
	users := []*common.UserInfo{
		{Name: "user1", Email: "user1@example.com"},
		{Name: "user2", Email: "user2@example.com"},
		{Name: "user23", Email: "user2@example.com"},
		{Name: "user42", Email: "user2@example.com"},
	}

	for _, user := range users {
		if err := stream.Send(user); err != nil {
			return err
		}
	}
	return nil
}
func (c Common) GetByCID(req *common.FileCID, stream common.CommonService_GetByCIDServer) error {
	fmt.Println("i got a cid：", req.CID)
	file, err := os.Open("static/ZhiLvYunZhang.zip")
	if err != nil {
		return err
	}
	for {
		buf := make([]byte, 1024)
		_, err = file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		if err := stream.Send(&common.FileResponse{Files: buf}); err != nil {
			return err
		}
	}
	return nil
}

func (c ClientStreamType) UploadFile(stream common.UploadFileServiceStreaming_UploadFileServer) error {
	for i := 0; i < 10; i++ {
		repos, err := stream.Recv()
		fmt.Println(repos, err, "1")
		if err == io.EOF {
			break
		}
	}
	stream.SendAndClose(&common.UploadFileResponse{Status: []byte("i am done")})
	return nil
}

func main() {
	// 监听端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	// 创建一个gRPC服务器实例。
	s := grpc.NewServer()
	// CommonServer := Common{}
	ClintStream := ClientStreamType{}

	// 将server结构体注册为gRPC服务。
	// common.RegisterCommonServiceServer(s, CommonServer)
	common.RegisterUploadFileServiceStreamingServer(s, ClintStream)

	fmt.Println("grpc server running :8080")
	// 开始处理客户端请求。
	s.Serve(listen)
}

// 流式传输案例：
// 1、客户端不知道服务端什么时候结束
// 2、下载文件
