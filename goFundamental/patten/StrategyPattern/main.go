package main

import "fmt"

// 函数类型使用

// ===================================================================
// 1. 定义策略蓝图 (统一的函数类型)
// ===================================================================

// ShippingCostStrategy 是一个函数类型，它定义了所有运费计算策略的“蓝图”或“接口”。
// 任何符合这个函数签名的函数，都可以被看作是一种运费计算策略。
// 它的底层类型是: func(*Order) float64
type ShippingCostStrategy func(order *Order) float64

// ===================================================================
// 2. 定义核心业务对象
// ===================================================================

// Order 结构体代表一个订单。
type Order struct {
	Weight   float64 // 订单重量
	Distance float64 // 运输距离
	IsVIP    bool    // 是否为VIP客户

	// 这个字段是策略模式的核心。
	// 它可以持有任何一个符合 ShippingCostStrategy 蓝图的函数。
	shippingStrategy ShippingCostStrategy
}

// SetShippingStrategy 是一个辅助方法，用于为一个订单设置或更换运费计算策略。
func (o *Order) SetShippingStrategy(strategy ShippingCostStrategy) {
	o.shippingStrategy = strategy
}

// CalculateShippingCost 方法负责计算并返回订单的最终运fen.
// 它并不关心具体的计算逻辑，而是将计算任务委托给当前持有的策略函数。
func (o *Order) CalculateShippingCost() float64 {
	// 如果外部没有为订单指定任何策略，我们可以提供一个默认的策略。
	if o.shippingStrategy == nil {
		fmt.Println("-> (未指定策略, 使用默认的重量计算策略)")
		return calculateByWeight(o) // 调用默认策略
	}
	// 如果指定了策略，就调用该策略。
	return o.shippingStrategy(o)
}

// ===================================================================
// 3. 实现多个具体的策略函数
// ===================================================================

// calculateByWeight 策略一：按重量计算运费。
// 它的函数签名和 ShippingCostStrategy 完全一致。
func calculateByWeight(order *Order) float64 {
	fmt.Println("-> (执行策略: 按重量计算)")
	// 运费 = 重量 * 1.5元/kg
	return order.Weight * 1.5
}

// calculateByDistance 策略二：按距离计算运费。
// 它的函数签名和 ShippingCostStrategy 完全一致。
func calculateByDistance(order *Order) float64 {
	fmt.Println("-> (执行策略: 按距离计算)")
	// 运费 = 距离 * 0.5元/km
	return order.Distance * 0.5
}

// vipFreeShipping 策略三：为VIP客户提供免运费服务。
// 它的函数签名和 ShippingCostStrategy 完全一致。
func vipFreeShipping(order *Order) float64 {
	fmt.Println("-> (执行策略: VIP客户免运费)")
	return 0.0
}

// ===================================================================
// 4. 主函数 - 应用和演示
// ===================================================================

func main() {
	// --- Case 1: 一个普通订单，使用默认的重量计算策略 ---
	fmt.Println("--- 案例 1: 普通订单 ---")
	order1 := &Order{Weight: 10}
	cost1 := order1.CalculateShippingCost()
	fmt.Printf("最终运费: %.2f 元\n\n", cost1)

	// --- Case 2: 一个长途订单，我们为其动态指定距离计算策略 ---
	fmt.Println("--- 案例 2: 长途订单 ---")
	order2 := &Order{Distance: 500}
	// 关键点：因为 calculateByDistance 的签名和 ShippingCostStrategy 一致，
	// 所以它可以被直接当作参数传递给 SetShippingStrategy 方法。
	order2.SetShippingStrategy(calculateByDistance)
	cost2 := order2.CalculateShippingCost()
	fmt.Printf("最终运费: %.2f 元\n\n", cost2)

	// --- Case 3: 一个VIP订单，注入VIP免运费策略 ---
	fmt.Println("--- 案例 3: VIP 订单 ---")
	order3 := &Order{Weight: 20, Distance: 100, IsVIP: true}
	order3.SetShippingStrategy(vipFreeShipping)
	cost3 := order3.CalculateShippingCost()
	fmt.Printf("最终运费: %.2f 元\n\n", cost3)

	// --- 演示：同一个订单，在不同情况下可以更换策略 ---
	fmt.Println("--- 案例 4: 同一订单更换策略 ---")
	order4 := &Order{Weight: 5, Distance: 200}
	fmt.Println("初始状态:")
	cost4_initial := order4.CalculateShippingCost() // 使用默认策略
	fmt.Printf("最终运费: %.2f 元\n", cost4_initial)

	fmt.Println("\n升级为长途运输:")
	order4.SetShippingStrategy(calculateByDistance) // 更换为距离策略
	cost4_upgraded := order4.CalculateShippingCost()
	fmt.Printf("最终运费: %.2f 元\n", cost4_upgraded)
}
