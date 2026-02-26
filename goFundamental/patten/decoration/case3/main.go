package main

import (
	"errors"
	"fmt"
	"time"
)

// ==========================================
// 1. 核心层 (Core / Domain)
// 这一层定义业务逻辑和接口，绝对不依赖具体实现
// ==========================================

// Order 订单实体
type Order struct {
	ID    string
	Price int
}

// OrderRepository 核心抽象（接口/契约）
// 只要实现了这个接口，不管是存 MySQL 还是存文件，OrderService 都能用
type OrderRepository interface {
	Save(order *Order) error
}

// OrderService 核心业务服务
// 它只依赖 OrderRepository 接口，完全不知道数据库和缓存的存在
type OrderService struct {
	repo OrderRepository
}

func (s *OrderService) CreateOrder(order *Order) error {
	fmt.Println("--- 业务逻辑开始 ---")

	// 1. 纯粹的业务校验
	if order.Price < 0 {
		return errors.New("❌ 业务报错: 价格不能为负数")
	}

	// 2. 调用接口保存
	// Service 只管叫 "保存"，具体怎么存、存哪、要不要缓存，它不管
	if err := s.repo.Save(order); err != nil {
		return err
	}

	fmt.Println("--- 业务逻辑结束 ---")
	return nil
}

// ==========================================
// 2. 基础设施层 (Infrastructure / Implementation)
// 这一层负责具体的干活
// ==========================================

// MySQLRepo 具体的数据库实现 (基础实现)
type MySQLRepo struct {
	// 模拟数据库连接字符串
	DBUrl string
}

func (m *MySQLRepo) Save(order *Order) error {
	// 模拟数据库操作耗时
	fmt.Printf("   [MySQL] 正在写入数据库 (Url: %s)... \n", m.DBUrl)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("   [MySQL] ✅ 订单 %s 已保存到 MySQL\n", order.ID)
	return nil
}

// ==========================================
// 3. 装饰器层 (Decorator / Proxy)
// 这一层负责增强功能，比如加缓存、加日志
// ==========================================

// CachedOrderRepo 缓存装饰器
// 它既是一个 Repo (实现了接口)，又包含了一个 Repo (持有下层)
type CachedOrderRepo struct {
	next  OrderRepository // 持有 "下一层"
	redis string          // 模拟 Redis 客户端
}

func (c *CachedOrderRepo) Save(order *Order) error {
	// Step 1: 先调用 "下一层" (通常是数据库)
	// 如果数据库保存失败，缓存也就不用写了，直接返回错误
	if err := c.next.Save(order); err != nil {
		return err
	}

	// Step 2: 数据库成功后，执行 "增强逻辑" (写缓存)
	// 这里就是原本可能污染 Service 的代码，现在被隔离在这里
	c.setCache(order)

	return nil
}

func (c *CachedOrderRepo) setCache(order *Order) {
	fmt.Printf("   [Redis] ⚡️ 正在写入缓存 (Client: %s)...\n", c.redis)
	fmt.Printf("   [Redis] ✅ 订单 %s 缓存已设置\n", order.ID)
}

// ==========================================
// 4. 组装层 (Wiring / Composition)
// 在 main 函数里决定怎么拼装积木
// ==========================================

func main() {
	// 准备一个订单
	order := &Order{ID: "ORDER-2023-001", Price: 100}

	fmt.Println("====== 场景 1: 只有数据库 (V1 版本) ======")

	// 1. 创建基础的 MySQL Repo
	mysqlRepo := &MySQLRepo{DBUrl: "root:1234@tcp(localhost:3306)"}

	// 2. 注入 Service
	serviceV1 := &OrderService{repo: mysqlRepo}

	// 3. 运行
	serviceV1.CreateOrder(order)

	fmt.Println("\n\n====== 场景 2: 数据库 + Redis 缓存 (V2 版本) ======")
	fmt.Println(">> 注意：我们完全没有修改 OrderService 的代码！ <<")

	// 1. 还是先有 MySQL Repo
	mysqlRepoV2 := &MySQLRepo{DBUrl: "root:1234@tcp(localhost:3306)"}

	// 2. 创建缓存装饰器，把 MySQL Repo "包" 进去
	// 就像给手机套了个壳
	cachedRepo := &CachedOrderRepo{
		next:  mysqlRepoV2, // <--- 关键：把 MySQL 塞给 Cache
		redis: "Redis:6379",
	}

	// 3. 注入 Service
	// Service 根本不知道传入的是个 "带壳的" Repo，它只知道这个 Repo 能 Save
	serviceV2 := &OrderService{repo: cachedRepo}

	// 4. 运行
	serviceV2.CreateOrder(order)
}
