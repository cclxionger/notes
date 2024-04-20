package quick1

func QuickSort(n int, arr []int) {
	quick(0, n-1, arr)
}
func quick(left, right int, arr []int) {
	if left >= right { //当left=right，也就是只有一个元素的时候，结束递归
		return
	}
	index := partition(left, right, arr)
	quick(left, index-1, arr)  //记得是循环quick
	quick(index+1, right, arr) //记得是循环quick
}
func partition(left, right int, arr []int) int {
	i := left
	j := right
	element := arr[left] //基准元素
	for i < j {
		for i < j && arr[j] > element {
			j--
		}
		for i < j && arr[i] <= element {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	} //此时i=j
	arr[left], arr[j] = arr[j], arr[left]
	return i
}
