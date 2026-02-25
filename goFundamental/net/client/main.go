package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	for {
		fmt.Print("Input: ")
		text, _ := reader.ReadString('\n')

		// 发送到服务端
		conn.Write([]byte(text))

		// 读取服务端回显
		reply, _ := serverReader.ReadString('\n')
		fmt.Println("Server:", reply)
	}
}
