package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Reader：从控制台读取
	reader := bufio.NewReader(os.Stdin)

	// Writer：写到控制台（os.Stdout）
	writer := bufio.NewWriter(os.Stdout)

	for {
		fmt.Print("请输入内容: ")

		// 阻塞读取一行
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取失败:", err)
			continue
		}
		fmt.Println(text)
		// 去掉换行符
		text = strings.TrimSpace(text)
		// 附加一些内容
		output := fmt.Sprintf("你输入的是: %s | 附加内容: [OK]\n", text)
		// 使用 writer 写入
		_, err = writer.WriteString(output)
		if err != nil {
			fmt.Println("写入失败:", err)
			continue
		}
		// Writer 是带缓冲的，必须 Flush 才会真正输出
		writer.Flush()
	}
}
