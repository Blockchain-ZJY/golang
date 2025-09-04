package main

import "fmt"

type Order struct {
	Weight   float64
	Distance float64
	IsVIP    bool

	// 这个字段是策略模式的核心。
	shippingStrategy ShippingStrategy
}

type ShippingStrategy interface {
	CalculateShippingCost(order Order) float64
}

// input strategy must be a instance of ShippingStrategy
func (o *Order) SetShippingStrategy(strategy ShippingStrategy) {
	o.shippingStrategy = strategy
}

func (o Order) CalculateShippingCost() float64 {
	if o.shippingStrategy == nil {
		// default strategy
		return WeightStrategy{}.CalculateShippingCost(o)
	}
	return o.shippingStrategy.CalculateShippingCost(o)
}

type WeightStrategy struct{}

func (ws WeightStrategy) CalculateShippingCost(order Order) float64 {
	fmt.Println("WeightStrategy, 计算运费")
	return order.Weight * 10
}

type DistanceStrategy struct{}

func (ds DistanceStrategy) CalculateShippingCost(order Order) float64 {
	fmt.Println("DistanceStrategy, 计算运费")
	return order.Distance * 2
}

type vipFreeShipping struct{}

func (vips vipFreeShipping) CalculateShippingCost(order Order) float64 {
	fmt.Println("vipFreeShipping, 计算运费")
	return 0
}

func main() {
	order := &Order{Weight: 10, Distance: 100, IsVIP: true}

	order.SetShippingStrategy(WeightStrategy{})
	fmt.Println(order.CalculateShippingCost())

	order.SetShippingStrategy(DistanceStrategy{})
	fmt.Println(order.CalculateShippingCost())

	order.SetShippingStrategy(vipFreeShipping{})
	fmt.Println(order.CalculateShippingCost())
}
