package main

import (
	"context"
	"fmt"
	"grpc_study/common"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := ":8080"
	// 使用 grpc.Dial 创建一个到指定地址的 gRPC 连接。
	// 此处使用不安全的证书来实现 SSL/TLS 连接
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(fmt.Sprintf("grpc connect addr [%s] 连接失败 %s", addr, err))
	}
	defer conn.Close()
	// 初始化客户端
	// client := common.NewCommonServiceClient(conn)
	// stream, err := client.GetUserInfo(context.Background(), &common.UserId{
	// 	Id: "张三",
	// })

	// for {
	// 	response, err := stream.Recv()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println(response)
	// }

	// stream, err := client.GetByCID(context.Background(), &common.FileCID{
	// 	CID: "张三",
	// })

	// // fils write
	// file, err := os.OpenFile("static/filsdownload.zip", os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// writer := bufio.NewWriter(file)
	// for {1
	// 	response, err := stream.Recv()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println("写入%d", len(response.Files))
	// 	writer.Write(response.Files)
	// }
	// writer.Flush()
	// fmt.Println(result1, err)

	client := common.NewUploadFileServiceStreamingClient(conn)

	stream, err := client.UploadFile(context.Background())

	for i := 0; i < 10; i++ {
		stream.Send(&common.ClientRequest{
			Files: []byte("张三"),
		})
		fmt.Printf("发送%d\n", i)
	}

	respo, err := stream.CloseAndRecv()
	fmt.Println(respo, err)
}
