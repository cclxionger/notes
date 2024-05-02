package remove
////leetcode 27
func RemoveElement(nums []int, val int) int {//返回移除后数组的新长度
	slow := 0//慢指针，就是为了获得新数组的下标
	//快指针是为了获取新数组的值
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow ++
		}
	}
	return  slow //因为索引从1开始，slow当前索引=新数组长度+1
}