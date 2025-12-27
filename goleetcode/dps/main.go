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
	// memo := make(map[int]int)
	// arr := []int{1, 5, 2, 4, 3, 12, 20, 43, 12, 576, 42, 41, 78, 5178, 21, 1, 89, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	// max := -1
	// for i := 0; i < len(arr); i++ {
	// 	max = Getmax(max, dp(arr[:], i, memo))
	// }
	// fmt.Println(max)
	fmt.Println(climbStairs(4))
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
