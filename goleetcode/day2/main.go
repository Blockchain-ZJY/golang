package main

import "fmt"

// 3. 无重复字符的最长子串
type Substring struct {
	Start   int
	End     int
	Letters string
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	maxlen := 1
	for i := 0; i < len(s); i++ {
		// 每一个开头，都从下一个字符开始，记录最大长度
		maxfori := 1
		mapi := make(map[string]bool)
		mapi[string(s[i])] = true
		for j := i + 1; j < len(s); j++ {
			if j == len(s)-1 && !mapi[string(s[j])] {
				maxfori = j - i + 1
				break
			}
			if _, ok := mapi[string(s[j])]; ok { // 能获取到数据，有重复
				maxfori = j - i
				break
			}
			mapi[string(s[j])] = true
		}
		if maxfori > maxlen {
			maxlen = maxfori
		}
	}
	return maxlen
}

// better way
func lengthOfLongestSubstringBetter(s string) (ans int) { //      pw
	if len(s) == 0 {
		return 0
	}
	count := 0
	left := 0 //      l r             [l, r)
	right := 0
	mapi := make(map[byte]int)
	maxlen := 1
	for right < len(s) {
		newone := s[right]
		right++
		mapi[newone]++
		if mapi[newone] == 2 {
			maxlen = Max(maxlen, right-left-1)
			fmt.Println(maxlen, count)
			count++
			for {
				left++
				if mapi[s[left-1]] > 1 {
					mapi[s[left-1]]--
					break
				} else {
					delete(mapi, s[left-1])
				}
			}
		} else {
			maxlen = Max(maxlen, right-left)
		}
	}
	return maxlen
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 4. 寻找两个正序数组的中位数

/*12321

将数组上下摆放，需要找到 上面的数组i，下面数组j的位置，
使得左边的所有元素都小于右边的元素
同时由于长度固定，i+j必须满足 i+j = (m+n+1)/2

*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	ans := 0.0
	n := len(nums1)
	m := len(nums2)
	var longer, shorter []int
	if n < m {
		longer, shorter = nums2, nums1
	} else {
		longer, shorter = nums1, nums2
	}
	halflen := (m + n + 1) / 2
	low := 0
	high := len(shorter)
	// 短数组内进行二分查找（时间复杂度要求）
	for low <= high {
		i := (low + high) / 2 // i = 2 j = 6-2 = 4
		j := halflen - i
		if getVal(shorter, i-1) <= getVal(longer, j) && getVal(longer, j-1) <= getVal(shorter, i) {
			maxLeft := float64(Max(getVal(shorter, i-1), getVal(longer, j-1)))
			minRight := float64(Min(getVal(shorter, i), getVal(longer, j)))
			if (m+n)%2 == 0 { // 偶数
				ans = (maxLeft + minRight) / 2.0
			} else { // 奇数个
				ans = maxLeft
			}
			break
		}
		if getVal(shorter, i) < getVal(longer, j-1) {
			low = i + 1
		} else {
			high = i - 1
		}
	}
	return ans
}

func getVal(arr []int, i int) int {
	if i < 0 {
		return -1
	}
	if i > len(arr)-1 {
		return 1000000000
	}
	return arr[i]
}

// 5. 最长回文子串
func longestPalindrome(s string) string {
	ans := ""
	maxlen := 1
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			// fmt.Println(s[i : j+1])
			if isPalindrome(s[i : j+1]) {
				if j-i+1 >= maxlen {
					maxlen = j - i + 1
					ans = s[i : j+1]
				}
			}
		}
	}
	return ans
}
func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(lengthOfLongestSubstringBetter("ab"))
}
