package main

import (
	"all_sort/insert"
	"fmt"
)

func main() {
	var arr = [...]int{5, 3, 1, 6}
	sliceArr := make([]int,0)
	sliceArr = append(sliceArr, 0)
	for i := 0; i < len(arr); i++ {
		sliceArr = append(sliceArr, arr[i])
	}
	fmt.Println(sliceArr)
	insert.InsertSort1(len(arr), sliceArr)
	fmt.Println(sliceArr)
	//
	fmt.Println(arr)

}
