// 爬楼梯
package dynamic

//每次你可以爬 1 或 2 个台阶
//dp【i】表示爬到i个台阶的方法数量
func ClimbStairs(n int) int {
	dp := make([]int, 3)
	dp[1] = 1
	if n == 1 {
		return 1
	}
	dp[2] = 2
	if n == 2 {
		return 2
	}
	//进行状态压缩
	sum := dp[1] + dp[2]
	for i := 4; i < n+1; i++ {
		dp[1] = dp[2]
		dp[2] = sum
		sum = dp[1] + dp[2]
	}
	return sum
}
