package main

import (
	"container/heap"
	"fmt"
	"math"
	"slices"
	"sort"
)

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 2))
}

// 33. 搜索旋转排序数组s
func search(nums []int, target int) int {
	// 找到最小值,再在两边进行查询
	var findmin func([]int) int
	findmin = func(arr []int) int {
		l, r := 0, len(arr)-1
		for l <= r {
			mid := (l + r) / 2
			if mid == l && l == r {
				return mid
			} else if arr[mid] < arr[r] {
				r = mid
			} else {
				l = mid + 1
			}
		}
		return l
	}
	index := findmin(nums)
	l, r := 0, 0
	if target >= nums[index] && target <= nums[len(nums)-1] {
		l = index
		r = len(nums)
	} else {
		l = 0
		r = index
	}
	ans, ok := slices.BinarySearch(nums[l:r], target)
	if !ok {
		return -1
	}
	return l + ans
}

// 35. 搜索插入位置
func searchInsert(nums []int, target int) int {
	idx, _ := slices.BinarySearch(nums, target)

	return idx
}

type pair struct {
	dis, x int
}

type hp []pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].dis < h[j].dis
}

func (h *hp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *hp) Push(x any) {
	*h = append(*h, x.(pair))
}

func (h *hp) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minCost(n int, edges [][]int) int {
	h := &hp{{}} // 初始化并且查入 {0,0} 表示最开始只有节点0,节点0到0的距离是0
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	// 表示从x点到 g[x].to 点的距离是 g[x].wt
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
		g[y] = append(g[y], edge{x, wt * 2}) // 反转边
	}
	dis := make([]int, n)
	for i := range dis {
		dis[i] = math.MaxInt
	}
	dis[0] = 0
	for h.Len() > 0 {
		cur := heap.Pop(h).(pair)
		if cur.dis > dis[cur.x] {
			continue
		}
		if cur.x == n-1 { // 到达终点
			return cur.dis
		}
		for _, e := range g[cur.x] {
			if cur.dis+e.wt < dis[e.to] {
				dis[e.to] = cur.dis + e.wt
				heap.Push(h, pair{dis[e.to], e.to})
			}
		}
	}
	return -1
}

// 1200. 最小绝对差
func minimumAbsDifference(arr []int) (ans [][]int) {
	sort.Ints(arr)
	min := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] < min {
			min = arr[i] - arr[i-1]
		}
	}
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] == min {
			ans = append(ans, []int{arr[i-1], arr[i]})
		}
	}
	return
}

// 744. 寻找比目标字母大的最小字母
func nextGreatestLetter(letters []byte, target byte) byte {
	for _, v := range letters {
		if v > target {
			return v
		}
	}
	return letters[0]
}

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	fmt.Println(nums)
	min := nums[k-1] - nums[0]
	for i := k; i < len(nums); i++ {
		fmt.Println(nums[i], nums[i-k+1])
		if nums[i]-nums[i-k+1] < min {
			min = nums[i] - nums[i-k+1]
		}
	}
	return min
}

// 236. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if l == nil {
		return r
	}
	return l
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	index := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			index = i
			break
		}
	}
	inorderleft := inorder[:index]
	inorderright := inorder[index+1:]
	preorderleft := preorder[1 : index+1]
	preorderright := preorder[index+1:]
	root.Left = buildTree(preorderleft, inorderleft)
	root.Right = buildTree(preorderright, inorderright)
	return root
}

type Node struct {
	Val, Key  int
	Pre, Next *Node
}

type LRUCache struct {
	Cap       int
	Size      int
	Dummy     *Node
	KeytoNode map[int]*Node
}

func Constructor(capacity int) LRUCache {
	d := &Node{}
	d.Pre = d
	d.Next = d
	ktn := make(map[int]*Node)
	return LRUCache{
		Cap:       capacity,
		Size:      0,
		Dummy:     d,
		KeytoNode: ktn,
	}
}

func (l *LRUCache) Rm(x *Node) {
	x.Pre.Next = x.Next
	x.Next.Pre = x.Pre
}
func (l *LRUCache) PushFont(x *Node) {
	x.Pre = l.Dummy
	x.Next = l.Dummy.Next
	x.Pre.Next = x
	x.Next.Pre = x
}

func (l *LRUCache) Get(key int) (ans int) {
	node := l.KeytoNode[key]
	if node == nil {
		return -1
	}
	ans = node.Val
	l.Rm(node)
	l.PushFont(node)
	return
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.KeytoNode[key]; ok {
		node.Val = value
		l.Rm(node)
		l.PushFont(node)
		return
	}
	newnode := &Node{
		Val: value,
		Key: key,
	}
	l.KeytoNode[key] = newnode
	l.PushFont(newnode)
	l.Size++
	if l.Size > l.Cap {
		tail := l.Dummy.Pre
		l.Rm(tail)
		delete(l.KeytoNode, tail.Key)
	}

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
