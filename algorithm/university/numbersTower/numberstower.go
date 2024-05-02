package numberstower

func Numberstower(arr [][]int) [][]int {
	n := len(arr)
	//初始化dp
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, i+1)
	}
	for j := 0; j <= n-1; j++ {
		dp[n-1][j] = arr[n-1][j]
	}
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[i][j] = arr[i][j] + max(dp[i+1][j], dp[i+1][j+1])
		}

	}
	return dp
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
