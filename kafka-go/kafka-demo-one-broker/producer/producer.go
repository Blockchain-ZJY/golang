package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

// --- 配置信息 ---
const (
	// Kafka Broker 的地址
	kafkaBrokers = "localhost:9092,localhost:9093"
	// 我们要向哪个主题发送消息
	topic = "orders.created"
	// 每条消息发送的时间间隔
	produceInterval = 10 * time.Second
)

// Order 结构体，与我们之前的消费者代码保持一致
type Order struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Amount    float64   `json:"amount"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	log.Println("启动 Kafka 生产者...")

	// 1. 设置优雅关闭
	// 创建一个 context，当接收到关闭信号时，它会被取消
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 监听操作系统的中断信号 (Ctrl+C)
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// 2. 配置并创建 Kafka Writer (生产者)
	// Writer 是线程安全的，可以在整个应用程序中共享一个实例
	writer := &kafka.Writer{
		Addr:     kafka.TCP(strings.Split(kafkaBrokers, ",")...),
		Topic:    topic,               // 默认发送到的主题
		Balancer: &kafka.LeastBytes{}, // Balancer 决定消息被发送到哪个分区
		// Async: true, // 可以开启异步模式以提高吞吐量，但会牺牲一些可靠性
	}

	// 使用 defer 确保在程序退出时，Writer 被正确关闭。
	// 这非常重要，因为它可以确保所有缓冲区的消息都被发送出去。
	defer func() {
		log.Println("正在关闭 Kafka Writer...")
		if err := writer.Close(); err != nil {
			log.Fatalf("关闭 Kafka Writer 失败: %v", err)
		}
		log.Println("Kafka Writer 已成功关闭。")
	}()

	// 启动一个 goroutine 在后台持续生产消息
	go func() {
		orderCounter := 0
		for {
			select {
			// 如果 context 被取消 (因为收到了关闭信号)，则立即退出循环
			case <-ctx.Done():
				log.Println("生产者接收到关闭信号，停止生产消息...")
				return
			// 每隔 produceInterval 的时间，就执行一次生产逻辑
			case <-time.After(produceInterval):
				orderCounter++

				// 3. 创建一条模拟订单消息
				order := Order{
					ID:        fmt.Sprintf("ORD-%d", 10000+orderCounter),
					UserID:    fmt.Sprintf("USER-%d", 100+rand.Intn(10)),
					Amount:    float64(rand.Intn(20000)) / 100.0, // 0.00 to 200.00
					Currency:  "USD",
					CreatedAt: time.Now().UTC(),
				}

				// 4. 将 Go 结构体序列化为 JSON 字节
				orderBytes, err := json.Marshal(order)
				if err != nil {
					log.Printf("JSON 序列化失败: %v", err)
					continue // 跳过这条消息
				}

				// 5. 准备要发送的 Kafka 消息
				msg := kafka.Message{
					// Key 对于分区的选择很重要。相同 Key 的消息通常会进入同一个分区。
					// 这保证了同一个订单的所有相关事件的处理顺序。
					Key:   []byte(order.ID),
					Value: orderBytes,
				}

				// 6. 发送消息
				// WriteMessages 是一个同步操作，它会等待 Kafka 的确认
				err = writer.WriteMessages(ctx, msg)
				if err != nil {
					// 如果 context 被取消，这里会立即返回错误
					if ctx.Err() != nil {
						// 这是预期的退出，不需要打印错误
						break
					}
					log.Printf("写入消息失败: %v", err)
				} else {
					// 成功发送后，可以从 Writer 的统计信息中获取分区和偏移量
					stats := writer.Stats()
					log.Printf("成功发送消息 -> 主题: %s, 订单ID: %s", stats.Topic, order.ID)
				}
			}
		}
	}()

	log.Printf("已连接到 Kafka, 正在向主题 '%s' 发送消息...", topic)
	log.Println("按 Ctrl+C 退出。")

	// 主 goroutine 阻塞在这里，等待接收到关闭信号
	<-sigchan

	// 一旦接收到信号，就调用 cancel()
	// 这会触发上面 goroutine 中 `case <-ctx.Done()` 的逻辑
	log.Println("接收到关闭信号，开始优雅关闭...")
	cancel()

	// 给一点时间让 defer 函数执行完毕
	time.Sleep(1 * time.Second)
}
