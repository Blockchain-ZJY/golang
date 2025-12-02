package main

import (
	"fmt"
)

func readBinaryWatch(turnedOn int) []string {
	leds := []int{8, 4, 2, 1, 32, 16, 8, 4, 2, 1}
	res := []string{}
	path := []int{} // 存索引，而不是值
	backtrackIndex(leds, turnedOn, 0, path, &res)
	return res
}

func backtrackIndex(leds []int, k, start int, path []int, res *[]string) {
	if len(path) == k {
		hour, minute := 0, 0
		for _, idx := range path {
			if idx < 4 { // 索引区分小时
				hour += leds[idx]
			} else {
				minute += leds[idx]
			}
		}
		if hour < 12 && minute < 60 {
			*res = append(*res, fmt.Sprintf("%d:%02d", hour, minute))
		}
		return
	}
	for i := start; i < len(leds); i++ {
		path = append(path, i)
		backtrackIndex(leds, k, i+1, path, res)
		path = path[:len(path)-1]
	}
}

func combine(n int, k int) [][]int {
	arr := []int{}
	res := [][]int{}
	path := []int{}
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}
	var backtracking func(n int, k int, start int)
	backtracking = func(n int, k int, start int) {
		if len(path) == k {
			// 必须复制 path
			res = append(res, append([]int(nil), path...))
			return
		}
		for i := start; i < n; i++ {
			path = append(path, arr[i])
			backtracking(n, k, i+1)
			path = path[:len(path)-1] // 回溯
		}
	}
	backtracking(n, k, 0)
	return res
}

// 17. 电话号码的字母组合
func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}
	wordMap := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}
	res := []string{}

	var backtrackingLetter func(index int, path string)
	backtrackingLetter = func(index int, path string) {
		if index == len(digits) {
			res = append(res, path)
			return
		}
		letters := wordMap[digits[index]]
		for i := 0; i < len(letters); i++ {
			backtrackingLetter(index+1, path+string(letters[i]))
		}
	}
	backtrackingLetter(0, "")
	return res
}

// 22. 括号生成
func generateParenthesis(n int) []string {
	path := "("
	res := []string{}
	var backtracking func(path string)
	backtracking = func(path string) {
		if len(path) == 2*n {
			if Stackmatch(path) {
				res = append(res, path)
			}
			return
		}
		backtracking(path + string("("))
		backtracking(path + string(")"))
	}
	backtracking(path)
	return res
}

func Stackmatch(s string) bool {
	cstack := CharStack{}
	for _, c := range s {
		if string(c) == "(" {
			cstack.Push(c)
		} else {
			topv, _ := cstack.Peek()
			if string(topv) != "(" {
				return false
			}
			cstack.Pop()
		}
	}
	return cstack.IsEmpty()
}

func main() {
	fmt.Println(generateParenthesis(3))
}

type CharStack struct {
	items []rune
}

// 入栈
func (s *CharStack) Push(ch rune) {
	s.items = append(s.items, ch)
}

// 出栈
func (s *CharStack) Pop() (rune, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

// 查看栈顶
func (s *CharStack) Peek() (rune, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

// 判断是否为空
func (s *CharStack) IsEmpty() bool {
	return len(s.items) == 0
}

// 栈大小
func (s *CharStack) Size() int {
	return len(s.items)
}
