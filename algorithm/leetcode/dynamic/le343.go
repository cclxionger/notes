package dynamic

// leetcode 343
// 可以先写2-10找规律  n从2开始
/*     n
2：1    1*1
3：2	1*2   而不是1*dp[2]大
4：4	2*2
5：6	2*3
6：9	3*3
7：12	2*5=2*6=4*3=12相当于把5转换成dp[5]了  	1*6 = 1*9
8：
9：
10:	3* 7 = 3 * 12 = 36
所以状态转移方程：dp【i】 = j * (i-j)   或者j * dp[i-j]
然后这个随着j增加,dp[i]也可以增加，随意要比较dp[i]之前的值和当前dp[i]的值
*/
func integerBreak(n int) int {
	//dp[i] 表示数字i 对应分解后的最大值
	dp := make([]int, n+1)
	//初始化dp
	dp[2] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			//然后这个随着j增加,dp[i]也可以增加，随意要比较dp[i]之前的值和当前dp[i]的值

			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}
func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
