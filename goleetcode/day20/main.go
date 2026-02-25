package main

import (
	"bytes"
	"cmp"
	"fmt"
	"math/bits"
	"slices"
	"sort"
)

func main() {

	a := []byte("AAAA/BBBBB")
	b := append(a, 'C')
	fmt.Println(string(a), len(a), cap(a), "len(a) cap(a)")
	fmt.Println(string(b), len(b), cap(b), "len(b) cap(b)")
	index := bytes.IndexByte(a, '/')
	fmt.Println(string(b), "string(b)")
	c := a[index+1:]
	b = append(b, "CCCCCCC"...)
	fmt.Println(len(b), cap(b), "len(b)")

	fmt.Println(string(c))
}

func groupAnagrams(strs []string) (ans [][]string) {
	m := make(map[[26]int][]string)
	for i := range strs {
		tep := [26]int{}
		for j := 0; j < len(strs[i]); j++ {
			tep[strs[i][j]-'a']++
		}
		m[tep] = append(m[tep], strs[i])
	}
	for _, v := range m {
		ans = append(ans, v)
	}
	return
}

func twoSum(nums []int, target int) (ans []int) {
	m := make(map[int]int)
	m[nums[0]] = 0
	for i := 1; i < len(nums); i++ {
		index, ok := m[target-nums[i]]
		if ok {
			ans = append(ans, i, index)
			return
		}
		m[nums[i]] = i
	}
	return
}

// 1356. 根据数字二进制下 1 的数目排序
func sortByBits(arr []int) []int {
	slices.SortFunc(arr, func(a, b int) int {
		return cmp.Or(bits.OnesCount(uint(a))-bits.OnesCount(uint(b)), a-b)
	})
	return arr
}

// 3634. 使数组平衡的最少移除数目
func minRemoval(nums []int, k int) (ans int) {
	maxlen := 1
	sort.Ints(nums)
	l, r := 0, 1
	for r < len(nums) {
		// 找到最长的平衡子数组
		if nums[l]*k >= nums[r] {
			r++
			maxlen = max(maxlen, r-l)
		} else {
			l++
		}
	}
	ans = len(nums) - maxlen
	return
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
