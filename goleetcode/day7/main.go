package main

import (
	"fmt"
	"math"
	"sort"
)

const MOD = 1000000007

// 3623. 统计梯形的数目 I
func countTrapezoids(points [][]int) int {
	maps := make(map[int]int)
	for _, v := range points {
		maps[v[1]]++
	}
	pos := []int{}
	for _, v := range maps {
		pos = append(pos, (v%MOD*(v-1)%MOD/2)%MOD)
	}
	res := 0
	for i := 0; i < len(pos)-1; i++ {
		for j := i + 1; j < len(pos); j++ {
			res += (pos[i] * pos[j]) % MOD
		}
	}
	return res % MOD
}

// 3623. 统计梯形的数目 I
func countTrapezoids1(points [][]int) int {
	maps := make(map[int]int)
	for _, v := range points {
		maps[v[1]]++
	}
	pos := []int{}
	for _, v := range maps {
		pos = append(pos, (v*(v-1)/2)%MOD)
	}
	totaledges := 0
	for _, v := range pos {
		totaledges += v
	}
	res := 0
	for _, cur := range pos {
		res += (cur * (totaledges - cur)) % MOD
	}
	res = res * ((MOD + 1) / 2) % MOD
	return res
}

func main() {
	// fmt.Println(groupAnagrams1([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	// fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	// fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	// fmt.Println(longestConsecutive([]int{1, 0, 1, 2}))
	// fmt.Println(longestConsecutive([]int{0, 0}))
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{0, 1})
	moveZeroes([]int{1, 0, 0})
	moveZeroes([]int{1, 0, 1})

}

// 49. 字母异位词分组
// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for k, v := range strs {
		//string can convert to []byte
		b := []byte(v)
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		key := string(b)
		groups[key] = append(groups[key], v)
		fmt.Println(k, b, key)
	}
	ans := [][]string{}
	for _, v := range groups {
		ans = append(ans, v)
	}
	return ans
}

func groupAnagrams1(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		// 统计字母频率
		count := make([]int, 26)
		for _, ch := range s {
			count[ch-'a']++
		}
		key := fmt.Sprint(count)
		fmt.Println(key)
		groups[key] = append(groups[key], s)
	}

	ans := [][]string{}
	for _, v := range groups {
		ans = append(ans, v)
	}
	return ans
}

// 128. 最长连续序列
func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	fmt.Println(nums)
	if len(nums) <= 1 {
		return len(nums)
	}
	ans := 0
	start, r := 0, 1
	count := 0
	for r < len(nums) {
		if nums[r] == nums[r-1] {
			r++
			count++
			continue
		}
		if nums[r-1] == nums[r]-1 {
			r++
		} else {
			// fmt.Println(count)
			ans = int(math.Max(float64(ans), float64(r-start-count)))
			count = 0
			start = r
			r++
		}
	}

	return int(math.Max(float64(ans), float64(r-start-count)))
}

// 283. 移动零
// [1,0,1,0,3,12]
func moveZeroes(nums []int) {

	l, r := 0, 0
	for l < len(nums)-1 && r < len(nums)-1 {
		for nums[l] != 0 && l < len(nums)-1 {
			l++
		}
		for nums[r] == 0 && r < len(nums)-1 {
			r++
		}
		fmt.Println(l, r)
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		} else {
			r++
		}

	}
	fmt.Println(nums)
}

//

func moveZeroes1(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
