// 不同路径 只能向右或者向下走  有障碍
package dynamic

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)    //m行
	n := len(obstacleGrid[0]) //n列
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	//起始位置或者终点位置有障碍物，直接返回0
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	for i := 0; i < m && obstacleGrid[i][0] != 1; i++ { //遍历每一行初始化
		dp[i][0] = 1
	}
	for j := 0; j < n && obstacleGrid[0][j] != 1; j++ { //遍历每一列初始化
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

//进行优化
func UniquePathsWithObstacles2(obstacleGrid [][]int) int {
	m := len(obstacleGrid)    //m行
	n := len(obstacleGrid[0]) //n列
	dp := make([]int, n)
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	for j := 0; j < n; j++ { //初始化第一行
		if obstacleGrid[0][j] == 1 {
			break
		}
		dp[j] = 1
	}
	for i := 1; i < m; i++ { //从第一行开始
		if obstacleGrid[i][0] == 1 { //每次都初始化每一行的第一列
			dp[0] = 0
		}
		for j := 1; j < n; j++ {

			if obstacleGrid[i][j] == 0 {
				dp[j] += dp[j-1]
			} else {
				dp[j] = 0
			}

		}
	}

	return dp[n-1]
}
