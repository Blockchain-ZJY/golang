package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

// --- Kafka 及应用配置 ---
const (
	// Kafka Broker 地址
	kafkaBrokers = "localhost:9092,localhost:9093"
	// 主题配置
	topicOrderCreated   = "orders.created"   // 接收新订单的主题
	topicOrderProcessed = "orders.processed" // 处理成功订单的主题
	topicOrderFailed    = "orders.failed"    // 处理失败订单的死信主题
	// 消费者组 ID
	groupID = "order-processor-group"
	// 并发处理的工作单元数量
	numWorkers = 5
)

// Order 定义了订单的数据结构
type Order struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	CreatedAt     time.Time `json:"created_at"`
	Status        string    `json:"status,omitempty"` // 用于记录处理结果
	FailureReason string    `json:"failure_reason,omitempty"`
}

// --- 全局组件 ---
var (
	logger      *slog.Logger  // 结构化日志记录器
	kafkaWriter *kafka.Writer // 用于向其他主题生产消息
)

func main() {
	// 初始化结构化日志
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("启动订单处理服务...")

	// 创建一个可以被取消的 context，用于优雅关闭
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 初始化 Kafka Writer，用于将处理结果发送到其他主题
	// 这是一个通用的 Writer，可以向任何主题发送消息
	kafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP(strings.Split(kafkaBrokers, ",")[0]),
		Balancer: &kafka.LeastBytes{},
	}
	defer func() {
		if err := kafkaWriter.Close(); err != nil {
			logger.Error("关闭 Kafka Writer 失败", "error", err)
		}
	}()

	// 配置 Kafka Reader
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        strings.Split(kafkaBrokers, ","),
		GroupID:        groupID,
		Topic:          topicOrderCreated,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: 0,    // 禁用自动提交
	})
	defer func() {
		if err := reader.Close(); err != nil {
			logger.Error("关闭 Kafka Reader 失败", "error", err)
		}
		logger.Info("Kafka Reader 已成功关闭")
	}()

	logger.Info("消费者已连接", "topic", topicOrderCreated, "groupID", groupID)

	// --- 设置工作池和消息分发 ---
	// 创建一个 channel，用于从 Reader 接收消息并分发给 Worker
	messages := make(chan kafka.Message, numWorkers)
	var wg sync.WaitGroup // WaitGroup 用于等待所有 worker 完成任务

	// 启动指定数量的 Worker Goroutine
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i+1, &wg, messages, reader)
	}

	// --- 优雅关闭处理 ---
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// 启动一个 goroutine，在后台从 Kafka 拉取消息并发送到 channel
	go func() {
		for {
			// FetchMessage 会阻塞直到获取到新消息或 context 被取消
			msg, err := reader.FetchMessage(ctx)
			if err != nil {
				if ctx.Err() != nil { // Context 被取消，正常退出
					return
				}
				logger.Error("从 Kafka 拉取消息失败", "error", err)
				continue
			}
			// 将消息发送到 channel，等待 worker 消费
			messages <- msg
		}
	}()

	logger.Info("服务已就绪，等待订单消息...", "workers", numWorkers)

	// 等待关闭信号
	<-sigchan
	logger.Info("接收到关闭信号，正在进行优雅关闭...")

	// 信号触发后，关闭 context
	cancel()

	// 等待所有 worker 优雅地完成当前任务
	close(messages) // 关闭 channel，让 worker 退出循环
	wg.Wait()       // 等待所有 worker 执行完毕

	logger.Info("所有任务已处理完毕，服务已关闭")
}

// worker 是实际处理订单消息的函数
func worker(ctx context.Context, id int, wg *sync.WaitGroup, messages <-chan kafka.Message, reader *kafka.Reader) {
	defer wg.Done()
	logger := logger.With("worker_id", id) // 为每个 worker 创建一个带 ID 的子 logger
	logger.Info("Worker 已启动")

	// 这是一个阻塞操作，会一直等待，知道message 关闭 woker函数才结束
	for msg := range messages {
		// 解析订单
		var order Order
		if err := json.Unmarshal(msg.Value, &order); err != nil {
			logger.Error("JSON 解析失败", "error", err, "offset", msg.Offset)
			// 对于无法解析的消息，直接发送到死信队列
			forwardToDLQ(ctx, &order, msg.Value, "Invalid JSON format")
			// 即使处理失败，也要提交 offset，防止消息被重复消费
			if err := reader.CommitMessages(ctx, msg); err != nil {
				logger.Error("提交 offset 失败", "error", err, "offset", msg.Offset)
			}
			continue
		}

		l := logger.With("order_id", order.ID, "offset", msg.Offset)
		l.Info("开始处理订单")

		// --- 模拟订单处理逻辑 ---
		// 可能是调用其他微服务、操作数据库、检查库存等
		// 这里我们随机模拟成功或失败
		processedOrder, err := processOrder(order)
		if err != nil {
			l.Error("订单处理失败", "error", err)
			// 将失败的订单发送到死信队列
			forwardToDLQ(ctx, &processedOrder, nil, err.Error())
		} else {
			l.Info("订单处理成功")
			// 将处理成功的订单发送到下一个主题
			forwardToProcessed(ctx, processedOrder)
		}

		// --- 手动提交 Offset ---
		// 无论成功还是失败，只要逻辑走完，就提交 offset
		// 这样可以确保每条消息只被处理一次
		if err := reader.CommitMessages(ctx, msg); err != nil {
			l.Error("提交 offset 失败", "error", err)
		} else {
			l.Info("Offset 提交成功")
		}
	}
	logger.Info("Worker 已关闭")

}

// processOrder 模拟订单处理的核心业务逻辑
func processOrder(order Order) (Order, error) {
	// 模拟耗时操作
	time.Sleep(time.Duration(100+time.Now().UnixNano()%100) * time.Millisecond)

	// 模拟随机失败，例如库存不足或支付验证失败
	if time.Now().UnixNano()/100%10 == 0 {
		fmt.Println(time.Now().UnixNano())
		order.Status = "FAILED"
		return order, fmt.Errorf("库存YOUYOUOYU不足")
	}

	order.Status = "PROCESSED"
	return order, nil
}

// forwardToProcessed 将处理成功的订单发送到 `orders.processed` 主题
func forwardToProcessed(ctx context.Context, order Order) {
	orderBytes, err := json.Marshal(order)
	if err != nil {
		logger.Error("序列化成功订单失败", "error", err, "order_id", order.ID)
		return
	}

	err = kafkaWriter.WriteMessages(ctx, kafka.Message{
		Topic: topicOrderProcessed,
		Key:   []byte(order.ID),
		Value: orderBytes,
	})
	if err != nil {
		logger.Error("开始处理订单发送消息到 processed 主题失败", "error", err, "order_id", order.ID)
	}
}

// forwardToDLQ 将处理失败的订单发送到死信队列 `orders.failed`
func forwardToDLQ(ctx context.Context, order *Order, originalValue []byte, reason string) {
	order.FailureReason = reason

	var valueBytes []byte
	var err error

	if originalValue != nil {
		// 如果是原始消息，直接使用
		valueBytes = originalValue
	} else {
		// 否则序列化包含失败原因的订单对象
		valueBytes, err = json.Marshal(order)
		if err != nil {
			logger.Error("序列化失败订单失败", "error", err, "order_id", order.ID)
			return
		}
	}

	err = kafkaWriter.WriteMessages(ctx, kafka.Message{
		Topic: topicOrderFailed,
		Key:   []byte(order.ID),
		Value: valueBytes,
		// 可以添加 Header 来存储额外的元数据
		Headers: []kafka.Header{{Key: "failure-reason", Value: []byte(reason)}},
	})
	if err != nil {
		logger.Error("发送消息到 failed (DLQ) 主题失败", "error", err, "order_id", order.ID)
	}
}
