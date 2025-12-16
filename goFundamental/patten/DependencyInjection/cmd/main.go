package main

import (
	"context"
	"fmt"
	di "gojourney/goFundamental/patten/DependencyInjection"
)

func main() {
	// 1. 创建依赖
	// 我们在这里手动注入“依赖项”。
	// 在大型应用程序中，您可能会使用像 google/wire 或 uber-go/dig 这样的 DI 容器。
	dbConnString := "postgres://user:pass@localhost:5432/mydb"
	userRepo := di.NewSQLUserRepository(dbConnString)

	// 2. 将依赖项注入服务
	userService := di.NewUserService(userRepo)

	// 3. 使用服务
	ctx := context.Background()
	err := userService.Register(ctx, "u1", "Alice", "alice@example.com")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	u, _ := userService.GetUser(ctx, "u1")
	fmt.Printf("找到用户: %+v\n", u)
}
