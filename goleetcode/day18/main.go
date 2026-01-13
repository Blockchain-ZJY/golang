package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

// 53. 最大子数组和
// dp[i] 表示以i结尾的最大连续子数组和
// 对于i这个位置, 可以入股,也可以单干 保障自己最大
func maxSubArray(nums []int) (ans int) {
	dp := make([]int, len(nums)+2)
	dp[0] = nums[0]
	ans = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(dp[i], ans)
	}
	fmt.Println(dp)
	return
}

// 152. 乘积最大子数组
// dp[i] 表示以i为结尾的最大子数组乘积
func maxProduct(nums []int) int {
	ans, preMax, preMin := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		a := preMax * nums[i]
		b := preMin * nums[i]
		preMax = max(nums[i], max(a, b))
		preMin = min(nums[i], min(a, b))
		ans = max(preMax, ans)
	}
	return ans
}

// dp[i] 表示和为 i 的完全平方数的最少数量
// 279. 完全平方数
func numSquaresTest(n int) int {

	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = 1000000
	}
	dp[0] = 0
	dp[1] = 1
	if n == 2 {
		return 2
	}
	dp[2] = 2

	//对于一个数 i 能被很多小于i的去减
	for i := 3; i <= n; i++ {
		for j := 0; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	fmt.Println(dp)
	return dp[n]
}

// 198. 打家劫舍
// dp[i] 表示到i所拿到最大的金额
func rob(nums []int) int {
	dp := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = nums[0]
		} else if i == 1 {
			dp[i] = max(nums[1], nums[0])
		} else {
			dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		}
	}
	fmt.Println(dp)
	return dp[len(nums)-1]
}

// 139. 单词拆分
// dp[i] 表示以 0-i 的字符串是否能被表示
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i < len(s); i++ {
		if !dp[i] {
			continue
		}
		for _, word := range wordDict {
			end := i + len(word)
			if end <= len(s) && s[i:end] == word {
				dp[end] = true
			}
		}
	}
	return dp[len(s)]
}

// ababcbaca defegdehijhklij
// 双指针实现
func partitionLabels(s string) (ans []int) {
	n := len(s)
	if n == 0 {
		return
	}

	// 统计每个字符总出现次数
	right := make(map[byte]int)
	for i := 0; i < n; i++ {
		right[s[i]]++
	}
	left := make(map[byte]int)
	start := 0
	for i := 0; i < n; i++ {
		c := s[i]
		left[c]++
		right[c]--
		// 判断当前 left 中的字符是否都不再出现在 right 中
		canCut := true
		for k := range left {
			if right[k] > 0 {
				canCut = false
				break
			}
		}
		if canCut {
			ans = append(ans, i-start+1)
			left = make(map[byte]int)
			start = i + 1
		}
	}
	return
}

// ababcbaca defegdehijhklij
// 双指针实现
func partitionLabelsGreedy(s string) (ans []int) {
	n := len(s)
	if n == 0 {
		return
	}
	maxri := make(map[byte]int)
	for i := 0; i < n; i++ {
		maxri[s[i]] = i
	}
	l := 0
	maxpoi := 0
	for i := 0; i < n; i++ {
		maxpoi = max(maxpoi, maxri[s[i]])
		if i == maxpoi {
			ans = append(ans, i-l+1)
			l = i + 1
		}
	}
	return
}

// 121. 买卖股票的最佳时机
func maxProfit(prices []int) int {
	// 用map记录前i个元素最小值
	m := make(map[int]int)
	minprs := prices[0]
	for i := 0; i < len(prices); i++ {
		minprs = min(minprs, prices[i])
		m[i] = minprs
	}
	maxpro := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-m[i-1] > 0 {
			maxpro = max(maxpro, prices[i]-m[i-1])
		}
	}
	return maxpro
}

// 45. 跳跃游戏 II
// dp[i] 表示跳到i的位置最小的次数
// 往前推
func jump(nums []int) int {
	dp := make([]int, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for i := 1; i < len(nums)+1; i++ {
		for j := i - 1; j >= 0; j-- {
			if i-j <= nums[j] {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(nums)-1]
}

// 45. 跳跃游戏 II
// 维护一个[l,r]表示的当前轮次能达到的最近最远距离
// 往前推直到[l,r]包含了最后一个位置
func jumpBFS(nums []int) int {
	l, r := 0, 0
	ans := 0
	for r < len(nums)-1 {
		maxr := 0
		for i := l; i <= r; i++ {
			maxr = max(maxr, i+nums[i])
		}
		ans++
		l = r + 1
		r = maxr
	}
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummy := &ListNode{}
	cur := dummy
	for carry != 0 || l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		fmt.Println(sum)
		cur.Next = &ListNode{
			Val: sum % 10,
		}
		carry = sum / 10
	}
	return dummy.Next
}
func twoSum(nums []int, target int) (ans []int) {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[target-nums[i]]; ok {
			ans = append(ans, i, v)
			return
		}
		m[nums[i]] = i
	}
	return
}

func findPairsHahs(nums []int, k int) (ans int) {
	//用值作为k 统计其频率
	fre := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		fre[nums[i]]++
	}
	if k == 0 {
		for _, v := range fre {
			if v >= 2 {
				ans++
			}
		}
	} else {
		for x := range fre {
			if _, ok := fre[x+k]; ok {
				ans++
			}
		}
	}
	return
}

// 532. 数组中的 k-diff 数对 双指针
func findPairs(nums []int, k int) (ans int) {
	sort.Ints(nums)
	// fmt.Println(nums)
	l, r := 0, 1
	for r < len(nums) {
		fmt.Println(l, r)
		dif := nums[r] - nums[l]
		if dif == k {
			ans++
			r++
			for r < len(nums) && r > 0 && nums[r] == nums[r-1] {
				r++
			}
		} else if dif < k {
			r++
		} else {
			l++
		}
		if l == r {
			r++
		}
	}
	return ans
}

// dp[i] 表示兑换 i 所需要的最小硬币数
// dp[i] = min(dp[i],dp[i-j]+1) j属于 coins 索引 []int{1, 2, 5}, 11
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i-c >= 0 && dp[i-c] != math.MaxInt {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}

// 279. 完全平方数
// dp[i] 表示和为 i 的完全平方数的最少数量
// dp[i] = min(dp[i],dp[i-j*j]+1)
func numSquares(n int) int {
	dp := make([]int, n+2)
	for i := 0; i <= n+1; i++ {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n+1; i++ {
		for j := 1; j*j <= i; j++ {
			if i == 13 {
				fmt.Println(dp[i], j)
			}
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

type smh []int  // 最小堆（右半边）
type bigh []int // 最大堆（左半边）
// 最小堆实现
func (s smh) Len() int           { return len(s) }
func (s smh) Less(i, j int) bool { return s[i] < s[j] }
func (s smh) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s *smh) Push(x any)        { *s = append(*s, x.(int)) }
func (s *smh) Pop() any {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

// 最大堆实现
func (b bigh) Len() int           { return len(b) }
func (b bigh) Less(i, j int) bool { return b[i] > b[j] }
func (b bigh) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b *bigh) Push(x any)        { *b = append(*b, x.(int)) }
func (b *bigh) Pop() any {
	old := *b
	n := len(old)
	x := old[n-1]
	*b = old[:n-1]
	return x
}

type MedianFinder struct {
	sum int
	s   *smh
	b   *bigh
}

func Constructor() MedianFinder {
	smheap := &smh{}
	bigh := &bigh{}
	heap.Init(smheap)
	heap.Init(bigh)
	return MedianFinder{
		sum: 0,
		s:   smheap,
		b:   bigh, // 大根堆放左半边小数
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.sum++
	if this.b.Len() == 0 || num <= (*this.b)[0] {
		heap.Push(this.b, num)
	} else {
		heap.Push(this.s, num)
	}
	// 平衡两个堆
	if this.b.Len() > this.s.Len()+1 {
		heap.Push(this.s, heap.Pop(this.b))
	} else if this.s.Len() > this.b.Len() {
		heap.Push(this.b, heap.Pop(this.s))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.sum%2 == 0 {
		return (float64((*this.b)[0]) + float64((*this.s)[0])) / 2.0
	}
	return float64((*this.b)[0])
}
