package main

import (
	"container/heap"
	"fmt"
	"math"
	"slices"
	"sort"
)

func main() {

	fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
}

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	mid := (l + r) / 2
	for l < r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// 74. 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1

	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			if matrix[i][sort.SearchInts(matrix[i], target)] != target {
				return false
			} else {
				return true
			}
		} else {
			i++
		}
	}
	return false
}

// 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	x := sort.SearchInts(nums, target)
	y := sort.SearchInts(nums, target+1)
	if x == len(nums) || nums[x] != target {
		//target not exi
		return []int{-1, -1}
	}
	return []int{x, y - 1}
}

// 33. 搜索旋转排序数组
func search(nums []int, target int) int {
	index := findMin(nums)
	n := len(nums)

	// target 在右半段
	if target >= nums[index] && target <= nums[n-1] {
		i := sort.SearchInts(nums[index:], target)
		if i < len(nums[index:]) && nums[index+i] == target {
			return index + i
		}
		return -1
	}

	// target 在左半段
	i := sort.SearchInts(nums[:index], target)
	if i < index && nums[i] == target {
		return i
	}
	return -1
}

// 153. 寻找旋转排序数组中的最小值
func findMin(nums []int) int {
	l, r := 0, len(nums)-1 // [l,r]
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if nums[l] <= nums[mid] && nums[mid] <= nums[r] {
			return l
		} else if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

type IntHeap []int

func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // i 的 优先级跟高
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	len := len(old)
	x := old[len-1]
	*h = (*h)[0 : len-1]
	return x
}

//3075. 幸福值最大化的选择方案

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Ints(happiness)
	ans := 0
	index := 0
	for i := len(happiness) - 1; i >= 0; i-- {
		if happiness[i]-index >= 0 {
			ans += happiness[i] - index
		} else {
			return int64(ans)
		}
		index++
		k--
		if k == 0 {
			return int64(ans)
		}
	}
	return int64(ans)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 23. 合并 K 个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}
	a, b := lists[0], lists[1]
	lists = append(lists[2:], mergetwo(a, b))
	return mergeKLists(lists)
}

func mergetwo(a, b *ListNode) *ListNode {
	dum := &ListNode{}
	t := dum
	for a != nil && b != nil {
		if a.Val < b.Val {
			t.Next = a
			a = a.Next
			t = t.Next
		} else {
			t.Next = b
			b = b.Next
			t = t.Next
		}
	}
	if a != nil {
		t.Next = a
	} else {
		t.Next = b
	}

	return dum.Next
}

// 2483. 商店的最少代价
// YYNYNY
func bestClosingTime(customers string) int {
	presumN := []int{}
	presumY := []int{}
	presumN = append(presumN, 0)
	presumY = append(presumY, 0)
	sum := 0
	for i := 0; i < len(customers); i++ {
		if customers[i] == 'N' {
			sum++
		}
		presumN = append(presumN, sum)
	}
	sum = 0
	for i := len(customers) - 1; i >= 0; i-- {
		if customers[i] == 'Y' {
			sum++
		}
		presumY = append(presumY, sum)
	}
	slices.Reverse(presumY)
	ans := []int{}
	for i := 0; i < len(customers)+1; i++ {
		ans = append(ans, presumN[i]+presumY[i])
	}
	min := math.MaxInt32
	res := 0
	for i := len(customers); i >= 0; i-- {
		if ans[i] <= min {
			min = ans[i]
			res = i
		}
	}
	fmt.Println(ans)
	return res
}

// 1353. 最多可以参加的会议数目
func maxEvents(events [][]int) int {
	mx := 0
	for i := range events {
		mx = max(events[i][1], mx)
	}
	slices.SortFunc(events, func(a, b []int) int {
		return a[0] - b[0]
	})
	ans := 0
	// 对于day,堆顶就是能消费的最优值
	// 堆里放的是所有已经开始但还没结束的会议的 endDay。
	hp := &IntHeap{}
	heap.Init(hp)
	j := 0
	for day := 1; day <= mx; day++ {
		// 先删除当天已经参加不了会议
		for hp.Len() > 0 && (*hp)[0] < day {
			heap.Pop(hp)
		}
		for j < len(events) && events[j][0] <= day { // 小于当前的值
			heap.Push(hp, events[j][1])
			j++
		}
		if hp.Len() > 0 {
			heap.Pop(hp)
			ans++
		}
	}
	return ans
}

// 455. 分发饼干
func findContentChildren(g []int, s []int) (ans int) {
	sort.Ints(s)
	sort.Ints(g)
	// g []int{1, 2, 3},  s []int{3}
	for i, j := len(g)-1, len(s)-1; i >= 0 && j >= 0; {
		if s[j] >= g[i] {
			ans++
			i--
			j--
		} else {
			i--
		}
	}
	return
}

// 跳跃游戏
func canJump(nums []int) (ans bool) {
	if len(nums) == 0 {
		return true
	}
	mx := nums[0]
	for i := 1; i < len(nums); i++ {
		if mx >= len(nums)-1 {
			return true
		}
		// 2 0 0 3
		if i > mx {
			return false
		}
		mx = max(i+nums[i], mx)
	}
	return
}

// 435. 无重叠区间
// 对于两个  	A: [s1, e1] B: [s2, e2]
// 选完某个区间后，未来能选的区间尽可能多-> 结束时间越早越好
func eraseOverlapIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[1] - b[1]
	})
	fmt.Println(intervals)
	ans := 0
	endtime := -5000000
	for i := range intervals {
		if intervals[i][0] < endtime {
			ans++
		} else {
			endtime = intervals[i][1]
		}
	}
	return ans
}

// 452. 用最少数量的箭引爆气球
func findMinArrowShots(points [][]int) (ans int) {

	// 右端点进行排序,每个气球都要被安排一次,是设定为左端点还是右端点?
	// 右端点  原因排序过后的 要不不相交, 相交的话一定右端点是交点(排序后)
	// 换句话说排序过后的右端点们, 相邻的两个区间,后面的一个左端点要不超过
	// 前面的右边, 要不不超过,所以能覆盖
	slices.SortFunc(points, func(a, b []int) int {
		return a[1] - b[1]
	})
	for i := 0; i < len(points); i++ {
		x := points[i][1] // 获取右端点
		for i+1 < len(points) && points[i+1][0] <= x {
			i++
		}
		ans++
	}
	return
}
