package main

import "fmt"

func main() {
	str := "Hello, 面试鸭!"
	byteArray := []byte(str) // 转换时发生内存拷贝
	byteArray[0] = 'h'
	fmt.Println(str)
	fmt.Println(byteArray)
	fmt.Println(string(byteArray))
}
