package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 2, 1, 2, 3, 3, 3, 3, 2, 3, 1, 3, 2}, 2))
}

type KV struct {
	key   int
	value int
}

type freh []KV

// Len implements [heap.Interface].
func (f freh) Len() int {
	return len(f)
}

// Less implements [heap.Interface].
func (f freh) Less(i int, j int) bool {
	return f[i].value < f[j].value
}

// Pop implements [heap.Interface].
func (f *freh) Pop() any {
	x := (*f)[len(*f)-1]
	(*f) = (*f)[:len(*f)-1]
	return x
}

// Push implements [heap.Interface].
func (f *freh) Push(x any) {
	(*f) = append((*f), x.(KV))
}

// Swap implements [heap.Interface].
func (f *freh) Swap(i int, j int) {
	(*f)[i], (*f)[j] = (*f)[j], (*f)[i]
}

func topKFrequentHeap(nums []int, k int) (ans []int) {
	h := &freh{}
	heap.Init(h)
	feq := make(map[int]int)
	for _, v := range nums {
		feq[v]++
	}
	for key, v := range feq {
		heap.Push(h, KV{key, v})
		if len((*h)) > k {
			heap.Pop(h)
		}
	}
	for h.Len() > 0 {
		ans = append(ans, heap.Pop(h).(KV).key)
	}

	return
}

// 347. 前 K 个高频元素
func topKFrequent(nums []int, k int) (ans []int) {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	kv := make([]KV, len(freq))
	for ks, counts := range freq {
		kv = append(kv, KV{ks, counts})
	}
	sort.Slice(kv, func(i, j int) bool {
		return kv[i].value > kv[j].value
	})
	for i := 0; i < k; i++ {
		ans = append(ans, kv[i].key)
	}
	return
}

type Heap []int

// Len implements [heap.Interface].
func (h *Heap) Len() int {
	return len((*h))
}

// Less implements [heap.Interface].
func (h *Heap) Less(i int, j int) bool {
	return (*h)[i] < (*h)[j]
}

// Pop implements [heap.Interface].
func (h *Heap) Pop() any {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

// Push implements [heap.Interface].
func (h *Heap) Push(x any) {
	*h = append((*h), x.(int))
}

// Swap implements [heap.Interface].
func (h *Heap) Swap(i int, j int) {
	(*h)[j], (*h)[i] = (*h)[i], (*h)[j]
}

// 215. 数组中的第K个最大元素
func findKthLargest(nums []int, k int) int {
	h := &Heap{}
	heap.Init(h)
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
		if len(*h) > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}

/*
189. 轮转数组
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
*/
func rotate(nums []int, k int) {
	k = k % len(nums)
	var rv func(nums []int)
	rv = func(nums []int) {
		l, r := 0, len(nums)-1
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	rv(nums)
	rv(nums[:k])
	rv(nums[k:])
}

// 56. 合并区间
/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi]。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
*/
func merge(intervals [][]int) (ans [][]int) {
	//按照起始的点进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans = append(ans, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := ans[len(ans)-1]
		cur := intervals[i]
		if cur[0] > last[1] {
			ans = append(ans, cur)
		} else {
			if cur[1] > last[1] {
				last[1] = cur[1]
			}
		}
	}
	return
}
