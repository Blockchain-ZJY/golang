package main

import "fmt"

// 定义一个接口
type Animal interface {
	Speak() string
}

// Dog 类型实现了 Animal 接口
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// Cat 类型实现了 Animal 接口
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

// Person 类型实现了 Animal 接口
type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello!"
}

// 多态函数：接收 Animal 接口
func MakeSound(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Person{Name: "John"}}

	// 遍历不同类型，调用接口方法
	for _, animal := range animals {
		MakeSound(animal)
	}
}
