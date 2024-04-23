package main

import (
	"fmt"
	towpointer "towpointer/towPointer"
)

func main() {
	// nums := []int{-4, -1, 0, 3, 10}
	// result := make([]int,len(nums),len(nums))
	// result = towpointer.SortedSquares(nums)
	// fmt.Println(result)
	arr2 := []int {2,3,1,2,4,3}
	n := towpointer.MinSubArrayLen(7,arr2)
	fmt.Println(n)
}
