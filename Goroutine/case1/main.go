package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// 启动一个新的Goroutine来执行say("World")
	go say("World")

	// main函数自身也是一个Goroutine
	say("Hello main routine")
}
