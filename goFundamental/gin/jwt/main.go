package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User 代表此示例中的简单用户
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes
	// 公开路由
	r.GET("/login", LoginHandler)

	// Protected routes group
	// 受保护的路由组
	auth := r.Group("/auth")
	auth.Use(JWTAuthMiddleware())
	{
		auth.GET("/profile", ProfileHandler)
	}
	return r
}

func main() {

	r := setupRouter()
	r.Run(":8080")
}

func LoginHandler(c *gin.Context) {
	// 从 URL 参数获取用户名和密码
	username := c.Query("username")
	password := c.Query("password")

	// 简单校验参数是否为空
	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
		return
	}

	// Mock authentication (Replace with database check)
	// 演示用: 用户名 "admin", 密码 "password"
	if username == "admin" && password == "password" {
		tokenString, err := GenerateToken(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

func ProfileHandler(c *gin.Context) {
	// Retrieve username set by the middleware
	// 获取中间件设置的用户名
	username, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to your profile", "username": username})
}
