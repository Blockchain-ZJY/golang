package main

import (
	"container/heap"
	"fmt"
	"math"
	"slices"
)

type IntHeap []int

func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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
func main() {
	a := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(a)
	// fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
	// fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

// 小根堆
func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < len(nums); i++ {
		if i < k {
			heap.Push(h, nums[i])
		} else if nums[i] > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}
	return (*h)[0]
}

// 股票平滑下跌阶段的数目
func getDescentPeriods(prices []int) int64 {
	ans := 0
	sum := 0
	for i := 0; i < len(prices); i++ {
		sum = 1
		for i+1 < len(prices) && prices[i+1]+1 == prices[i] {
			sum++
			i++
		}
		fmt.Println(i, sum)
		ans += sum * (sum + 1) / 2
	}
	return int64(ans)
}

// 缺失的第一个正数
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		if nums[i] > n || nums[i] <= 0 {
			nums[i] = -1
		}
	}

	for i := range nums {
		for nums[i] > 0 && nums[i] != i+1 {
			if nums[i] == nums[nums[i]-1] {
				nums[i] = -1
				break
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	fmt.Println(nums)
	for i := 0; i < n; i++ {
		if nums[i] == -1 {
			return i + 1
		}

	}
	return n + 1
}

// 136. 只出现一次的数字
func singleNumber(nums []int) int {
	sig := 0
	for i := range nums {
		sig ^= nums[i]
	}
	return sig
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0)
	push := func(x int) { //如果新元素 nums[x] 比队尾对应的值大，就把队尾弹出（保证队列递减
		for len(queue) > 0 && nums[x] > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, x)
	}
	for i := 0; i < k; i++ {
		push(i)
	}

	ans := make([]int, 0)
	ans = append(ans, nums[queue[0]])
	for i := k; i < len(nums); i++ {
		if queue[0] < i-k+1 {
			queue = queue[1:]
		}
		push(i)
		ans = append(ans, nums[queue[0]])
	}
	return ans
}

func setZeroes(matrix [][]int) {
	for x, v := range matrix {
		for k, s := range v {
			if s == 0 {
				for i := 0; i < len(v); i++ {
					if matrix[x][i] == 0 {
						continue
					}
					matrix[x][i] = math.MaxInt64
				}
				for i := 0; i < len(matrix); i++ {
					if matrix[i][k] == 0 {
						continue
					}
					fmt.Println(i, x)
					matrix[i][k] = math.MaxInt64
				}
			}
		}
	}
	for x := range matrix {
		for y := range matrix[x] {
			if matrix[x][y] == math.MaxInt64 {
				matrix[x][y] = 0
			}
		}
	}
}

// 用第一行第一列来记录是是否有需要值为零
func setZeroebetter(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	firstRow := slices.Contains(matrix[0], 0)
	firstCol := false
	for _, row := range matrix {
		if row[0] == 0 {
			firstCol = true
			break
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if firstCol {
		for _, row := range matrix {
			row[0] = 0
		}
	}
	if firstRow {
		clear(matrix[0])
	}
}
