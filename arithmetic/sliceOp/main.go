package main

import (
	"fmt"
	"seliceop/remove"
)

func main() {
	arr := []int{1, 2, 3, 3, 5}
	l := remove.RemoveElement(arr, 3)
	fmt.Println(l)
	

	
}
