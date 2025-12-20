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
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	solveSudoku(board)
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

// 37. 解数独
func solveSudoku(board [][]byte) {
	count := 0
	location := make(map[int][]int)
	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				location[count] = []int{i, j}
				count++
			}
		}
	}
	// 判断在第 i 行或第 j 列是否存在值 v
	match := func(i, j int, v byte) bool {
		// 检查第 i 行
		for col := 0; col < 9; col++ {
			if board[i][col] == v {
				return false // 行里已经有 v
			}
		}
		// 检查第 j 列
		for row := 0; row < 9; row++ {
			if board[row][j] == v {
				return false // 列里已经有 v
			}
		}
		startRow := (i / 3) * 3
		startCol := (j / 3) * 3
		for r := startRow; r < startRow+3; r++ {
			for c := startCol; c < startCol+3; c++ {
				if board[r][c] == v {
					return false
				}
			}
		}
		return true // 行和列都没有 v，可以放置
	}
	found := false
	var fitin func(n int)
	fitin = func(n int) {
		if found {
			return
		}
		if n == count {
			found = true
			return
		}
		nexti := location[n][0]
		nextj := location[n][1]
		for v := byte('1'); v <= '9'; v++ {
			if match(nexti, nextj, v) {
				board[nexti][nextj] = v
				fitin(n + 1)
				if found {
					return
				}
				board[nexti][nextj] = '.'
			}
		}
	}
	fitin(0)
	// for _, v := range board {
	// 	fmt.Println(string(v))
	// }

	// 这里 result 就是完整的答案路径 fmt.Println("路径:", string(result))
}
