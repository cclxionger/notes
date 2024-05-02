package dynamic
//斐波那契数
func Fib(n int) int {
	//一开始可能用用一个长度为n的切片，然后进行状态压缩
	dp := make([]int, 2)
	dp[0] = 0
	if n == 0 {
		return 0
	}
	dp[1] = 1
	if n == 1 {
		return 1
	}
	//进行状态压缩
	sum := dp[0] + dp[1]
	for i := 3; i < n+1; i++ {
		dp[0] = dp[1]
		dp[1] = sum
		sum = dp[0] + dp[1]
	}
	return sum

}
