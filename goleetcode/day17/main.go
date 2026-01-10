package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(totalNQueens(4))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const MOD = 1000000007

func maxProduct(root *TreeNode) (ans int) {
	if root == nil {
		return 0
	}
	m := make(map[*TreeNode]int)
	var getsum func(root *TreeNode) int
	getsum = func(root *TreeNode) (sum int) {
		if root == nil {
			return 0
		}
		t := root.Val + getsum(root.Left) + getsum(root.Right)
		m[root] = t
		return t
	}
	// 求出完整树的和,并且尽量均分
	total := getsum(root)
	//最小差值
	minsub := math.MaxInt
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		//以当前节点进行切分计算该节点为root的总和
		thisside := m[root]
		otherside := total - thisside
		if abs(otherside-thisside) < minsub {
			minsub = abs(otherside - thisside)
			ans = (otherside * thisside) % MOD
		}
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)
	return
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 129. 求根节点到叶节点数字之和
func sumNumbers(root *TreeNode) (ans int) {
	anslist := []int{}
	path := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil {
			sum := 0
			for _, v := range path {
				sum = sum*10 + v
			}
			anslist = append(anslist, sum)
			path = path[:len(path)-1]
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		path = path[:len(path)-1]
	}

	dfs(root) // 作为参数传递是不需要回溯的,值类型,作为全局变量是需要回溯的
	for i := 0; i < len(anslist); i++ {
		ans += anslist[i]
	}
	return
}

func sumNumbersBetter(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, cur int) int {
		if node == nil {
			return 0
		}
		cur = cur*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return cur
		}
		return dfs(node.Left, cur) + dfs(node.Right, cur)
	}
	return dfs(root, 0)
}

// 124. 二叉树中的最大路径和
func maxPathSum(root *TreeNode) (ans int) {
	// 以当前节点为根节点计算最大路径和
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		// 可能出现负数,如果是负数就不选
		maxleft := max(0, dfs(root.Left))
		maxright := max(0, dfs(root.Right))
		ans = max(maxleft+maxright+root.Val, ans)
		//给他的父节点使用,用max是保证只选择最大的一条边
		return max(maxleft, maxright) + root.Val
	}
	ans = root.Val
	dfs(root)
	return
}

// 84. 柱状图中最大的矩形  单调栈确定每个i的左右边界从而计算最大面积
// 左右边界指的是以当前元素为高,计算能最大延伸多长(也就是在左右找到第一个比height[i]小的元素)
// 所以在单调递增栈里面,每次来了一个新的元素如果比栈顶元素大,对与当前元素左边界就是栈顶元素,需要压栈操作
// 如果比栈顶元素小,需要出栈,对于出栈的元素第一个比它小的就是当前元素
// 入栈的时候  栈顶元素就是当前元素的左边界
// 出栈的时候  出栈元素的右边界就是当前元素
func largestRectangleArea(heights []int) int {
	left := make([]int, len(heights)+1)
	right := make([]int, len(heights)+1)
	for i := range left {
		left[i] = -1
		right[i] = len(heights)
	}
	sk := Stack{
		items: []int{},
	}
	for i := range heights {
		// 如果栈顶元素小于 当前元素 入栈, 确定当前元素的左边界
		if heights[sk.Peek()] < heights[i] {
			left[i] = sk.Peek()
			sk.Push(i)
			continue
		}
		for !sk.IsEmpty() && heights[sk.Peek()] > heights[i] {
			right[sk.Peek()] = i
			sk.Pop()
		}
		if !sk.IsEmpty() {
			left[i] = sk.Peek()
		}
		sk.Push(i)
	}
	ans := 0
	for i := 0; i < len(heights); i++ {
		w := right[i] - left[i] - 1
		h := heights[i]
		ans = max(ans, w*h)
	}
	return ans
}

type Stack struct {
	items []int
}

// 入栈
func (s *Stack) Push(ch int) {
	s.items = append(s.items, ch)
}

// 出栈
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

// 查看栈顶
func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		return 0
	}
	return s.items[len(s.items)-1]
}

// 判断是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// 栈大小
func (s *Stack) Size() int {
	return len(s.items)
}
func numIslands(grid [][]byte) (ans int) {
	var dfs func(grid [][]byte, i, j int, visited [][]bool)
	dfs = func(grid [][]byte, i, j int, visited [][]bool) {
		// 边界条件
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
			return
		}
		if visited[i][j] || grid[i][j] == '0' {
			return
		}

		visited[i][j] = true
		// 四个方向
		dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, d := range dirs {
			dfs(grid, i+d[0], j+d[1], visited)
		}
	}
	n := len(grid)
	m := len(grid[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				ans++
				dfs(grid, i, j, visited)
			}
		}
	}
	return
}

func numIslandsBFS(grid [][]byte) (count int) {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				count++
				// BFS 队列
				queue := [][]int{{i, j}}
				visited[i][j] = true
				for len(queue) > 0 {
					cur := queue[0]
					queue = queue[1:]
					for _, d := range dirs {
						ni, nj := cur[0]+d[0], cur[1]+d[1]
						if ni >= 0 && nj >= 0 && ni < m && nj < n {
							if grid[ni][nj] == '1' && !visited[ni][nj] {
								visited[ni][nj] = true
								queue = append(queue, []int{ni, nj})
							}
						}
					}
				}
			}
		}
	}
	return count
}

func orangesRotting(grid [][]int) (ans int) {
	dis := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	m, n := len(grid), len(grid[0])
	minutes := -1
	fresh := 0
	q := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				q = append(q, []int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}
	if fresh == 0 {
		return 0
	}
	for len(q) > 0 {
		lenth := len(q)
		for lenth != 0 {
			curi, curj := q[0][0], q[0][1]
			q = q[1:]
			for _, d := range dis {
				ni, nj := curi+d[0], curj+d[1]
				if ni >= 0 && nj >= 0 && ni < m && nj < n && grid[ni][nj] == 1 {
					fresh--
					q = append(q, []int{ni, nj})
					grid[ni][nj] = 2
				}
			}
			lenth--
		}
		ans++
	}
	if fresh != 0 {
		return -1
	}
	return minutes
}

// 865. 具有所有最深节点的最小子树
func subtreeWithAllDeepestBFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	//层序遍历获取所有
	q := []*TreeNode{}
	parent := make(map[*TreeNode]*TreeNode)
	lastlayer := []*TreeNode{}
	q = append(q, root)
	for len(q) != 0 {
		l := len(q)
		lastlayer = []*TreeNode{}
		for i := 0; i < l; i++ {
			cur := q[0]
			q = q[1:]
			// 确定parent
			lastlayer = append(lastlayer, cur)
			if cur.Left != nil {
				q = append(q, cur.Left)
				parent[cur.Left] = cur
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
				parent[cur.Right] = cur
			}
		}
	}
	// 如果最深层只有一个节点，直接返回
	if len(lastlayer) == 1 {
		return lastlayer[0]
	}
	// 当前层的节点
	cur := lastlayer
	// 不断向上跳，直到只剩一个节点
	for len(cur) > 1 {
		// 用 map 去重
		next := make(map[*TreeNode]bool)
		for _, node := range cur {
			next[parent[node]] = true
		}
		// 把 map 转回 slice
		cur = cur[:0]
		for p := range next {
			cur = append(cur, p)
		}
	}
	// 返回唯一节点
	return cur[0]
}

// dfs(node) → (depth, ans)
// 对每个节点 node，DFS 返回：
// 这个节点为根的子树的最大深度 depth
// 这个子树中包含所有最深节点的最小子树的根 ans
func subtreeWithAllDeepestDFS(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode) (int, *TreeNode)
	dfs = func(node *TreeNode) (int, *TreeNode) {
		if node == nil {
			return -1, nil
		}
		ld, ln := dfs(node.Left)
		rd, rn := dfs(node.Right)
		if ld == rd {
			return ld + 1, node
		}
		if ld > rd {
			return ld + 1, ln
		}
		return rd + 1, rn
	}
	_, ans := dfs(root)
	return ans
}

// 236. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) (ans *TreeNode) {
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) (res int) {
		mid := 0
		if root == nil {
			return 0
		}
		l := dfs(root.Left)
		r := dfs(root.Right)
		if root == p || root == q {
			mid++
		}
		total := l + r + mid
		if total == 2 && ans == nil {
			ans = root

		}
		return total
	}
	dfs(root)
	return
}

// 52. N 皇后 II 返回方案数量
func totalNQueens(n int) (ans int) {
	// 初始化棋盘
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	// 判断当前x,y是否满足
	var isok func(board [][]byte, x, y int) bool
	isok = func(board [][]byte, x, y int) bool {
		for i := 0; i < n; i++ {
			if board[x][i] == 'Q' || board[i][y] == 'Q' {
				return false
			}
			// 斜边是否满足
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					if board[i][j] == 'Q' {
						if i-j == x-y || i+j == x+y {
							return false
						}
					}
				}
			}
		}
		return true
	}

	var dfs func(board [][]byte, y int)
	dfs = func(board [][]byte, y int) {
		if y == n {
			ans++
			return
		}
		for i := 0; i < n; i++ {
			if isok(board, i, y) {
				board[i][y] = 'Q'
				dfs(board, y+1)
				board[i][y] = '.'
			}
		}
	}

	dfs(board, 0)
	return
}
