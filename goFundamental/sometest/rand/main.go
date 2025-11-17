package main

import "fmt"

type Book struct {
	title string
}

func (b *Book) printTitle() {
	fmt.Println(b.title)
}

func main() {
	b := Book{title: "Go Programming"}
	p := b
	b.printTitle() // 语法糖
	// 调用方法时，Go 会自动解引用
	p.printTitle() // 等价于 (*p).printTitle()
}
