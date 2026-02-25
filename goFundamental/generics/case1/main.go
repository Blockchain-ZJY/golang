package main

type TestAPI interface {
	Method1() error
	Method2() error
}

// 具体类型Tester1
type Tester1 struct{}

func (t *Tester1) Method1() error { return nil }
func (t *Tester1) Method2() error { return nil }

// 具体类型Tester2
type Tester2 struct{}

func (t *Tester2) Method1() error { return nil }
func (t *Tester2) Method2() error { return nil }
func (t *Tester2) Method3() error { return nil }

// Number 定义了一个约束，它是一个包含所有整数和浮点数类型的接口。
// 任何实现了这些类型之一的类型都满足这个约束。
type Number interface {
	int | int64 | float32 | float64
}

func Sort[T Number](a []T) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
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
	var t1 Tester1
	var t2 Tester2
	// 定义接口变量
	var api1, api2, api3 TestAPI
	api1 = &t1
	api2 = &t2
	api3 = &t2
	// 方法调用
	api1.Method1()
	api2.Method1()
	api3.Method2()
}
