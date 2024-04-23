package main

import (
	"binarysea/binsea"
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	index := binsea.Sear1(arr, 1)
	fmt.Println(index)
	index = binsea.Sear2(arr, 5)
	fmt.Println(index)
}
