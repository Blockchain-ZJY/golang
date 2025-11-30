package main

import (
	"fmt"
	"math"
	"sort"
)

//// 3512.使数组和能被K 整除的最少操作次数

func minOperations(nums []int, k int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	// if sum < k {
	// 	return sum
	// }
	return sum % k
}

// 12. 整数转罗马
func intToRoman(num int) string {
	THOUSANDS := []string{"", "M", "MM", "MMM"}
	HUNDREDS := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	TENS := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ONES := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	ans := ""
	ans = ans + THOUSANDS[num/1000] + HUNDREDS[num/100%10] + TENS[num/10%10] + ONES[num%10]
	return ans
}

// 12. 罗马转整数
func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	n := len(s)
	ans := m[s[n-1]]
	for i := n - 2; i >= 0; i-- {
		if m[s[i]] >= m[s[i+1]] {
			ans += m[s[i]]
		} else {
			ans -= m[s[i]]
		}
	}
	return ans
}
func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

// 14.最长公共前缀
func longestCommonPrefix(strs []string) string {
	strings := [][]byte{}
	if len(strs) == 1 {
		return strs[0]
	}
	for _, v := range strs {
		if v == "" {
			return ""
		}
	}
	minindex := 111110
	for _, v := range strs {
		minindex = min(minindex, len(v))
	}
	fmt.Println("1231", minindex)
	for _, v := range strs {
		strings = append(strings, []byte(v))
	}
	fmt.Println(strings)
	var index int
	for index = 0; index < minindex; index++ {
		for i := 0; i < len(strings)-1; i++ {
			if strings[i][index] != strings[i+1][index] {
				if index == 0 {
					return ""
				} else {
					return string([]byte(strs[0])[:index])
				}
			}
		}
	}
	return string([]byte(strs[0])[:index])

}
func longestCommonPrefix_rebuild(strs []string) string {
	firstone := strs[0]
	maxi := 0
	for i := 0; i < len(firstone); i++ {
		for _, eachstring := range strs[1:] {
			if i > len(eachstring)-1 || firstone[i] != eachstring[i] {
				maxi = i
				return firstone[:maxi]
			}
		}
	}

	return strs[0]
}
func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	pivot := nums[left]
	i, j := left, right

	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		for i < j && nums[i] <= pivot {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[left], nums[i] = nums[i], nums[left]

	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

// 15 三数之和
func threeSum(nums []int) [][]int {
	// 双指针
	ans := [][]int{}
	newnums := sortArray(nums)
	fmt.Println(newnums)
	for i := 0; i < len(newnums)-2; i++ {
		if newnums[i] > 0 {
			return ans
		}
		if i > 0 && newnums[i] == newnums[i-1] {
			continue
		}
		left, right := i+1, len(newnums)-1
		for right > left {
			if newnums[left]+newnums[right]+newnums[i] > 0 {
				right--
			} else if newnums[left]+newnums[right]+newnums[i] < 0 {
				left++
			} else {
				ans = append(ans, []int{newnums[i], newnums[left], newnums[right]})
				left++
				right--
				for left < right && newnums[right] == newnums[right+1] {
					right--
				}
				for left < right && newnums[left] == newnums[left-1] {
					left++
				}
				fmt.Println(ans)
			}
		}
	}
	return ans
}
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

// 1590. 使数组和能被 P 整除;
// 给你一个正整数数组 nums，请你移除 最短 子数组（可以为 空），使得剩余元素的 和 能被 p 整除。 不允许 将整个数组都移除。
// 请你返回你需要移除的最短子数组的长度，如果无法满足题目要求，返回 -1 。
// 子数组 定义为原数组中连续的一组元素。

// 去掉最少的子串,使得整个数组和能被p整除
// 分析: 前缀和+mod+hash
// X X X X [i X X j] X X X X
//
// sum[i:j] = prefix[j] - prefix[i-1]
// sum[i:j]%p === k === prefix[j]%p - prefix[i-1]%p
// 余数  k				 确定j     ->    确定i的位置
// 确定一个j 找到一个i使得prefix[i-1]%p = 定值
// presum[j]-presum[x]%p == k
func minSubarray(nums []int, p int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	k := total % p
	if k == 0 {
		return 0
	}

	n := len(nums)
	ans := n
	s := 0
	last := map[int]int{0: -1}

	for i, v := range nums {
		s += v
		mod := s % p
		last[mod] = i
		target := (mod - k + p) % p
		if j, ok := last[target]; ok {
			ans = min(ans, i-j)
		}
	}

	if ans == n {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}

func minSubarray1(nums []int, p int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	k := total % p //
	if k == 0 {
		return 0
	}
	maps := make(map[int]int)
	maps[0] = -1
	n := 0
	ans := len(nums)
	for index, v := range nums {
		n += v
		mod := n % p
		if j, ok := maps[(mod-k+p)%p]; ok {
			ans = min(ans, index-j)
		}
		maps[mod] = index
	}
	if ans == n {
		return -1
	}
	return ans
}

// 16. 最接近的三数之和
func threeSumClosest(nums []int, target int) int {
	// 	X X X X X X X X X
	// 	i l r
	sort.Ints(nums) // sorts in place
	fmt.Println(nums)
	n := len(nums)
	left, right := 0, 0
	threesum := 0
	cloestnum := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		left, right = i+1, n-1
		for left < right {
			threesum = nums[i] + nums[left] + nums[right]
			if math.Abs(float64(threesum-target)) < math.Abs(float64(cloestnum-target)) {
				cloestnum = threesum
			}
			if threesum-target == 0 {
				return target
			} else if threesum < target {
				left++
			} else {
				right--
			}
		}
	}
	return cloestnum
}
