package main

import (
	"container/heap"
	"fmt"
	"math"
	"slices"
	"sort"
)

func main() {
	fmt.Print(longestPalindrome("babad"))
	fmt.Print(longestPalindrome("cbbd"))

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

// 1351. 统计有序矩阵中的负数
func countNegatives(grid [][]int) (ans int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] < 0 {
				ans++
			}
		}
	}
	return
}

// 1351. 统计有序矩阵中的负数
func countNegativesSearch(grid [][]int) (ans int) {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		//对于每一行进行二分查找
		//3,2,1,-1 查找-1的位置
		x := sort.Search(len(grid[i]), func(k int) bool {
			return grid[i][k] < 0
		})
		ans += n - x

	}
	return ans
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
func combinationSum(candidates []int, target int) (ans [][]int) {
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

// 5. 最长回文子串
func longestPalindrome(s string) string {
	reslen := 0
	resstart := 0
	for i := 0; i < len(s); i++ {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > reslen {
				resstart = l
				reslen = r - l + 1
			}
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > reslen {
				resstart = l
				reslen = r - l + 1
			}
			l--
			r++
		}
	}
	return s[resstart : resstart+reslen]
}

// 840. 矩阵中的幻方
func numMagicSquaresInside(grid [][]int) (ans int) {
	if len(grid) < 3 || len(grid[0]) < 3 {
		return 0
	}

	isvalid := func(i, j int) bool {
		seen := make([]bool, 10)
		for r := 0; r < 3; r++ {
			for c := 0; c < 3; c++ {
				v := grid[i+r][j+c]
				if v < 1 || v > 9 || seen[v] {
					return false
				}
				seen[v] = true
			}
		}
		for v := 1; v <= 9; v++ {
			if !seen[v] {
				return false
			}
		}
		target := grid[i][j] + grid[i][j+1] + grid[i][j+2]
		for r := 0; r < 3; r++ {
			if grid[i+r][j]+grid[i+r][j+1]+grid[i+r][j+2] != target {
				return false
			}
		}
		for c := 0; c < 3; c++ {
			if grid[i][j+c]+grid[i+1][j+c]+grid[i+2][j+c] != target {
				return false
			}
		}
		if grid[i][j]+grid[i+1][j+1]+grid[i+2][j+2] != target {
			return false
		}
		if grid[i][j+2]+grid[i+1][j+1]+grid[i+2][j] != target {
			return false
		}
		return true
	}

	for i := 0; i <= len(grid)-3; i++ {
		for j := 0; j <= len(grid[0])-3; j++ {
			if isvalid(i, j) {
				ans++
			}
		}
	}
	return ans
}
