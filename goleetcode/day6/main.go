package main

import (
	"fmt"
	"sort"
)

// 2141. 同时运行 N 台电脑的最长时间
func maxRunTime(n int, batteries []int) int64 {
	sort.Ints(batteries)
	fmt.Println(batteries)
	return 0
}
func ismatch() {

}

// 20 有效括号
func isValid(s string) bool {
	stack := CharStack{}
	ismatch := func(a int32, b int32) bool {
		if string(a) == "(" && string(b) == ")" {
			return true
		}
		if string(a) == "{" && string(b) == "}" {
			return true
		}
		if string(a) == "[" && string(b) == "]" {
			return true
		}
		return false
	}
	for _, v := range s {
		if string(v) == "(" || string(v) == "{" || string(v) == "[" {
			stack.Push(v)
		} else {
			topv, _ := stack.Peek()
			if !ismatch(topv, v) {
				return false
			}
			stack.Pop()
		}
	}
	return stack.IsEmpty()
}
func main() {
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
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

// 18.四数之和

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	n := len(nums)
	if n == 4 {
		if nums[0]+nums[1]+nums[2]+nums[3] == target {
			return append(res, []int{nums[0], nums[1], nums[2], nums[3]})
		} else {
			return res
		}
	}
	sort.Ints(nums)
	l, r := 0, 0
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r = j+1, n-1
			for l < r {
				forsum := nums[l] + nums[r] + nums[i] + nums[j]
				if forsum == target {
					res = append(res, []int{nums[l], nums[r], nums[i], nums[j]})

					//退
					for nums[l+1] == nums[l] {
						l++
						if l == r {
							break
						}
					}
					l++
					for nums[r-1] == nums[r] && l < r {
						r--
						if l == r {
							break
						}
					}
					r--

				} else if forsum > target {
					r--
				} else {
					l++
				}
			}

		}
	}
	return res
}
