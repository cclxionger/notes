package dynamic

import (
	"math"
)

func MaxSubArray(nums []int) int {
	n := len(nums)
	//dp[i]表示以[i]为结尾的当前最大子数组和
	//状态转移方程主要是看后面这个数字对前面数字有没有削弱作用，选不选第i这个数字
	//dp[i]不选前面的数字的话就是为nums[i], 选了前面的数字话dp[i-1] + nums[i]
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
	}
	res := math.MinInt32
	for _, v := range dp {
		res = max(v, res)
	}
	return res
}
