package main

import (
	"container/heap"
	"fmt"
	"slices"
	"sort"
)

func main() {
	// ans := longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	// ans := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	// fmt.Println(threeSum([]int{0, 1, 1}))
	// fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
	// fmt.Println(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4))
	fmt.Println(bs([]int{1, 3, 5, 7, 9, 13, 16, 27}, 16))
}

// 二分查找
func bs(nums []int, target int) int {
	l, r := 0, len(nums)
	mid := (l + r) / 2
	for l < r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return -1
}

// 1 2 3 4 5
// 1 3 6 10 15
// k = 5
func addnum(nums []int, x int) []int {
	for len(nums) > 0 && x > nums[len(nums)-1] {
		nums = nums[:len(nums)-1]
	}
	nums = append(nums, x)
	return nums
}

func maxSlidingWindow(nums []int, k int) (ans []int) {

	st := []int{}
	for i := 0; i < k; i++ {
		st = addnum(st, nums[i])
	}
	ans = append(ans, st[0])
	for i := k; i < len(nums); i++ {
		st = addnum(st, nums[i])
		if nums[i-k] == st[0] {
			st = st[1:]
		}
		ans = append(ans, st[0])
	}
	return
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	s, f := head, head.Next
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	x := reverseList(s.Next)
	fmt.Println(x.Val, head.Val)
	for x != nil {
		if x.Val != head.Val {
			return false
		}
		x = x.Next
		head = head.Next
	}
	return true
}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	c := 0
	for a != nil && b != nil {
		a = a.Next
		b = b.Next
		if a == nil {
			for b != nil {
				b = b.Next
				c++
			}
			for i := 0; i < c; i++ {
				headB = headB.Next
			}
			break
		}
		if b == nil {
			for a != nil {
				a = a.Next
				c++
			}
			for i := 0; i < c; i++ {
				headA = headA.Next
			}
			break
		}
	}
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}

func permute(nums []int) (ans [][]int) {
	path := []int{}
	used := make([]bool, len(nums))
	n := len(nums)
	var dfs func()
	dfs = func() {
		if len(path) == n {
			t := []int{}
			for k := range path {
				t = append(t, k)
			}
			ans = append(ans, t)
		}
		for i := 0; i < len(nums); i++ {
			if !used[nums[i]] {
				used[nums[i]] = true
				path = append(path, nums[i])
				dfs()
				path = path[:len(path)-1]
				used[nums[i]] = false
			}
		}
	}
	dfs()
	return
}

func findAnagrams(s string, p string) []int {
	ns, np := len(s), len(p)
	if ns < np {
		return nil
	}
	var res []int
	cntP := [26]int{}
	cntS := [26]int{}
	for i := 0; i < np; i++ {
		cntP[p[i]-'a']++
		cntS[s[i]-'a']++
	}
	if cntP == cntS {
		res = append(res, 0)
	}
	for i := np; i < ns; i++ {
		cntS[s[i]-'a']++
		cntS[s[i-np]-'a']--

		if cntP == cntS {
			res = append(res, i-np+1)
		}
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var root *ListNode
	for head != nil {
		head.Next, head, root = root, head.Next, head
	}
	return root
}

func mergeKLists(lists []*ListNode) (ans *ListNode) {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	var mg2 func(a *ListNode, b *ListNode) *ListNode
	mg2 = func(a, b *ListNode) *ListNode {
		dy := &ListNode{}
		cur := dy
		for a != nil && b != nil {
			if a.Val < b.Val {
				cur.Next = a
				a = a.Next
			} else {
				cur.Next = b
				b = b.Next
			}
			cur = cur.Next
		}
		if a == nil {
			cur.Next = b
		}
		if b == nil {
			cur.Next = a
		}
		return dy.Next
	}
	for i := 0; i < len(lists); i++ {
		ans = mg2(lists[i], ans)
	}
	return
}

func swapPairs(head *ListNode) (ans *ListNode) {
	if head == nil || head.Next == nil {
		return head
	}
	ans = head.Next
	head.Next = swapPairs(ans.Next)
	ans.Next = head
	return
}
func swapPairs1(head *ListNode) (ans *ListNode) {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	ans = head.Next
	for head.Next != nil {
		n := head.Next
		head.Next = n.Next
		n.Next = head
		if pre != nil {
			pre.Next = n
		}
		pre = head
		head = head.Next
	}
	return

}

// a -> b -> c
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	newhead := head
	for i := 0; i < k-1; i++ {
		newhead = newhead.Next
		if newhead == nil {
			return head
		}
	}
	root := newhead
	head.Next, root, head = reverseKGroup(newhead.Next, k), head, head.Next
	for root != newhead {
		head.Next, root, head = root, head, head.Next
	}
	return newhead
}

func combinationSum(candidates []int, target int) (ans [][]int) {
	path := []int{}
	var dfs func(int, int)
	dfs = func(start int, sum int) {
		if sum > target {
			return
		}
		if sum == target {
			ans = append(ans, append([]int(nil), path...))
		}
		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			sum += candidates[i]
			dfs(i, sum)
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	sum := 0
	dfs(0, sum)
	return
}

func combine(n int, k int) (ans [][]int) {
	path := []int{}
	arr := []int{}
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}
	var bc func(k int, start int)
	bc = func(k, start int) {
		if len(path) == k {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for i := start; i < n; i++ {
			path = append(path, arr[i])
			bc(k, i+1)
			path = path[:len(path)-1]
		}
	}
	bc(k, 0)
	return
}

type WordHeap []word

// Len implements [heap.Interface].
func (w *WordHeap) Len() int {
	return len(*w)
}

// Less implements [heap.Interface].
func (w *WordHeap) Less(i int, j int) bool {
	if (*w)[i].frq == (*w)[j].frq {
		return (*w)[i].st > (*w)[j].st
	}
	return (*w)[i].frq < (*w)[j].frq // 小根堆
}

// Pop implements [heap.Interface].
func (w *WordHeap) Pop() any {
	x := (*w)[len(*w)-1]
	*w = (*w)[:len(*w)-1]
	return x
}

// Push implements [heap.Interface].
func (w *WordHeap) Push(x any) {
	*w = append(*w, x.(word))
}

// Swap implements [heap.Interface].
func (w *WordHeap) Swap(i int, j int) {
	(*w)[i], (*w)[j] = (*w)[j], (*w)[i]
}

type word struct {
	st  string
	frq int
}

func topKFrequent(words []string, k int) []string {
	WordHeap := &WordHeap{}
	heap.Init(WordHeap)
	m := make(map[string]int)
	for _, v := range words {
		m[v]++
	}
	for s, f := range m {
		heap.Push(WordHeap, word{s, f})
		if len(*WordHeap) > k {
			heap.Pop(WordHeap)
		}

	}
	ans := []string{}
	for WordHeap.Len() > 0 {
		ans = append(ans, heap.Pop(WordHeap).(word).st)
	}
	slices.Reverse(ans)
	return ans
}

func lengthOfLongestSubstring(s string) (ans int) {
	if len(s) == 0 {
		return 0
	}
	l, r := 0, 1
	m := make(map[byte]bool)
	ans = 1
	m[byte(s[l])] = true
	for r < len(s) {
		// 右端点不存在
		if !m[byte(s[r])] {
			m[byte(s[r])] = true
			r++
			ans = max(ans, r-l)
		} else {
			for m[byte(s[r])] == true {
				m[byte(s[l])] = false
				l++
			}
		}
	}
	return
}
func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	// 确定一个数,再做两数之和
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			return
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				//去重
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}
	return
}

func maxArea(height []int) (ans int) {
	l, r := 0, len(height)-1
	for l < r {
		ans = max(ans, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return
}

func moveZeroes(nums []int) {
	s, f := 0, 0
	for f < len(nums) {
		if nums[f] != 0 {
			nums[f], nums[s] = nums[s], nums[f]
			s++
		}
		f++
	}
	fmt.Println(nums)
}

func longestConsecutive(nums []int) (ans int) {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	for k := range m {
		if _, ok := m[k-1]; !ok {
			tem := 1
			start := k + 1
			for {
				if m[start] == true {
					start++
					tem++
				} else {
					ans = max(ans, tem)
					break
				}
			}
		}
	}
	return
}
