package main

import (
	"fmt"
	"strings"
)

func minDistance(P, T string) (int, [][]int) {
	m, n := len(P), len(T)
	//初始化dp
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 初始化DP表第0列和第0行
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// 填充DP表
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if P[i-1] == T[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// 返回最小差别数和DP表
	return dp[m][n], dp
}

// 三个数字的最小值
func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

func main() {

	P := "abc"
	T := "ab"
	fmt.Println("P为", P)
	fmt.Println("T为", T)
	minDiff, dpTable := minDistance(P, T)
	fmt.Println("表中数据如下:")
	//打印输出
	for _, row := range dpTable {
		fmt.Println(strings.Join(strings.Fields(fmt.Sprint(row)), " "))
	}
	fmt.Printf("二者最小差别数:%d\n", minDiff)
}
