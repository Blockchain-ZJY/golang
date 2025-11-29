package main

import (
	"fmt"
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

func main() {
	fmt.Println(threeSum([]int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10}))
}
