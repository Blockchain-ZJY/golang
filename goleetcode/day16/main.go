package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

func main() {
	fmt.Println(maxMatrixSum([][]int{{-1, 0, -1}, {-2, 1, 3}, {3, 2, 2}}))
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

// 198. 打家劫舍
func rob(nums []int) int {
	var dfs func(i int) int
	m := make(map[int]int)
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if _, ok := m[i]; ok {
			return m[i]
		}
		res := max(dfs(i-2)+nums[i], dfs(i-1))
		m[i] = res
		return res
	}
	return dfs(len(nums) - 1)
}

// 武林大会打擂台 O(1) 的空间复杂度实现 找到众数
func majorityElement(nums []int) (ans int) {
	hp := 0
	for _, x := range nums {
		if hp == 0 {
			ans, hp = x, 1
		} else if x == ans {
			hp++
		} else {
			hp--
		}
	}
	return
}

// 76. 最小覆盖子串
func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}
	target := [128]int{}
	for i := 0; i < len(t); i++ {
		target[t[i]]++
	}
	var ismatch func(a [128]int, target [128]int) bool
	ismatch = func(a, target [128]int) bool {
		for x, _ := range target {
			if target[x] > a[x] {
				return false
			}
		}
		return true
	}
	ans := [128]int{}
	for i := 0; i < len(t); i++ {
		ans[s[i]]++
	}
	if ismatch(ans, target) {
		return s[0:len(t)]
	}
	l, r := 0, len(t)
	resl, resr := 0, len(s)
	found := false
	for r < len(s) {
		ans[s[r]]++
		for ismatch(ans, target) {
			found = true
			if r-l < resr-resl {
				resl = l
				resr = r
			}
			ans[s[l]]--
			l++
		}
		r++
	}
	if !found {
		return ""
	}
	return s[resl : resr+1]
}

// 46. 全排列
func permute(nums []int) (ans [][]int) {
	path := []int{}
	vis := make([]bool, len(nums))
	var dfs func(length int)
	dfs = func(length int) {
		if length == len(nums) {
			t := []int{}
			t = append(t, path...)
			ans = append(ans, t)
		}
		for i := 0; i < len(nums); i++ {
			if !vis[i] {
				vis[i] = true
				path = append(path, nums[i])
				dfs(length + 1)
				path = path[:len(path)-1]
				vis[i] = false
			}
		}
	}
	dfs(0)
	return
}

// 1390. 四因数
func sumFourDivisors(nums []int) (ans int) {
	for x := range nums {
		t := 0
		sum := 0
		for j := 2; j <= int(math.Sqrt(float64(nums[x]))); j++ {
			if nums[x]%j == 0 {
				t = j
				sum++
			}
		}
		if sum == 1 && t != nums[x]/t {
			ans += t + 1 + nums[x] + nums[x]/t
		}
	}
	return
}

func largestEven(s string) string {
	b := []byte(s)
	for i := len(s) - 1; i >= 0; i-- {
		if (b[i]-'0')%2 == 0 {
			return s[:i+1]
		}
	}
	return ""
}

// Q2. 单词方块 II
func wordSquares(words []string) (ans [][]string) {

	path := []string{}
	vis := make([]bool, len(words)+1)

	var dfs func()
	dfs = func() {
		if len(path) == 4 {
			// 最后检查 bottom 的两个条件
			if path[3][0] == path[1][3] && path[3][3] == path[2][3] {
				ans = append(ans, append([]string{}, path...))
			}
			return
		}

		for i := 0; i < len(words); i++ {
			if vis[i] {
				continue
			}
			// 剪枝：根据当前 path 长度直接比较
			if len(path) == 1 { // left 必须首字母等于 top[0]
				if words[i][0] != path[0][0] {
					continue
				}
			}
			if len(path) == 2 { // right 必须首字母等于 top[3]
				if words[i][0] != path[0][3] {
					continue
				}
			}
			if len(path) == 3 { // bottom 必须首字母等于 left[3]，末字母等于 right[3]
				if words[i][0] != path[1][3] || words[i][3] != path[2][3] {
					continue
				}
			}

			vis[i] = true
			path = append(path, words[i])
			dfs()
			path = path[:len(path)-1]
			vis[i] = false
		}
	}

	dfs()
	sort.Slice(ans, func(i, j int) bool {
		for k := 0; k < 4; k++ {
			if ans[i][k] != ans[j][k] {
				return ans[i][k] < ans[j][k]
			}
		}
		return false
	})
	return
}

// 1975. 最大方阵和
func maxMatrixSum(matrix [][]int) int64 {
	ans, count, min := 0, 0, math.MaxInt64
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] < 0 {
				count++
				ans -= matrix[i][j]
			} else {
				ans += matrix[i][j]
			}
			if math.Abs(float64(matrix[i][j])) < math.Abs(float64(min)) {
				min = matrix[i][j]
			}
		}
	}
	if count%2 == 0 {
		return int64(ans)
	} else {
		return int64(ans - 2*int(math.Abs(float64(min))))
	}
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	k := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
