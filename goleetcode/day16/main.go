package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(generate(5))
}

func plusOne(digits []int) (ans []int) {
	carry := 1
	n := len(digits) - 1
	for n >= 0 || carry != 0 {
		sum := digits[n] + carry
		digits[n] = sum % 10
		carry = sum / 10
		n--
		if n == -1 && carry != 0 {
			ans = append(ans, 1)
			ans = append(ans, digits...)
			return
		}
	}
	return digits
}

type sk struct {
	arr []int
}

func (s *sk) top() int {
	return s.arr[len(s.arr)-1]
}
func (s *sk) push(a int) {
	s.arr = append(s.arr, a)
}

func (s *sk) pop() {
	s.arr = s.arr[:len(s.arr)-1]
}

// 394. 字符串解码 3[a2[cb]]
// func decodeString(s string) string {
// 	stack := &sk{
// 		arr: []byte{},
// 	}
// 	for i := 0; i < len(s); i++ {
// 		if unicode.IsDigit(rune(s[i])) || s[i] == '[' || unicode.IsLetter(rune(s[i])) {
// 			stack.push(s[i])
// 		} else {
// 			t := []byte{}
// 			for stack.top() != '[' {
// 				t = append(t, stack.top())
// 				stack.pop()
// 			} //3[a2[cb]]
// 			slices.Reverse(t)
// 			stack.pop()
// 			numBytes := []byte{}
// 			for len(stack.arr) > 0 && unicode.IsDigit(rune(stack.top())) {
// 				numBytes = append(numBytes, stack.top())
// 				stack.pop()
// 			}
// 			slices.Reverse(numBytes)
// 			rp := 0
// 			for _, b := range numBytes {
// 				rp = rp*10 + int(b-'0')
// 			}
// 			for i := 0; i < rp; i++ {
// 				for j := 0; j < len(t); j++ {
// 					stack.push(t[j])
// 				}
// 			}
// 		}
// 	}
// 	return string(stack.arr)
// }

// 739. 每日温度
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := &sk{
		arr: []int{},
	}
	stack.push((0))
	for i := 1; i < len(temperatures); i++ {
		for len(stack.arr) > 0 && temperatures[i] > temperatures[stack.top()] {
			ans[stack.top()] = i - stack.top()
			stack.pop()
		}
		stack.push(i)
	}
	return ans
}

// 961. 在长度 2N 的数组中找出重复 N 次的元素
func repeatedNTimes(nums []int) (ans int) {
	m := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			ans = nums[i]
			return
		}
		m[nums[i]] = struct{}{}
	}
	return
}

// 31. 下一个排列
// [1,2,3]
// [1,3,2]
// [2,1,3]
// [2,3,1]
// [3,1,2]
// [3,2,1]
func nextPermutation(nums []int) {
	// 1. 找到第一个下降的坐标i
	// 2. 找到i之后的位置j(满足递增顺序) 交换
	// 3. reverse i之后的部分
	for i := len(nums) - 1; i >= 0; i-- {
		if i != len(nums)-1 && nums[i] < nums[i+1] {
			j := len(nums) - 1
			for j > i && nums[j] <= nums[i] {
				j--
			}
			fmt.Println()
			nums[i], nums[j] = nums[j], nums[i]
			slices.Reverse(nums[i+1:])
			fmt.Println(nums)
			return
		} else if i == 0 {
			slices.Reverse(nums)
			return
		}
	}
	fmt.Println(nums)
}

// 杨辉三角
func generate(numRows int) (ans [][]int) {
	for i := 0; i < numRows; i++ {
		t := []int{}
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				t = append(t, 1)
			} else {
				if i > 1 {
					t = append(t, ans[i-1][j-1]+ans[i-1][j])
				}
			}
		}
		ans = append(ans, t)
	}
	return ans
}
