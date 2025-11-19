package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/bytedance/sonic"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 写一个关于性能优化的 面试题目

func CopyFile(src, dst string) error {
	// 打开源文件
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// 创建目标文件
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// 使用 io.Copy 拷贝内容
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	// 确保写入完成
	err = destFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// 切片拷贝
	// for 循环遍历逐个拷贝-> copy 函数拷贝
	needcopy := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		needcopy[i] = i
	}
	copyed := make([]int, 1000)
	len := copy(copyed, needcopy)
	fmt.Println(len)

	// 文件拷贝
	// 先读取文件内容到内存，再写入到目标文件-> 使用 io 包的 Copy 函数
	// src := "static/test.zip"
	// dst := "static/test_copy.zip"

	// err := CopyFile(src, dst)
	// if err != nil {
	// 	fmt.Println("拷贝失败:", err)
	// } else {
	// 	fmt.Println("拷贝成功!")
	// }

	// 字符串拼接
	str1 := "Hello, "
	str2 := "21321"

	var builder strings.Builder // 是一种高效的字符串拼接方式,先声明一个 strings.Builder

	builder.WriteString(str1)
	builder.WriteString(str2)
	fmt.Println(builder.String())

	u := User{Name: "Alice", Age: 25}

	// 序列化为 JSON
	data, err := sonic.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data)) // 输出: {"name":"Alice","age":25}

	// 美化输出（带缩进）
	pretty, err := sonic.MarshalIndent(u, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(pretty))

	//log

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// 基本日志
	logger.Info("服务启动成功")

	// 带字段的日志
	logger.Info("用户登录",
		slog.String("username", "alice"),
		slog.Int("age", 25),
	)

	// 错误日志
	err1 := doSomething()
	if err1 != nil {
		logger.Error("操作失败",
			slog.String("op", "doSomething"),
			slog.Any("error", err),
		)
	}

	// Debug 日志（默认级别可能不会显示，需要配置 handler）
	logger.Debug("调试信息", slog.String("detail", "这里是一些调试内容"))
}
func doSomething() error {
	return fmt.Errorf("模拟错误")
}
