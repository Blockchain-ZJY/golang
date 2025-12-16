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

func main() {
	r := gin.Default()

	// Public routes
	// 公开路由
	r.POST("/login", LoginHandler)

	// Protected routes group
	// 受保护的路由组
	auth := r.Group("/auth")
	auth.Use(JWTAuthMiddleware())
	{
		auth.GET("/profile", ProfileHandler)
	}

	r.Run(":8080")
}

func LoginHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mock authentication (Replace with database check)
	// For demo: username "admin", password "password"
	// 模拟身份验证 (请替换为数据库检查)
	// 演示用: 用户名 "admin", 密码 "password"
	if user.Username == "admin" && user.Password == "password" {
		tokenString, err := GenerateToken(user.Username)
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
