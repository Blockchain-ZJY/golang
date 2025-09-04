package main

import "fmt"

// 会跑的接口
type Runner interface {
	Run()
}

// 会吃的接口
type Eater interface {
	Eat()
}

// 组合接口：动物接口，既会跑又会吃
type Animal interface {
	Runner
	Eater
}

// 定义一个Dog类型
type Dog struct {
	Name string
}

// 为Dog实现Runner接口
func (d Dog) Run() {
	fmt.Printf("%s 在跑...\n", d.Name)
}

// 为Dog实现Eater接口
func (d Dog) Eat() {
	fmt.Printf("%s 在吃...\n", d.Name)
}

func main() {
	var animal Animal
	dog := Dog{Name: "旺财"}

	// 因为Dog同时实现了Run()和Eat()方法，所以它也实现了Animal接口
	animal = dog
	animal.Run() // 输出: 旺财 在跑...
	animal.Eat() // 输出: 旺财 在吃...
}

/*

**案例解析：**

`Animal` 接口通过嵌入 `Runner` 和 `Eater` 接口，要求实现它的类型必须同时具备 `Run()` 和 `Eat()` 两个方法。这是一种强大的代码组织方式。

### 总结

| 知识点       | 描述                                                                       | 关键语法/概念                              |
| :----------- | :------------------------------------------------------------------------- | :----------------------------------------- |
| **定义与实现** | 接口是方法签名的集合，类型通过实现其所有方法来**隐式实现**接口。             | `type Sayer interface { Say() }`           |
| **多态**     | 接口变量可以持有任何实现了该接口的具体类型的值，常用于函数参数以提高复用性。 | `func Action(s Sayer)`                     |
| **空接口**   | `interface{}` 不包含任何方法，因此所有类型都实现了它，可用于存储任意类型的值。 | `var v interface{}`                         |
| **类型断言** | 用于检查接口变量底层存储的具体类型，并将其转换为该具体类型。               | `value, ok := i.(T)`                       |
| **类型选择** | `switch` 的一种特殊形式，用于更优雅地对多种类型进行判断和处理。              | `switch v := i.(type) { case T: ... }`     |
| **接口组合** | 通过在一个接口中嵌入其他接口来创建新的接口。                               | `type Animal interface { Runner; Eater }` |

掌握接口是从Go语言入门到进阶的关键一步，它深刻地影响着Go代码的设计哲学和项目结构。*/
