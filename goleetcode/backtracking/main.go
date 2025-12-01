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
func main() {
	fmt.Println(letterCombinations("23"))
}
