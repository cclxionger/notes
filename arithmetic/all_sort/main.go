package main

import (
	"all_sort/bubble"
	"fmt"
)

func main() {
	var arr = [...]int{5, 3, 1, 6}
	bubble.BubbleSort(len(arr), arr[:])
	//
	fmt.Println(arr)

}
