package main

import "fmt"

func testrune() {
	s := "你好，世界"
	runes := []rune(s)
	for _, r := range runes {
		fmt.Printf("%c ", r)
	}

}

func main() {
	s := "今天happy"
	fmt.Println(len(s))         //输出11
	fmt.Println(len([]rune(s))) //输出7
	fmt.Println(string([]rune(s)[0]))
	testrune()
}
