package dynamic

import "fmt"

//不同路径 只能向右或者向下走
func UniquePaths(m int, n int) int {
	dp := make([][]int, m) //注意初始化二位切片方法
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ { //第一列值全为1,遍历每一行
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ { //第一行值全为1,遍历每一列
		dp[0][j] = 1
	}
	//状态方程推导：
	//如果（1，1）dp【1】【1】 = 2，dp【0】【2】 = 1
	//dp【1】【2】= 1+2 = 3
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]

}

//优化后
func UniquePaths2(m int, n int) int {
	dp := make([]int, n)
	// if m == 1 || n == 1 {
	// 	return 1
	// }
	
	//状态方程推导：
	//如果（1，1）dp【1】【1】 = 2，dp【0】【2】 = 1
	//dp【1】【2】= 1+2 = 3
	for j := 0; j < n; j++ {//初始化第一行
		dp[j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	fmt.Println(dp)
	return dp[n-1]

}
