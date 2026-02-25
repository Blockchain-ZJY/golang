package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// 监听端口
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server started on :8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}

		// 每个连接一个 goroutine
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {

		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected")
			return
		}

		fmt.Printf("Received: %s", msg)

		// 回显给客户端 这是写给客户端的
		conn.Write([]byte("Echo: " + msg))
	}
}
