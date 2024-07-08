package main

import "fmt"

// 单指针快排
func quickSort(arr []int) {
	split(arr, 0, len(arr)-1)
}
func split(arr []int, left, right int) {
	// 如果left<right,把基准元素小或者大的一分为二
	// 直至left = right 也就是只有一个元素的时候就不分割了
	if left < right {
		pivotIndex := partition(arr, left, right)
		// 是递归split方法
		split(arr, left, pivotIndex-1)
		split(arr, pivotIndex+1, right)

	}
}
func partition(arr []int, left, right int) int {
	// 举例  1 5 3 0 -1
	// 交换后1 0 -1 5 3
	// 此时swapIndex 指向5，应该把-1和1交换位置，所以把swap(pivotIndex, swapIndex-1)
	pivotIndex := left
	swapIndex := left + 1
	for i := swapIndex; i <= right; i++ {
		if arr[pivotIndex] > arr[i] {
			swap(arr, i, swapIndex)
			swapIndex++
		}
	}
	swap(arr, pivotIndex, swapIndex-1)
	return swapIndex - 1
}

// 交换切片俩个元素
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func main() {
	arr := []int{2, 5, 0, -1, 3}
	fmt.Println(arr)
	quickSort(arr)
	fmt.Println(arr)
}
