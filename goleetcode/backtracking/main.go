package main

import (
	"fmt"
	"sort"
)

func readBinaryWatch(turnedOn int) []string {
	leds := []int{8, 4, 2, 1, 32, 16, 8, 4, 2, 1}
	res := []string{}
	path := []int{} // 存索引，而不是值
	backtrackIndex(leds, turnedOn, 0, path, &res)
	return res
}

func backtrackIndex(leds []int, k, start int, path []int, res *[]string) {
	if len(path) == k {
		hour, minute := 0, 0
		for _, idx := range path {
			if idx < 4 { // 索引区分小时
				hour += leds[idx]
			} else {
				minute += leds[idx]
			}
		}
		if hour < 12 && minute < 60 {
			*res = append(*res, fmt.Sprintf("%d:%02d", hour, minute))
		}
		return
	}
	for i := start; i < len(leds); i++ {
		path = append(path, i)
		backtrackIndex(leds, k, i+1, path, res)
		path = path[:len(path)-1]
	}
}

func combine(n int, k int) [][]int {
	arr := []int{}
	res := [][]int{}
	path := []int{}
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}
	var backtracking func(n int, k int, start int)
	backtracking = func(n int, k int, start int) {
		if len(path) == k {
			// 必须复制 path
			res = append(res, append([]int(nil), path...))
			return
		}
		for i := start; i < n; i++ {
			path = append(path, arr[i])
			backtracking(n, k, i+1)
			path = path[:len(path)-1] // 回溯
		}
	}
	backtracking(n, k, 0)
	return res
}

// 17. 电话号码的字母组合
func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}
	wordMap := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}
	res := []string{}

	var backtrackingLetter func(index int, path string)
	backtrackingLetter = func(index int, path string) {
		if index == len(digits) {
			res = append(res, path)
			return
		}
		letters := wordMap[digits[index]]
		for i := 0; i < len(letters); i++ {
			backtrackingLetter(index+1, path+string(letters[i]))
		}
	}
	backtrackingLetter(0, "")
	return res
}

// 22. 括号生成
func generateParenthesis(n int) []string {
	path := "("
	res := []string{}
	var backtracking func(path string)
	backtracking = func(path string) {
		if len(path) == 2*n {
			if Stackmatch(path) {
				res = append(res, path)
			}
			return
		}
		backtracking(path + string("("))
		backtracking(path + string(")"))
	}
	backtracking(path)
	return res
}

func Stackmatch(s string) bool {
	cstack := CharStack{}
	for _, c := range s {
		if string(c) == "(" {
			cstack.Push(c)
		} else {
			topv, _ := cstack.Peek()
			if string(topv) != "(" {
				return false
			}
			cstack.Pop()
		}
	}
	return cstack.IsEmpty()
}

func main() {
	// board := [][]byte{
	// 	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	// }
	fmt.Println(pyramidTransition("AAAA", []string{"AAB", "AAC", "BCD", "BBE", "DEF"}))
	// solveSudoku(board)
}

type CharStack struct {
	items []rune
}

// 入栈
func (s *CharStack) Push(ch rune) {
	s.items = append(s.items, ch)
}

// 出栈
func (s *CharStack) Pop() (rune, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

// 查看栈顶
func (s *CharStack) Peek() (rune, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

// 判断是否为空
func (s *CharStack) IsEmpty() bool {
	return len(s.items) == 0
}

// 栈大小
func (s *CharStack) Size() int {
	return len(s.items)
}

// 37. 解数独
func solveSudoku(board [][]byte) {
	count := 0
	location := make(map[int][]int)
	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				location[count] = []int{i, j}
				count++
			}
		}
	}
	// 判断在第 i 行或第 j 列是否存在值 v
	match := func(i, j int, v byte) bool {
		// 检查第 i 行
		for col := 0; col < 9; col++ {
			if board[i][col] == v {
				return false // 行里已经有 v
			}
		}
		// 检查第 j 列
		for row := 0; row < 9; row++ {
			if board[row][j] == v {
				return false // 列里已经有 v
			}
		}
		startRow := (i / 3) * 3
		startCol := (j / 3) * 3
		for r := startRow; r < startRow+3; r++ {
			for c := startCol; c < startCol+3; c++ {
				if board[r][c] == v {
					return false
				}
			}
		}
		return true // 行和列都没有 v，可以放置
	}
	found := false
	var fitin func(n int)
	fitin = func(n int) {
		if found {
			return
		}
		if n == count {
			found = true
			return
		}
		nexti := location[n][0]
		nextj := location[n][1]
		for v := byte('1'); v <= '9'; v++ {
			if match(nexti, nextj, v) {
				board[nexti][nextj] = v
				fitin(n + 1)
				if found {
					return
				}
				board[nexti][nextj] = '.'
			}
		}
	}
	fitin(0)
	// for _, v := range board {
	// 	fmt.Println(string(v))
	// }

	// 这里 result 就是完整的答案路径 fmt.Println("路径:", string(result))
}

// 39. 组合总和
func combinationSum(candidates []int, target int) (ans [][]int) {
	var backtracking func(start, sum int)
	path := []int{}
	backtracking = func(start, sum int) {
		if sum > target {
			return
		}
		if sum == target {
			fmt.Println(path)
			ans = append(ans, append([]int(nil), path...))
		}
		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			backtracking(i, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	backtracking(0, 0)
	return
}

// 46. 全排列-回溯法
// 题目大意：给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
func permute(nums []int) (ans [][]int) {
	n := len(nums)
	path := []int{}
	m := make(map[int]bool)
	for i := 0; i < n; i++ {
		m[i] = false
	}
	var bc func()
	bc = func() {
		if len(path) == n {
			t := append([]int(nil), path...)
			ans = append(ans, t)
		}
		for i := 0; i < n; i++ {
			// fmt.Println(path)
			if m[i] == false {
				path = append(path, nums[i])
				m[i] = true
				bc()
				m[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	bc()
	return ans
}

// 47. 全排列 II-回溯法 排列要从0开始,多一个used数组来判断是否被选择过
// 题目大意：给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
// 输入：nums = [1,1,2]
// 输出：[[1,1,2], [1,2,1], [2,1,1]]
func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	used := make([]bool, n)
	path := make([]int, 0, n)

	var dfs func()
	dfs = func() {
		if len(path) == n {
			tmp := append([]int(nil), path...)
			ans = append(ans, tmp)
			return
		}

		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			// 1 2 2
			//相同的数递归出来的结果是一致的,比如  我第一轮选了1
			// 在 1 2 之间排列 1 2 / 2 1
			// 第二次我选择第二个1 同样是在剩下 1 2 之间排序(去掉重复结果,所以要排序)
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs()
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	dfs()
	return ans
}

// 39. 组合总和
// 题目大意：给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
// candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
// 输入：candidates = [2,3,6,7], target = 7
// 输出：[[2,2,3],[7]]
func combinationSum1(candidates []int, target int) (ans [][]int) {
	path := []int{}
	var dfs func(start, sum int, path []int)
	dfs = func(start, sum int, path []int) {
		if sum > target {
			return
		}
		if sum == target {
			t := append([]int(nil), path...)
			ans = append(ans, t)
			return
		}
		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			sum += path[len(path)-1]
			dfs(i, sum, path) // 可以被无限次选择所以从i开始
			sum -= path[len(path)-1]
			path = path[:len(path)-1]
		}
	}
	sum := 0
	dfs(0, sum, path)
	return ans
}

// 40. 组合总和 II 会有重复的方案,同理,存在同样的值需要去重,否则选出来的结果会重复
// 题目大意：给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用 一次 。
// 注意：解集不能包含重复的组合。
// 输入：candidates = [10,1,2,7,6,1,5], target = 8
// 输出：[[1,1,6],[1,2,5],[1,7],[2,6]]
func combinationSum2(candidates []int, target int) (ans [][]int) {
	//
	sort.Ints(candidates)
	path := []int{}
	var dfs func(start, sum int, path []int)
	dfs = func(start, sum int, path []int) {
		if sum > target {
			return
		}
		if sum == target {
			t := append([]int(nil), path...)
			ans = append(ans, t)
			return
		}
		// 1 1 2 5 6 7 10
		for i := start; i < len(candidates); i++ {
			// i > start 决定了第一次出现的同代表不去掉
			// 后面再出现一样的就不算了
			if i > start && candidates[i-1] == candidates[i] {
				continue
			}
			path = append(path, candidates[i])
			sum += path[len(path)-1]
			dfs(i+1, sum, path) // 同一个位置不能重复加,从下一个开始加
			sum -= path[len(path)-1]
			path = path[:len(path)-1]
		}
	}
	sum := 0
	dfs(0, sum, path)
	return ans
}

// 子集 I nums 中的所有元素 互不相同(不需要排序)
// 题目大意：给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
func subsets(nums []int) (ans [][]int) {
	path := []int{}
	var dfs func(start int, path []int)
	dfs = func(start int, path []int) {
		// 每层都加,不需要剪枝 !!!!
		t := append([]int(nil), path...)
		ans = append(ans, t)
		// 1 2 2 5 6 7 10
		for i := start; i < len(nums); i++ {
			// i > start 决定了第一次出现相同代表不去掉
			// 后面再出现一样的就不算了
			if i > start && nums[i-1] == nums[i] {
				continue
			}
			path = append(path, nums[i])
			dfs(i+1, path) // 同一个位置不能重复加,从下一个开始加
			path = path[:len(path)-1]
		}
	}
	dfs(0, path)
	return ans
}

// 子集 II
// 题目大意：给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
// 输入：nums = [1,2,2]
// 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
func subsetsWithDup(nums []int) (ans [][]int) {
	sort.Ints(nums)
	path := []int{}
	var dfs func(start int, path []int)
	dfs = func(start int, path []int) {
		// 每层都加,不需要剪枝 !!!!
		t := append([]int(nil), path...)
		ans = append(ans, t)
		// 1 1 2 5 6 7 10
		for i := start; i < len(nums); i++ {
			// i > start 决定了第一次出现的同代表不去掉
			// 后面再出现一样的就不算了
			if i > start && nums[i-1] == nums[i] {
				continue
			}
			path = append(path, nums[i])
			dfs(i+1, path) // 同一个位置不能重复加,从下一个开始加
			path = path[:len(path)-1]
		}
	}
	dfs(0, path)
	return ans
}

// 79. 单词搜索
// 题目大意：给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// 输出：true
func exist(board [][]byte, word string) bool {
	freq := map[byte]int{}
	for i := range board {
		for j := range board[0] {
			freq[board[i][j]]++
		}
	}
	for i := 0; i < len(word); i++ {
		freq[word[i]]--
		if freq[word[i]] < 0 {
			return false
		}
	}
	var dfs func(i, j, k int) bool
	// 二维数组要次次make
	used := make([][]bool, len(board))
	for i := range used {
		used[i] = make([]bool, len(board[0]))
	}

	// i,j位置的数值是否是word[k]
	dfs = func(i, j, k int) bool {
		// 1. 越界
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
			return false
		}
		// 2. 已访问
		if used[i][j] {
			return false
		}
		// 3. 字符不匹配
		if board[i][j] != word[k] {
			return false
		}
		// 4. 匹配到最后一个字符
		if k == len(word)-1 {
			return true
		}
		// 5. 标记访问
		used[i][j] = true
		// 6. 四方向搜索，只要一个成功就 true
		if dfs(i+1, j, k+1) || dfs(i-1, j, k+1) ||
			dfs(i, j+1, k+1) || dfs(i, j-1, k+1) {
			return true
		}
		// 7. 回溯
		used[i][j] = false
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// 分割回文串
// 题目大意：给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]
func partition(s string) (ans [][]string) {

	var isPalindrome func(s string, left, right int) bool
	isPalindrome = func(s string, left int, right int) bool {
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}
	path := []string{}
	var dfs func(end int)
	dfs = func(end int) {
		if end == len(s) {
			tm := make([]string, len(path))
			copy(tm, path)
			ans = append(ans, tm)
		}
		for i := end; i < len(s); i++ {
			if isPalindrome(s, end, i) {
				path = append(path, s[end:i+1])
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return
}

// 756. 金字塔转换矩阵
func pyramidTransition(bottom string, allowed []string) bool {
	n := len(bottom)
	m := make(map[[2]byte][]byte)
	for i := range allowed {
		t := [2]byte{allowed[i][0], allowed[i][1]}
		m[t] = append(m[t], []byte(allowed[i][2]))
	}
	board := make([]string, len(bottom))
	board[0] = bottom

	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {

	}
	return false
}
