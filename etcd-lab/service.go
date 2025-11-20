package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	// 允许通过命令行参数指定端口，方便启动多个
	port := flag.Int("port", 8000, "service port")
	flag.Parse()

	// 服务地址 (key)
	serviceAddr := fmt.Sprintf("127.0.0.1:%d", *port)
	key := "/services/user/" + serviceAddr

	// 1. 连接 etcd
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 2. 创建租约 (Lease) - 10秒过期
	leaseResp, err := cli.Grant(context.Background(), 10)
	if err != nil {
		log.Fatal(err)
	}

	// 3. 注册服务 (Put + Lease)
	_, err = cli.Put(context.Background(), key, "running", clientv3.WithLease(leaseResp.ID))
	if err != nil {
		log.Fatal(err)
	}

	// 4. 自动续租 (KeepAlive)
	ch, err := cli.KeepAlive(context.Background(), leaseResp.ID)
	if err != nil {
		log.Fatal(err)
	}

	// 处理续租响应的协程
	go func() {
		for range ch {
			// 只是为了证明在续租，实际生产不用打印
			// fmt.Printf("心跳正常: %s \n", serviceAddr)
		}
	}()

	fmt.Printf("🟢 服务 [%s] 已启动，注册成功！(按 Ctrl+C 停止)\n", serviceAddr)

	// 5. 阻塞直到收到退出信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 模拟优雅退出 (也可以不写这一步，直接杀进程测试 etcd 的自动过期)
	fmt.Printf("🔴 服务 [%s] 正在停止...\n", serviceAddr)
	cli.Revoke(context.Background(), leaseResp.ID) // 立即撤销租约
}