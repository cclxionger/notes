package dynamic

import (
	"fmt"
	"sort"
)

func CanPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 || len(nums) == 1 {
		return false
	}
	sum /= 2
	dp := make([]int, sum + 10)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		for j := sum; j >= 0; j-- {
			if j-val >= 0 && dp[j-val] == 1 {
				dp[j] = 1
			}
		}
	}
	fmt.Println(dp)
	// temp := 0
	// for i := 0; i < len(nums); i++ {
	// 	dp[nums[i]] = 1

	// 	for j := i; j < len(nums); j++ {
	// 		if dp[j] != 0 {
	// 			temp += nums[j]
	// 		}
	// 		dp[temp] = 1
	// 	}
	// }
	// fmt.Println(dp)
	if dp[sum] == 1 {
		return true
	}
	return false
}

// 这个方法在1，1，2，2中 不行
// 因为这样会4，2   应该1+2，1+2，返回true
func canPartition1(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return false
	}
	sort.Ints(nums)
	i := 0
	j := n - 1
	numL := nums[0]
	numR := nums[n-1]
	for {
		if i+1 == j {
			break
		}
		if numL <= numR {
			i++
			numL += nums[i]
		} else {
			j--
			numR += nums[j]
		}
	}
	if i+1 == j && numL == numR {
		return true
	}
	return false
}
