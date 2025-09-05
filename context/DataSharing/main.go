package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "name", "context data sharing")
	GetUser(ctx)
}

func GetUser(ctx context.Context) {
	// 获取用户名
	fmt.Println(ctx.Value("name"))
}
