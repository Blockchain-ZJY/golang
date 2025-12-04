package main

import (
	"fmt"
	"golang/leetcode/day/commen"
	"math/big"
)

// 42. 接雨水 单调栈
func trap(height []int) int {
	stk := Stack{}
	stk.Push(0)
	sum := 0
	for i := 1; i < len(height); i++ {
		for !stk.IsEmpty() && height[i] > height[stk.Peek()] {
			mid := stk.Peek()
			stk.Pop()
			if !stk.IsEmpty() {
				h := min(height[stk.Peek()], height[i]) - height[mid]
				w := i - stk.Peek() - 1
				sum += h * w
			}
		}
		stk.Push(i)
	}
	return sum
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}

type Stack struct {
	items []int
}

// 入栈
func (s *Stack) Push(ch int) {
	s.items = append(s.items, ch)
}

// 出栈
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

// 查看栈顶
func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		return 0
	}
	return s.items[len(s.items)-1]
}

// 判断是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// 栈大小
func (s *Stack) Size() int {
	return len(s.items)
}

// 2211. 道路上碰撞次数
// 用栈来记录当前状态和栈顶的关系
func countCollisions(directions string) int {
	stk := commen.Stack{}
	sum := 0
	stk.Push(rune(directions[0]))
	for i := 1; i < len(directions); i++ {
		switch directions[i] {
		case 'S':
			for stk.Peek() == 'R' {
				sum++
				stk.Pop()
			}
			for stk.Peek() == 'L' {
				stk.Pop()
			}
			stk.Push('S')
		case 'L':
			if stk.Peek() == 'S' {
				sum += 1
			}
			if stk.Peek() == 'R' {
				sum += 2
				stk.Pop()
				for stk.Peek() == 'R' {
					sum += 1
					stk.Pop()
				}
				stk.Push('S')
			}
		default:

			stk.Push('R')
		}
	}
	return sum
}

//438. 找到字符串中所有字母异位词

func findAnagrams(s string, p string) []int {
	primes := []*big.Int{
		big.NewInt(2), big.NewInt(3), big.NewInt(5), big.NewInt(7), big.NewInt(11),
		big.NewInt(13), big.NewInt(17), big.NewInt(19), big.NewInt(23), big.NewInt(29),
		big.NewInt(31), big.NewInt(37), big.NewInt(41), big.NewInt(43), big.NewInt(47),
		big.NewInt(53), big.NewInt(59), big.NewInt(61), big.NewInt(67), big.NewInt(71),
		big.NewInt(73), big.NewInt(79), big.NewInt(83), big.NewInt(89), big.NewInt(97),
		big.NewInt(101),
	}

	// 计算模式串乘积
	target := big.NewInt(1)
	for _, v := range p {
		target.Mul(target, primes[v-'a'])
	}

	n := len(p)
	if len(s) < n {
		return nil
	}

	// 初始化窗口乘积
	window := big.NewInt(1)
	for i := 0; i < n; i++ {
		window.Mul(window, primes[s[i]-'a'])
	}

	ans := []int{}
	if window.Cmp(target) == 0 {
		ans = append(ans, 0)
	}

	// 滑动窗口
	for i := n; i < len(s); i++ {
		// 移除左边字符
		window.Div(window, primes[s[i-n]-'a'])
		// 加入右边字符
		window.Mul(window, primes[s[i]-'a'])

		if window.Cmp(target) == 0 {
			ans = append(ans, i-n+1)
		}
	}

	return ans
}
