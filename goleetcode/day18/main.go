package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(generate(5))
}

// 61. 旋转链表
func rotateRight(head *ListNode, k int) (newhead *ListNode) {
	n := 1
	cur := head
	for cur != nil {
		n++
		if cur.Next == nil {
			cur.Next = head
			break
		}
		cur = cur.Next
	}
	fmt.Println(n)
	k = k % n
	for i := 0; i < n-k; i++ {
		if i == n-k-1 {
			newhead = head.Next
			head.Next = nil
			break
		}
		head = head.Next
	}
	return
}

// 198. 打家劫舍
func rob(nums []int) int {
	dp := make([]int, len(nums)+1)
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	} else {
		dp[0], dp[1] = nums[0], max(nums[0], nums[1])
		for i := 3; i <= len(nums); i++ {
			dp[i] = max(dp[i-2]+nums[i], dp[i-1])
		}
	}
	return dp[len(nums)]
}

// 杨辉三角
func generate(numRows int) [][]int {
	dp := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		dp[i] = make([]int, i+1)
	}
	fmt.Println(dp)
	for i := 0; i < numRows; i++ {
		dp[i][0] = 1
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			fmt.Println(dp[i][j])
		}
		dp[i][i] = 1
	}
	return dp
}

// 287. 寻找重复数
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if fast == slow {
			break
		}
	}
	head := 0
	for slow != head {
		slow = nums[slow]
		head = nums[head]
	}
	return slow
}

// 75. 颜色分类
// 把1 2 看成一个东西
// 在从后面分
func sortColors(nums []int) {
	n := len(nums)
	j := 0 // 当前需要交换的地方
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	for i := j; i < n; i++ {
		if nums[i] == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	for _, p := range prerequisites {
		g[p[1]] = append(g[p[1]], p[0])
	}

	colors := make([]int, numCourses)
	// 返回 true 表示找到了环
	var dfs func(int) bool
	dfs = func(x int) bool {
		colors[x] = 1 // x 正在访问中
		for _, y := range g[x] {
			// 情况一：colors[y] == 1，表示发生循环依赖，找到了环
			// 情况二：colors[y] == 0，未知，继续递归 y 获取信息
			// 情况三：colors[y] == 2，继续递归 y 只会重蹈覆辙，和之前一样无法找到环
			if colors[y] == 1 || colors[y] == 0 && dfs(y) {
				return true // 找到了环
			}
		}
		colors[x] = 2 // x 完全访问完毕，从 x 出发无法找到环
		return false  // 没有找到环
	}

	for i, c := range colors {
		if c == 0 && dfs(i) {
			return false // 有环
		}
	}
	return true // 没有环
}

type Node struct {
	Val, Key  int
	Next, Pre *Node
}
type LRUCache struct {
	size      int
	dummy     *Node
	KeytoNode map[int]*Node
}

func Constructor(capacity int) LRUCache {
	dummy := &Node{}
	dummy.Next = dummy
	dummy.Pre = dummy
	m := make(map[int]*Node)
	return LRUCache{
		size:      capacity,
		KeytoNode: m,
		dummy:     dummy,
	}
}

func (l *LRUCache) Remove(x *Node) {
	x.Pre.Next = x.Next
	x.Next.Pre = x.Pre
}

// 在链表头加上一个节点x
func (l *LRUCache) PushFront(x *Node) {
	x.Pre = l.dummy
	x.Next = l.dummy.Next
	x.Pre.Next = x
	x.Next.Pre = x
}

func (l *LRUCache) Get(key int) int {
	node := l.KeytoNode[key]
	// 不存在该值
	if node == nil {
		return -1
	}
	l.Remove(node)
	l.PushFront(node)
	return node.Val
}

func (l *LRUCache) Put(key int, value int) {
	node := l.KeytoNode[key]
	if node != nil {
		node.Val = value
		l.Remove(node)
		l.PushFront(node)
		return
	}
	// 当前节点为空
	newnode := &Node{
		Val: value,
		Key: key,
	}

	l.PushFront(newnode)
	l.KeytoNode[key] = newnode
	if l.size < len(l.KeytoNode) {
		backnode := l.dummy.Pre
		delete(l.KeytoNode, backnode.Key)
		l.Remove(l.dummy.Pre)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

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

// func Constructor() MedianFinder {
// 	smheap := &smh{}
// 	bigh := &bigh{}
// 	heap.Init(smheap)
// 	heap.Init(bigh)
// 	return MedianFinder{
// 		sum: 0,
// 		s:   smheap,
// 		b:   bigh, // 大根堆放左半边小数
// 	}
// }

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
