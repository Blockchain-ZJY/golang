package main

import "fmt"

// Number 定义了一个约束，它是一个包含所有整数和浮点数类型的接口。
// 任何实现了这些类型之一的类型都满足这个约束。
type Number interface {
	int | int64 | float32 | float64
}

// SumNumbers 是一个泛型函数。
// [T Number] 是类型参数声明：
// - T 是类型参数的名称，它在函数内部代表一个具体的类型。
// - Number 是 T 必须满足的约束。
func SumNumbers[T Number](numbers []T) (ans T) {
	var total T // total 的类型是 T，如果传入 []int，它就是 int；如果传入 []float64，它就是 float64
	for _, n := range numbers {
		total += n
	}
	return total
}

func main() {
	// --- 使用泛型函数处理不同类型的切片 ---

	// 1. 用于 int 切片
	intSlice := []int{1, 2, 3, 4, 5}
	sumOfInts := SumNumbers(intSlice) // Go 编译器会自动推断 T 的类型是 int
	fmt.Printf("Generic sum of ints: %d (Type: %T)\n", sumOfInts, sumOfInts)

	// 2. 用于 float64 切片
	float64Slice := []float64{1.1, 2.2, 3.3, 4.4}
	sumOfFloat64s := SumNumbers(float64Slice) // 编译器推断 T 的类型是 float64
	fmt.Printf("Generic sum of floats: %f (Type: %T)\n", sumOfFloat64s, sumOfFloat64s)

	// 3. 用于 int64 切片
	int64Slice := []int64{100, 200, 300}
	sumOfInt64s := SumNumbers(int64Slice) // 编译器推断 T 的类型是 int64
	fmt.Printf("Generic sum of int64s: %d (Type: %T)\n", sumOfInt64s, sumOfInt64s)
}
