//一个很常见的案例，我有一个获取ip的协程，但是这是一个耗时操作，用户随时可能会取消
//如果用户取消了，那么之前那个获取协程的函数就要停止运行

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wait = sync.WaitGroup{}

func main() {
	t1 := time.Now()
	wait.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {

		ip, err := GetIp(ctx)
		fmt.Println(ip, err)
	}()
	wait.Add(1)
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
		wait.Done()
	}()

	wait.Wait()
	fmt.Println("执行完成", time.Since(t1))
}

func GetIp(ctx context.Context) (ip string, err error) {
	fmt.Println("开始获取ip")
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("提前取消", ctx.Err().Error())
			err = ctx.Err()
			wait.Done()
			return
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("获取ip完成")
	ip = "192.168.200.1"
	wait.Done()
	return
}
