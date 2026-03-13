package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/emirpasic/gods/stacks/arraystack"
)

// 1
func twoSum(nums []int, target int) []int {
	maps := make(map[int]int)
	maps[nums[0]] = 0
	ans := []int{}
	for i := 1; i < len(nums); i++ {
		if _, ok := maps[target-nums[i]]; ok {
			ans = append(ans, []int{i, maps[target-nums[i]]}...)
			return ans
		}
		maps[nums[i]] = i
	}
	return ans
}

// 2 字母异位词分组
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, v := range strs {
		count := [26]int{}
		for _, c := range v {
			count[c-'a']++
		}
		m[count] = append(m[count], v)
	}
	ans := [][]string{}
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

// 3 最长连续序列  100,4,200,1,3,2
func longestConsecutive(nums []int) int {
	a := make(map[int]bool)
	for _, v := range nums {
		a[v] = true
	}
	ans := math.MinInt16
	for k, _ := range a {
		if _, ok := a[k-1]; !ok { // 是起始位置
			temp := 1
			cur := k + 1
			for !a[cur] {
				temp++
				cur++
			}
			ans = max(ans, temp)
		}
	}
	return ans
}

// 4.移动零
// 输入: nums = [0,1,0,3,12]
// 输入: nums = [1,0,0,3,12]
// 以右端点移动,左端点
// 输出: [1,3,12,0,0]
func moveZeroes(nums []int) {
	l, r, n := 0, 0, len(nums)
	for r < n {
		if nums[r] != 0 {
			nums[r], nums[l] = nums[l], nums[r]
			l++
		}
		r++
	}
}

// 5. 盛最多水的容器 注意是当前较小的去缩小
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		ans = max(ans, min(height[l], height[r])*(r-l))
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return ans
}

// 6 三数之和
//
//	注意同一个解要跳过
func threeSum(nums []int) [][]int {
	// 双指针
	sort.Ints(nums)
	ans := [][]int{}
	n := len(nums)
	//确定i 双指针
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			return ans
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for nums[l] == nums[l-1] && l+1 < n {
					l++
				}
				for nums[r+1] == nums[r] && r-1 > 0 {
					r--
				}
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}

	}
	return ans
}

// 7 接雨水

// 接雨水 单调栈-栈存的是下标
func trap(height []int) int {
	n := len(height)
	if n == 1 {
		return 0
	}
	stk := arraystack.New()
	stk.Push(0)
	sum := 0
	for i := 1; i < n; i++ {
		for !stk.Empty() {
			p, _ := stk.Peek()
			if height[p.(int)] >= height[i] {
				break
			}
			mid := p.(int)
			stk.Pop()
			if _, ok := stk.Peek(); !ok {
				break
			}
			left, _ := stk.Peek()
			h := min(height[left.(int)], height[i]) - height[mid]
			w := i - left.(int) - 1
			sum += h * w
		}
		stk.Push(i)
	}
	return sum
}

// 8 找到字符串中所有字母异位词
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	res := []int{}
	m := make(map[[26]int]bool)
	target := [26]int{}
	ans := [26]int{}
	for i := range p {
		target[p[i]-'a']++
		ans[s[i]-'a']++
	}
	m[target] = true
	if m[ans] {
		res = append(res, 0)
	}
	// fmt.Println(ans, target, m)
	for i := 1; i < len(s)-len(p)+1; i++ {
		fmt.Println(i)
		ans[s[i-1]-'a']--
		ans[s[i+len(p)-1]-'a']++
		if m[ans] {
			res = append(res, i)
		}
	}
	return res
}

// 239. 滑动窗口最大值
// 单调队列
func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	ans := []int{}
	push := func(i int) {
		for len(q) != 0 && nums[len(q)-1] < nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	for i, _ := range nums {
		push(i)
		if q[0] < i-k+1 {
			q = q[1:]
		}
		if i >= k-1 {
			ans = append(ans, nums[q[0]])
		}
	}
	return ans
}

// 234. 回文链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针,后半部分进行翻转
func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	f, s := head.Next, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	var rvslist func(h *ListNode) *ListNode
	rvslist = func(h *ListNode) *ListNode {
		var root *ListNode
		for h != nil {
			h.Next, root, h = root, h, h.Next
		}
		return root
	}
	root := rvslist(s.Next)
	for root != nil {
		if root.Val != head.Val {
			return false
		}
		root = root.Next
		head = head.Next
	}
	return true
}

// 21. 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	head := &ListNode{}
	p := head
	for list1 != nil || list2 != nil {
		if list1 == nil {
			p.Next = list2
			return head.Next
		}
		if list2 == nil {
			p.Next = list1
			return head.Next
		}
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2

			list2 = list2.Next
		}
		p = p.Next
	}
	return head.Next
}
