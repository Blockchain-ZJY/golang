package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var serviceMap sync.Map // 用来在本地存储当前可用的 IP

func main() {
	// 1. 连接 etcd
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	prefix := "/services/user/"

	// 2. 第一次启动，先获取当前所有的服务 (GET)
	resp, err := cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}

	for _, kv := range resp.Kvs {
		serviceMap.Store(string(kv.Key), string(kv.Value))
	}
	printServices() // 打印初始列表

	// 3. 启动 Watch 协程，实时监听变化
	go func() {
		watchChan := cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
		for watchResp := range watchChan {
			for _, event := range watchResp.Events {
				key := string(event.Kv.Key)
				switch event.Type {
				case mvccpb.PUT:
					fmt.Printf("\n[🔔 新增节点] %s 上线了！\n", key)
					serviceMap.Store(key, string(event.Kv.Value))
				case mvccpb.DELETE:
					fmt.Printf("\n[⚠️ 节点移除] %s 下线了！\n", key)
					serviceMap.Delete(key)
				}
				printServices() // 每次变化都重新打印列表
			}
		}
	}()

	// 阻塞主程
	select {}
}

// 辅助函数：打印当前内存里所有的服务
func printServices() {
	fmt.Println("--------- 当前可用服务列表 ---------")
	count := 0
	serviceMap.Range(func(key, value interface{}) bool {
		fmt.Printf(" -> %s \n", key)
		count++
		return true
	})
	if count == 0 {
		fmt.Println(" (空) 暂无服务可用")
	}
	fmt.Println("----------------------------------")
}
