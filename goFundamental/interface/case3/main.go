package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

func main() {
	var s Speaker
	d := Dog{}

	s = d // 接口赋值

	// 类型断言：将接口转换为具体类型
	dog, ok := s.(Dog)
	if ok {
		fmt.Println("转换成功:", dog)
		dog.Speak()
	} else {
		fmt.Println("转换失败")
	}

	// 错误的断言，触发运行时错误
	// cat := s.(string) // panic: interface conversion: main.Dog is not string
}
