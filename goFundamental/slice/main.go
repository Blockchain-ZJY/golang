package main

import (
	"fmt"
)

func myAppend(s []int) []int {
	// 这里 s 虽然改变了，但并不会影响外层函数的 s
	s = append(s, 100)
	return s
}

func myAppendPtr(s *[]int) {
	// 会改变外层 s 本身
	*s = append(*s, 100)
}

func main() {
	s := []int{1, 1, 1}
	fmt.Printf("slice variable addr: %p\n", &s)
	x := myAppend(s)
	fmt.Println(s)
	s = append(s, 2)
	fmt.Printf("slice variable addr: %p\n", &s)
	fmt.Printf("slice variable addr: %p,cap of x: %d\n", &x, cap(x))
	fmt.Println(x)
	fmt.Println(s)
}
