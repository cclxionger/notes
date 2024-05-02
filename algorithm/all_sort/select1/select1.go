package select1
//选择排序，每次比较相邻的俩个，最大的放在最后
func SelectSort(n int, arr []int) { //长度为n的切片
	for i := 0; i < len(arr) - 1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[j + 1] {
				arr[i],arr[j] = arr[j],arr[i]
			}
		}
	}
}
