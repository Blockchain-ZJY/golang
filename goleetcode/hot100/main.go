package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}

func rotate(nums []int, k int) {
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		for nums[i] <= n && nums[i] > 0 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	// 找第一个不匹配的位置
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func productExceptSelf(nums []int) []int {
	length := len(nums)

	// L 和 R 分别表示左右两侧的乘积列表
	L, R, answer := make([]int, length), make([]int, length), make([]int, length)

	L[0] = 1
	for i := 1; i < length; i++ {
		L[i] = nums[i-1] * L[i-1]
	}

	R[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		R[i] = nums[i+1] * R[i+1]
	}

	// 对于索引 i，除 nums[i] 之外其余各元素的乘积就是左侧所有元素的乘积乘以右侧所有元素的乘积
	for i := 0; i < length; i++ {
		answer[i] = L[i] * R[i]
	}
	return answer
}

// 1  2  3 4
// 1  1  2 6
// 12 12 6 1
func merge(intervals [][]int) (ans [][]int) {
	// 排序
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	for i := 0; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]
		// 合并区间：下一个区间的起点 <= 当前区间的终点
		for i+1 < len(intervals) && intervals[i+1][0] <= end {
			end = max(end, intervals[i+1][1])
			i++
		}
		ans = append(ans, []int{start, end})
	}

	return ans
}

// dp[i] 表示以i结尾最大子数组和最大和
func maxSubArray(nums []int) (ans int) {
	ans = math.MinInt32
	dp := make([]int, len(nums)+2)
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(ans, dp[i])
	}
	return
}
func minWindow(s string, t string) string {
	tgt := make(map[byte]int, 0)
	src := make(map[byte]int, 0)
	for i := range t {
		tgt[t[i]]++
	}
	isok := func(src map[byte]int, tgt map[byte]int) bool {
		for i := range tgt {
			if src[i] < tgt[i] {
				return false
			}
		}
		return true
	}
	ansl, ansr := 0, math.MaxInt32
	l, r := 0, 0
	for r < len(s) {
		src[s[r]]++
		r++
		fmt.Println(src)
		for isok(src, tgt) {
			if (ansr - ansl) > r-l {
				ansr = r
				ansl = l
			}
			src[s[l]]--
			l++
		}
	}
	if ansl == ansr {
		return ""
	}
	return s[ansl:ansr]
}

// 560. 和为 K 的子数组
func subarraySum(nums []int, k int) (ans int) {
	sum := 0
	m := make(map[int]int) // key :前缀和, v: 当前和为当前前缀和的总数
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		ans += m[sum-k]
		m[sum]++
	}
	return
}
