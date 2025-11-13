package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, "https://cloudflare-eth.com")
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	defer client.Close()

	fmt.Println("成功连接到以太坊网络")
}
