package binsea

func Sear1(arr []int, target int) int { //左闭右闭 ,返回target在切片中的索引
	left := 0
	right := len(arr) - 1
	for {
		
		if !(left <= right) { //退出循环
			break
		}
		mid := (left + (right - left)) >> 1
		if arr[mid] > target { //数组中的mid大了，往左边寻找
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			return mid
		}

	}
	return -1 //查找不到返回-1
}
func Sear2(arr []int, target int) int { //左闭右开 ,返回target在切片中的索引
	left := 0
	right := len(arr) //left在范围内，right不在
	for {
		if !(left < right) { //退出循环
			break
		}
		mid := (left + right) >> 1 //小心left+right超范围，尽量写成（left+（right-left）） / 2
		if arr[mid] > target {     //数组中的mid大了，往左边寻找
			right = mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			return mid
		}

	}
	return -1 //查找不到返回-1
}
