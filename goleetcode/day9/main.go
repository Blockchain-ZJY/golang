package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	// fmt.Println(maxSlidingWindow([]int{1, -1}, 1))
	fmt.Println(countTriples(5))
	// fmt.Println(int(math.Sqrt(float64(25))))
}

// 239. 滑动窗口最大值
// 单调队列
func maxSlidingWindow(nums []int, k int) []int {
	temp := nums[0]
	ans := []int{}
	if len(nums) == k {
		for i := 0; i < len(nums); i++ {
			temp = max(nums[i], temp)
		}
		ans = append(ans, temp)
		return ans
	}
	q := Deque{}
	q.PushBack(nums[0])
	if k == 1 {
		ans = append(ans, nums[0])
	}
	for i := 1; i < len(nums); i++ {
		pv, exi := q.PeekFront()
		if exi && i > k-1 && pv == nums[i-k] {
			q.PopFront()
		}
		backv, ok := q.PeekBack()
		for ok && backv < nums[i] {
			q.PopBack()
			backv, ok = q.PeekBack()
		}
		q.PushBack(nums[i])
		if i >= k-1 {
			k, _ := q.PeekFront()
			ans = append(ans, k)
		}
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

type Deque struct {
	items []int
}

// 队首入队
func (d *Deque) PushFront(item int) {
	d.items = append([]int{item}, d.items...)
}

// 队尾入队
func (d *Deque) PushBack(item int) {
	d.items = append(d.items, item)
}

// 队首出队
func (d *Deque) PopFront() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque is empty")
	}
	item := d.items[0]
	d.items = d.items[1:]
	return item, nil
}

// 队尾出队
func (d *Deque) PopBack() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque is empty")
	}
	item := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return item, nil
}

// 查看队首元素
func (d *Deque) PeekFront() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}
	return d.items[0], true
}

// 查看队尾元素
func (d *Deque) PeekBack() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}
	return d.items[len(d.items)-1], true
}

// 判断是否为空
func (d *Deque) IsEmpty() bool {
	return len(d.items) == 0
}

// 获取长度
func (d *Deque) Len() int {
	return len(d.items)
}

// 1925. 统计平方和三元组的数目
func countTriples(n int) int {
	res := 0
	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ans := int(math.Sqrt(float64(i*i + j*j)))
			if ans <= n && ans*ans == i*i+j*j {
				// fmt.Println("i,j,ans", i, j, ans, ans*ans == i*i+j*j, ans <= n)
				res += 2
			}
		}
	}
	return res
}
