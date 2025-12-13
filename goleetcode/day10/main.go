package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// minWindow("", "")
	// specialTriplets([]int{6, 3, 6})
	// specialTriplets([]int{0, 1, 0, 0})
	// specialTriplets([]int{8, 4, 2, 8, 4})
	fmt.Println(countPermutations([]int{3, 2, 3, 1}))
}

func minWindow(s string, t string) string {
	tint := make([]int, 128)
	sint := make([]int, 128)

	for _, v := range t {
		tint[v-'a']++
		sint[v-'a']++
	}
	ismatch := func(small, large []int) bool {
		for k := range small {
			if small[k] > large[k] {
				return false
			}
		}
		return true
	}
	start, end := 0, 0
	anss, anse := 0, 0
	ans := len(s) + 1
	for end < len(s) {
		sint[s[end]-'a']++
		for ismatch(tint, sint) {
			ans = min(ans, end-start+1)
			anss = start
			anse = end
			sint[s[start]-'a']--
			start++
		}
		end++
	}
	if ans == len(s)+1 {
		return ""
	}
	return s[anss : anse+1]

}

// 209. 长度最小的子数组
// 前缀和解法
//
//	2 3 1 2 4 3
//
// 0 2 5 6 8 12 15
func minSubArrayLen(target int, nums []int) int {
	presum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		presum[i] = presum[i-1] + nums[i-1]
	}

	minLen := len(nums) + 1
	for left := 1; left <= len(nums); left++ {
		// 找到最小的 j，使得 presum[j] >= target + presum[left]
		// 找到第一个满足条件的j
		j := sort.SearchInts(presum, target+presum[left-1])
		if j <= len(nums) {
			minLen = min(minLen, j-(left-1))
		}
	}

	if minLen == len(nums)+1 {
		return 0 // 没有满足条件的子数组
	}
	return minLen
}

// 双指针解法
//
//	2 3 1 2 4 3
func minSubArrayLenPoint(target int, nums []int) int {
	n := len(nums)
	start, end := 0, 0
	sum := 0
	ans := math.MaxInt32
	for end < n {
		sum += nums[end]
		for sum >= target {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 3583. 统计特殊三元组
func specialTriplets(nums []int) int {
	mod := 1000000007
	ans := 0
	ml := make(map[int]int)
	mr := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mr[nums[i]]++
	}
	mr[nums[0]]--
	for i := 1; i < len(nums)-1; i++ {
		mr[nums[i]]--
		ml[nums[i-1]]++
		// fmt.Println("i,ml,mr", nums[i], ml, mr)
		ans += (ml[nums[i]*2] * mr[nums[i]*2]) % mod
		ans = ans % mod
	}
	// fmt.Println(ans)
	return ans
}

// 53. 最大子数组和
// func maxSubArray(nums []int) int {

// }
const MOD = 1000000007

// 3577. 统计计算机解锁顺序排列数
func countPermutations(complexity []int) int {
	start := complexity[0]
	sort.Ints(complexity)
	if complexity[0] == complexity[1] || complexity[0] != start {
		return 0
	}
	// 全排列
	// A2 1 = 2*1
	// An-1 n-1
	ans := 1
	n := len(complexity)
	for i := 1; i < n; i++ {
		ans *= i
		ans = ans % MOD
	}
	return ans
}
