package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// LoggerMiddleware 是一个正确的 Gin 中间件
// 它不需要接收 next 作为参数，因为它将通过 c.Next() 来控制流程
func LoggerMiddleware() gin.HandlerFunc {
	// 1. 返回一个符合 gin.HandlerFunc 签名的函数
	return func(c *gin.Context) {
		// 2. 请求到达中间件时，执行前置操作
		now := time.Now()
		// 3. ✨ 使用 c.Next() 来调用后续的处理函数（可能是其他中间件，也可能是最终的业务Handler）
		c.Next()
		// 4. 后续的处理函数执行完毕后，程序会回到这里，执行后置操作
		log.Printf("url: %s, latency: %v, status: %d", c.Request.URL, time.Since(now), c.Writer.Status())
	}
}
func MyMiddleware(c *gin.Context) {
	// 在处理函数之前执行的逻辑
	fmt.Println("Middleware: 处理请求前")
	// 执行下一个中间件或处理函数
	c.Next()
	// 在处理函数之后执行的逻辑
	fmt.Println("Middleware: 处理请求后")
}

func main() {
	r := gin.Default()
	r.Use(LoggerMiddleware())
	r.Use(MyMiddleware)
	r.GET("/ping", ping)
	r.POST("/ping1", ping)
	r.Run(":9000")
}
