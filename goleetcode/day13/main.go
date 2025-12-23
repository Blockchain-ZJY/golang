package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	// board2 := [][]byte{
	// 	{'.', '.', '4', '.', '.', '.', '6', '3', '.'},
	// 	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	// 	{'5', '.', '.', '.', '.', '.', '.', '9', '.'},
	// 	{'.', '.', '.', '5', '6', '.', '.', '.', '.'},
	// 	{'4', '.', '3', '.', '.', '.', '.', '.', '1'},
	// 	{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
	// 	{'.', '.', '.', '1', '.', '.', '.', '.', '.'},
	// 	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	// 	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	// }

	// fmt.Println(maxProfit([]int{5, 4, 3}, []int{1, 1, 0}, 2))
	fmt.Println(findAllPeople(6, [][]int{{1, 2, 5}, {2, 3, 8}, {1, 5, 10}}, 1))
	// fmt.Println(maxProfit([]int{5, 4, 3}, []int{1, 1, 0}, 2))
	// fmt.Println(isValidSudoku(board2))
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
	for i := 0; i < len(meetings); i++ {
		// 不越界并且是同一时间
		graph := make(map[int][]int)
		vis := make(map[int]bool)
		x, y := meetings[i][0], meetings[i][1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)

		for ; i < len(meetings)-1 && meetings[i][2] == meetings[i+1][2]; i++ {
			x, y := meetings[i+1][0], meetings[i+1][1]
			graph[x] = append(graph[x], y)
			graph[y] = append(graph[y], x)
		}
		fmt.Println(graph)
		var dfs func(start int)
		dfs = func(start int) {
			vis[start] = true
			knows[start] = true
			for _, neighbor := range graph[start] {
				if !vis[neighbor] {
					dfs(neighbor)
				}
			}
		}

		for node := range graph {
			if knows[node] {
				dfs(node)
			}
		}
	}

	ans := []int{}
	for node := range knows {
		ans = append(ans, node)
	}
	return ans
}
