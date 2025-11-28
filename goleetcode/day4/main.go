package main

import (
	"fmt"
	"strconv"
)

// 7. 整数翻转
func reverse(x int) int {
	if x == 0 {
		return 0
	}
	symble := true
	if x < 0 {
		symble = false
		x = -x
	}
	ans := 0
	LENS := 0
	y := x
	for y > 0 {
		LENS++
		y /= 10
	}

	for x > 0 {
		LENS--
		ans = intPow(10, LENS)*(x%10) + ans
		x /= 10
	}
	if ans > intPow(2, 31)-1 {
		return 0
	}
	if !symble {
		ans = -ans
	}

	return ans
}

// 9.回文数

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	i := 0
	j := len(s) - 1
	for {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
		if j < i {
			break
		}
	}
	return true
}
func intPow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

// 10. 正则表达式匹配

// 匹配任意单个字符
const AnyOneChar = '.'

// 匹配零个或多个前面的那一个元素
const RepeatChar = '*'

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

// 11. 盛最多水的容器
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	ans := 0
	for i < j {
		ans = max(ans, (min(height[i], height[j]) * (j - i)))
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return ans
}
func main() {
	// fmt.Println(isMatch("AAA", "AB*A*C*"))
	fmt.Println(maxArea([]int{1, 2, 4, 3}))
}
