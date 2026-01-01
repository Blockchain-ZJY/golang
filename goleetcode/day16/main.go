package main

import (
	"fmt"
)

func main() {
	// fmt.Println(decodeString("3[a2[cb]]"))
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
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
