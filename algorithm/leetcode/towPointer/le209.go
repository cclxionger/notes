package towpointer

import "math"

// lettcode 209
func MinSubArrayLen(target int, nums []int) int {
	var n = len(nums)
	result := n + 1//不初始化为n，因为n可能刚好就满足条件
	//使用滑动窗口，思路就是一直向右，直到大于=target，
	//滑动窗口向右移动左指针进行向右移动
	//
	res := 0//当前总和
	le := n + 1//当前长度  
	for i,j := 0,0; j < n; j++ {//j代表终止位置
		res += nums[j]
		for res >= target{//当 当前总和大于等于目标值sum之后（因为nums中可以有0）。进行滑动，
			le = j - i + 1
			result = int (math.Min(float64(le),float64(result)))
			res -= nums[i]
			i ++
		}
	}
	if  result == n + 1{
		return 0
	}
	return result
}