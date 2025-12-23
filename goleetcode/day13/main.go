package main

import (
	"fmt"
	"sort"
)

func main() {
	// // fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	// // board2 := [][]byte{
	// // 	{'.', '.', '4', '.', '.', '.', '6', '3', '.'},
	// // 	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	// // 	{'5', '.', '.', '.', '.', '.', '.', '9', '.'},
	// // 	{'.', '.', '.', '5', '6', '.', '.', '.', '.'},
	// // 	{'4', '.', '3', '.', '.', '.', '.', '.', '1'},
	// // 	{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
	// // 	{'.', '.', '.', '1', '.', '.', '.', '.', '.'},
	// // 	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	// // 	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	// // matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	// matrix1 := [][]int{{7}, {9}, {6}}
	// // fmt.Println(maxProfit([]int{5, 4, 3}, []int{1, 1, 0}, 2))
	// // fmt.Println(findAllPeople(6, [][]int{{1, 2, 5}, {2, 3, 8}, {1, 5, 10}}, 1))
	// fmt.Println(spiralOrder(matrix1))
	// // fmt.Println(maxProfit([]int{5, 4, 3}, []int{1, 1, 0}, 2))
	// // fmt.Println(isValidSudoku(board2))
	// a := &TreeNode{
	// 	Left:  nil,
	// 	Right: nil,
	// 	Val:   0,
	// }
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	// // matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// fmt.Println(isValidBST(a))
	fmt.Println(searchMatrix(matrix, -1))
}

// 35. 搜索插入位置
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)

	for l < r {
		mid := (l + r) / 2
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

/*
每行必须包含 1–9 且不重复。
每列必须包含 1–9 且不重复。
每个 3×3 宫格必须包含 1–9 且不重复。
如果初始题目已经违反这些规则（例如某行有两个相同数字），则必然无解。
*/
func isValidSudoku(board [][]byte) bool {
	// count := 0
	for i := range board {
		c := make(map[byte]bool)
		for j := range board[i] {
			if board[i][j] == '.' {
				continue
			} else if c[board[i][j]] {
				return false
			} else {
				// count++
				c[board[i][j]] = true
			}
		}
	}
	for i := range board {
		l := make(map[byte]bool)
		for j := range board[i] {
			if board[j][i] == '.' {
				continue
			} else if l[board[j][i]] {
				return false
			} else {
				l[board[j][i]] = true
			}
		}
	}

	// if count < 17 {
	// 	return false
	// }
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			m := make(map[byte]bool)
			for x := i * 3; x < (i+1)*3; x++ {
				for y := j * 3; y < (j+1)*3; y++ {
					if board[x][y] == '.' {
						continue
					}
					if m[board[x][y]] {
						fmt.Println(i, j)
						return false
					} else {
						m[board[x][y]] = true
					}
				}
			}
		}
	}
	return true
}

// 3652. 按策略买卖股票的最佳时机
func maxProfit(prices []int, strategy []int, k int) int64 {
	ans := 0
	for i := 0; i < len(strategy); i++ {
		ans += strategy[i] * prices[i]
	}
	delta := 0
	for j := 0; j < k/2; j++ {
		delta += -strategy[j] * prices[j]
	}
	for j := k / 2; j < k; j++ {
		delta += (1 - strategy[j]) * prices[j]
	}
	best := ans + delta
	if delta < 0 {
		best = ans
	}
	for i := k; i < len(prices); i++ {
		//新增的元素加入
		delta += (1 - strategy[i]) * prices[i]
		delta -= prices[i-k/2]
		//删除的元素
		delta += strategy[i-k] * prices[i-k]
		best = max(best, ans+delta)
	}
	return int64(best)
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	// 1. 按照时间排序
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})
	// 2. 同一时间内,所有的连接的人构成一个图
	// 例如 [1,4,3] [1,5,3] [2,4,3] 1 2 4 5 四个节点构成一个图
	// 如果其中有一个人知道秘密,图遍历所有节点标记知道秘密的人

	knows := map[int]bool{0: true, firstPerson: true}
	for i := range meetings {
		// 不越界并且是同一时间
		graph := make(map[int][]int)
		x, y := meetings[i][0], meetings[i][1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
		for ; i+1 < len(meetings) && meetings[i][2] == meetings[i+1][2]; i++ {
			x, y := meetings[i][0], meetings[i][1]
			graph[x] = append(graph[x], y)
			graph[y] = append(graph[y], x)
		}
		var dfs func(start int)
		dfs = func(start int) {
			if knows[start] {
				return
			}
			knows[start] = true
			for _, neighbor := range graph[start] {
				dfs(neighbor)
			}
		}
		fmt.Println(knows)
		// 对与当前图,遍历所有节点
		for node := range graph {
			fmt.Println("node", node)
			if knows[node] {
				dfs(node)
			}
		}
		// fmt.Println(i, graph)
	}

	ans := []int{}
	for node := range knows {
		ans = append(ans, node)
	}
	return ans
}

// 54. 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	ans := []int{}
	var getans func(n, m, startx, starty int)
	getans = func(n, m, startx, starty int) {
		if n < 1 || m < 1 {
			return
		}
		for i := 0; i < n; i++ {
			ans = append(ans, matrix[startx][starty+i])
			fmt.Println("往→走")
			fmt.Println(matrix[startx][starty+i])
		}
		if m <= 1 {
			return
		}
		for i := 1; i < m; i++ {
			ans = append(ans, matrix[startx+i][starty+n-1])
			fmt.Println("往↓走")
			fmt.Println(matrix[startx+i][starty+n-1])
		}
		if n-2 < 0 {
			return
		}
		for i := n - 2; i >= 0; i-- {
			ans = append(ans, matrix[startx+m-1][starty+i])
			fmt.Println("往←走")
			fmt.Println(matrix[startx+m-1][starty+i])
		}

		for i := m - 2; i >= 1; i-- {
			ans = append(ans, matrix[startx+i][starty])
			fmt.Println("往↑走")
			fmt.Println(matrix[startx+i][starty])
		}
		getans(n-2, m-2, startx+1, starty+1)
	}
	getans(len(matrix[0]), len(matrix), 0, 0)
	return ans
}

func minDeletionSize(strs []string) (ans int) {
	n, m := len(strs), len(strs[0])
	a := make([]string, n) // 最终得到的字符串数组
	//
next:
	for j := 0; j < m; j++ {
		for i := 0; i < n-1; i++ {
			if a[i]+string(strs[i][j]) > a[i+1]+string(strs[i+1][j]) {
				// j 列不是升序，必须删
				ans++
				continue next
			}
			// j 列是升序，不删更好
			for i, s := range strs {
				a[i] += string(s[j])
			}
		}
	}
	return
}

// 旋转图像
// 矩阵顺时针旋转90° = 对角线翻转 + 每行对调
func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	//每行对调
	for i := 0; i < len(matrix); i++ {
		l, r := 0, len(matrix)-1
		for l < r {
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
			l++
			r--
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre *TreeNode

// 中序遍历
func isValidBST(root *TreeNode) bool {

	if root == nil {
		return true
	}
	left := isValidBST(root.Left)
	if pre != nil && root.Val <= pre.Val {
		return false
	}
	pre = root
	right := isValidBST(root.Right)
	return left && right
}

// 240. 搜索二维矩阵 II
func searchMatrix(matrix [][]int, target int) bool {
	// 从右上角开始
	i, j := 0, len(matrix[0])-1
	for {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
		} else {
			i++
		}
		if i > len(matrix)-1 || j < 0 {
			return false
		}
		fmt.Println(matrix[i][j])
	}
	return false
}

// 59. 螺旋矩阵 II
func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	v := 1
	var getans func(n, startx, starty int)
	getans = func(n, startx, starty int) {
		if n < 1 {
			return
		}
		for i := 0; i < n; i++ {
			ans[startx][starty+i] = v
			v++
		}
		if n <= 1 {
			return
		}
		for i := 1; i < n; i++ {
			ans[startx+i][starty+n-1] = v
			v++
		}
		if n-2 < 0 {
			return
		}
		for i := n - 2; i >= 0; i-- {
			ans[startx+n-1][starty+i] = v
			v++
		}
		for i := n - 2; i >= 1; i-- {
			ans[startx+i][starty] = v
			v++
		}
		getans(n-2, startx+1, starty+1)
	}
	getans(n, 0, 0)
	return ans
}
