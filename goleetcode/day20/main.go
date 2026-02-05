package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	// a := []byte("AAAA/BBBBB")
	// fmt.Println(cap(a), len(a))
	// index := bytes.IndexByte(a, '/')
	// b := a[:index]
	// fmt.Println(len(b), "len(b)")
	// c := a[index+1:]
	// b = append(b, "CCCCCCC"...)
	// fmt.Println(len(b), "len(b)")
	// fmt.Println(string(a))
	// fmt.Println(string(b))
	// fmt.Println(string(c))

	fmt.Println(threeSum([]int{1, 2, 0, 1, 0, 0, 0, 0}))
}

// 15. 三数之和
func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for nums[l] == nums[l-1] && l+1 < len(nums) {
					l++
				}
				for nums[r] == nums[r+1] && r-1 > 0 {
					r--
				}
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return
}

// 打家劫舍
func rob(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	if len(nums) == 1 {
		return dp[0]
	}
	dp[1] = max(dp[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	fmt.Println(dp)
	return dp[len(nums)-1]
}

// 118. 杨辉三角
func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		ans[i] = make([]int, i+1)
	}
	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				ans[i][j] = 1
			} else {
				fmt.Println(i, j)
				ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
			}
		}
	}
	return ans
}

// 70. 爬楼梯
func climbStairs(n int) int {
	dp := make([]int, n+2)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func constructTransformedArray(nums []int) []int {
	ans := make([]int, len(nums))
	m := len(nums)
	for i := range nums {
		ans[i] = nums[((nums[i]+i)%m+m)%m]
	}
	return ans
}

func singleNumber(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans ^= v
	}
	return ans
}

// 169. 多数元素
func majorityElement(nums []int) int {
	target := nums[0]
	hp := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == target {
			hp++
		} else {
			hp--
			if hp == 0 {
				target = nums[i]
				hp++
			}
		}
	}
	return target
}

// 75. 颜色分类
func sortColors(nums []int) {
	l, r, n := 0, 0, len(nums)
	for r < n {
		if nums[r] == 0 {
			nums[r], nums[l] = nums[l], nums[r]
			l++
		}
		r++
	}
	r = l
	for r < n {
		if nums[r] == 1 {
			nums[r], nums[l] = nums[l], nums[r]
			l++
		}
		r++
	}
}

// 31. 下一个排列
func nextPermutation(nums []int) {
	// 找到第一个下降的所以 i
	// i的后面,找到第一个大于nums[i]的坐标j,交换
	// 翻转后续的数组
	for i := len(nums) - 1; i >= 0; i-- {
		if i != len(nums)-1 && nums[i] < nums[i+1] {
			j := len(nums) - 1
			for j >= i && nums[j] <= nums[i] {
				j--
			}
			fmt.Println(nums[i], nums[j])
			nums[i], nums[j] = nums[j], nums[i]
			slices.Reverse(nums[i+1:])
			return
		}
	}
	slices.Reverse(nums)
}

// 寻找重复数
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for nums[slow] != nums[nums[fast]] {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	head := 0
	for slow != head {
		slow = nums[slow]
		head = nums[head]
	}
	return slow
}
