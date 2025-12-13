package main

import (
	"fmt"
	"math"
	"sort"
	"unicode"
)

func main() {
	// fmt.Println(validateCoupons([]string{"SAVE20", "", "PHARMA5", "SAVE@20"}, []string{"restaurant", "grocery", "pharmacy", "restaurant"}, []bool{true, true, true, true}))
	// fmt.Println(merge([][]int{{1, 4}}))
	rotate([]int{-1, -100, 3, 99}, 2)
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}

// 29. 两数相除
func divide(dividend int, divisor int) int {
	// 特殊情况：溢出
	if dividend == -1<<31 && divisor == -1 {
		return 1<<31 - 1
	}
	res := 0
	negative := (dividend < 0) != (divisor < 0)
	dvd := abs(dividend)
	dvs := abs(divisor)
	for dvd >= dvs {
		temp := dvs
		mutitime := 1
		for dvd >= (temp << 1) {
			temp <<= 1
			mutitime <<= 1 // 增加了多少倍  就是出出来
		}
		dvd -= temp
		res = res + mutitime
	}
	if negative {
		return -res
	}
	return res
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 3531. 统计被覆盖的建筑
// 输入: n = 3, buildings = [[1,2],[2,2],[3,2],[2,1],[2,3],[6,1]]
func countCoveredBuildings(n int, buildings [][]int) int {
	xmap := make(map[int][]int)
	ymap := make(map[int][]int)
	// min := -1
	// max := n + 1
	for _, v := range buildings {
		xmap[v[1]] = append(xmap[v[1]], v[0]) // y = 1 时 x有哪些
		ymap[v[0]] = append(ymap[v[0]], v[1]) // x = 1 时 y有哪些

	}
	ans := 0
	IsNotMinOrMax := func(nums []int, x int) bool {
		if len(nums) == 0 {
			return false
		}
		minVal, maxVal := nums[0], nums[0]
		// 一次遍历找最小和最大
		for _, v := range nums {
			if v < minVal {
				minVal = v
			}
			if v > maxVal {
				maxVal = v
			}
		}
		return x != minVal && x != maxVal
	}
	for _, v := range buildings {
		x := v[0]
		y := v[1]

		if len(xmap[y]) > 2 && len(ymap[x]) > 2 {
			if IsNotMinOrMax(xmap[y], x) && IsNotMinOrMax(ymap[x], y) {
				ans++
			}
		}
	}
	return ans
}

// 31. 下一个排列
// 1,2,3 -> 1,3,2 -> 2,1,3 -> 2,3,1 -> 3,1,2 -> 3,2,1 -> 1,2,3
// 1,2,3,4 -> 1,2,4,3 -> 1,3,2,4 -> 1,3,4,2 -> 1,4,2,3
// 1,2,3,4 -> 1,2,4,3 -> 1,
func nextPermutation(nums []int) {
	//要全原地转化,不用队列
	// 从右往左找到第一个下降的地方i,交换右边第一个大于他的nums[i]的值
	//重新翻转后续数组
	Reverse := func(nums []int) {
		left, right := 0, len(nums)-1
		for left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			ans := i + 1
			for j := i + 2; j < len(nums); j++ {
				if nums[j] > nums[i] {
					ans = j
				}
			}
			nums[i], nums[ans] = nums[ans], nums[i]
			Reverse(nums[i+1:])
			return
		}
		if i == 0 {
			Reverse(nums[:])
			return
		}
	}
	return
}

// 32. 最长有效括号
// func longestValidParentheses(s string) int {
// 	stack := arraystack.New() // empty
// 	// dp[i] 表示从i结尾最长的有效括号长度
// 	dp := func(s string) int {
// 		if len(s) == 1 || len(s) == 0 {
// 			return 0
// 		}
// 		stack := arraystack.New() // empty
// 	}
// }

//  53. 最大子数组和
//     -2,1,-3,4,-1,2,1,-5,4
//     -2,-1
//
// 前缀和
// 1.一边走一边算总账，然后用现在的总账减去你过去最穷（最小前缀和）时候的账。
// 2.现在的钱减去最穷时的钱，差价越大，说明中间这一段你赚得越多（这就是最大子数组和）。
func maxSubArray(nums []int) int {
	ans := math.MinInt
	minPreSum := 0
	preSum := 0
	for _, x := range nums {
		preSum += x                        // 当前的前缀和
		ans = max(ans, preSum-minPreSum)   // 减去前缀和的最小值
		minPreSum = min(minPreSum, preSum) // 维护前缀和的最小值
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	ans := [][]string{}
	for i := range code {
		temp := true

		for _, v := range code[i] {
			if !(unicode.IsLetter(v) || unicode.IsDigit(v) || v == '_') {
				temp = false
			}
		}
		if len(code[i]) == 0 {
			temp = false
		}
		fmt.Println(i, temp)
		if !(businessLine[i] == "electronics" || businessLine[i] == "grocery" || businessLine[i] == "pharmacy" || businessLine[i] == "restaurant") {
			temp = false
		}
		isActive[i] = temp && isActive[i]
		if isActive[i] {
			ans = append(ans, []string{code[i], businessLine[i]})
		}
	}
	order := map[string]int{
		"electronics": 1,
		"grocery":     2,
		"pharmacy":    3,
		"restaurant":  4,
	}
	sort.Slice(ans, func(i, j int) bool {
		if order[ans[i][1]] != order[ans[j][1]] {
			return order[ans[i][1]] < order[ans[j][1]]
		}
		// 如果业务线相同，再比较 code
		return ans[i][0] < ans[j][0]
	})
	ansa := []string{}
	for i := range ans {
		ansa = append(ansa, ans[i][0])
	}
	return ansa
}

// 56. 合并区间
func merge(intervals [][]int) [][]int {
	ans := [][]int{}
	//先按照第一个参数进行排序,依次计算
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1] // 起点相同，按终点排序
		}
		return intervals[i][0] < intervals[j][0] // 按起点排序
	})
	var bg int
	var end int
	for i := range intervals {
		if i == 0 {
			bg = intervals[i][0]
			end = intervals[i][1]
			if i == len(intervals)-1 {
				ans = append(ans, []int{bg, end})
			}
			continue
		}
		if intervals[i][0] <= end {
			end = max(end, intervals[i][1])
		} else {
			ans = append(ans, []int{bg, end})
			bg, end = intervals[i][0], intervals[i][1]

		}
		if i == len(intervals)-1 {
			ans = append(ans, []int{bg, end})
		}
	}
	return ans
}

// 189. 轮转数组
func rotate(nums []int, k int) {
	k = k % len(nums)
	left := append([]int(nil), nums[len(nums)-k:]...)
	right := append([]int(nil), nums[:len(nums)-k]...)
	// fmt.Println(left, right, nums)
	for i := range nums {
		if i < k {
			nums[i] = left[i]
		} else {
			nums[i] = right[i-k]
		}
	}
	fmt.Println(left, right, nums)
}
