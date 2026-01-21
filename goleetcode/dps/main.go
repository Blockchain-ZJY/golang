package main

import "fmt"

func Getmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp, 获取最大子序列长度 [1,5,2,4,3]
func dp(arr []int, i int, memo map[int]int) int { // 从第i个元素开始，计算最大子序列长度
	if i == len(arr)-1 {
		return 1
	}
	max := 1
	for j := i + 1; j < len(arr); j++ {
		if arr[j] > arr[i] {
			max = Getmax(max, dp(arr, j, memo)+1)
		}
	}
	memo[i] = max
	return max

}

func main() {
	// fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 1, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}}))
	fmt.Println(countSubstrings("aaa"))
	// fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}

// 70. 爬楼梯
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[len(dp)-1]
}

// [10,9,2,5,3,7,101,18]
// 300. 最长递增子序列
func lengthOfLIS(nums []int) (res int) {
	dp := []int{}
	// dp[i] 定义为 以i结尾的最长递增子序列的长度
	for i := 0; i < len(nums); i++ {
		dp = append(dp, 1)
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = Getmax(dp[j]+1, dp[i])
				res = Getmax(res, dp[i])
			}
		}
	}
	return
}

// 5. 最长回文子串 递归实现 dp[i,j]代表 s[i,j+1]是否是回文
func longestPalindrome(s string) string {
	resl, resr := 0, 0
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
				if dp[i][j] {
					if j-i > resr-resl {
						resr, resl = j, i
					}
				}
			}
		}
	}
	return s[resl : resr+1]
}

// 746. 使用最小花费爬楼梯
// dp[i] 表示到达下标i最小花费值
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i < n+1; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	fmt.Println(dp)
	return dp[n]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 62. 不同路径
// 数学计算 C_m+n-2 ^m-1
func uniquePaths(m int, n int) int {
	a := m + n - 2
	b := m - 1
	if n-1 < b {
		b = n - 1
	}
	pr := 1
	for i := 0; i < b; i++ {
		// 先乘后除（确保每一步都是整数，减少中间值大小，缓解溢出）
		pr = pr * (a - i) / (i + 1)
	}
	return pr
}

// 62. 不同路径
// dp计算 dp[i][j] 表示从0,0到i,j位置可能性数
func uniquePathsDP(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for j := 0; j < m; j++ {
		dp[j][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 63. 不同路径 II
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		dp[0][j] = 1
	}
	for j := 0; j < m; j++ {
		if obstacleGrid[j][0] == 1 {
			break
		}
		dp[j][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

//343. 整数拆分
// dp[i] 表示 i这个数 被拆分最大乘积是多少
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 0, 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//647. 回文子串
// 与第五题类似, 用dp[i][j] 表示从s[i,j+1]是否是回文串
func countSubstrings(s string) (ans int) {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
					ans++
				} else {
					dp[i][j] = dp[i+1][j-1]
					if dp[i][j] {
						ans++
					}
				}
			}
		}
	}
	return
}
