package main

import (
	"fmt"
)

func main() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(7, -3))
}

// 29. 两数相除
func divide(dividend int, divisor int) int {
	// 特殊情况：溢出
	if dividend == -1<<31 && divisor == -1 {
		return 1<<31 - 1
	}
	res := 0
	negative := (dividend < 0) != (divisor < 0)
	dvd := abs(dividend)
	dvs := abs(divisor)
	for dvd >= dvs {
		temp := dvs
		mutitime := 1
		for dvd >= (temp << 1) {
			temp <<= 1
			mutitime <<= 1 // 增加了多少倍  就是出出来
		}
		dvd -= temp
		res = res + mutitime
	}
	if negative {
		return -res
	}
	return res
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
