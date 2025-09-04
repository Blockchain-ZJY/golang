package main

import (
	"fmt"
	"sync"
)

func main() {
	userch := make(chan string)
	var wg sync.WaitGroup // 1. 声明一个 WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // 2. 每启动一个 goroutine，就让计数器加 1
		go func(i int) {
			defer wg.Done() // 4. goroutine 结束时，调用 Done() 让计数器减 1
			userch <- fmt.Sprintf("user_%d", i)
		}(i)
	}

	// 启动一个新的 goroutine 来等待所有的发送者完成工作，然后关闭 channel
	go func() {
		wg.Wait()     // 3. 阻塞，直到 WaitGroup 计数器归零
		close(userch) // 5. 关闭 channel
	}()

	// for range 会在 channel 关闭后自动结束循环
	for user := range userch {
		fmt.Println(user)
	}
}
