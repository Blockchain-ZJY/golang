package main

import (
	"fmt"
	"sync"
	"time"
)

// TokenBucket 令牌桶结构
type TokenBucket struct {
	capacity   int        // 桶容量
	tokens     int        // 当前令牌数
	rate       int        // 每秒生成令牌数
	lastRefill time.Time  // 上次补充时间
	mu         sync.Mutex // 并发安全
}

// NewTokenBucket 初始化令牌桶
func NewTokenBucket(rate, capacity int) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity, // 初始时桶满
		rate:       rate,
		lastRefill: time.Now(),
	}
}

// refill 根据时间补充令牌 lazy update 
func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	tb.lastRefill = now

	// 计算新增令牌数
	newTokens := int(elapsed * float64(tb.rate))
	if newTokens > 0 {
		tb.tokens += newTokens
		if tb.tokens > tb.capacity {
			tb.tokens = tb.capacity
		}
	}
}

// Allow 尝试获取令牌
func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

func main() {
	// 每秒生成 2 个令牌，桶容量 5
	limiter := NewTokenBucket(2, 5)

	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Println("Request", i, "allowed at", time.Now())
		} else {
			fmt.Println("Request", i, "denied at", time.Now())
		}
		time.Sleep(300 * time.Millisecond)
	}
}
