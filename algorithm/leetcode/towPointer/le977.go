package towpointer

//leetcode 977
func SortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	k := n - 1
	for i, j := 0, n-1; i <= j; { //i是左指针 j是右指针
		left := nums[i] * nums[i]
		right := nums[j] * nums[j]
		if left < right { //大的就放进新切片
			result[k] = right
			j--
			k--
		} else {
			result[k] = left
			i++
			k--
		}
	}
	return result
}
