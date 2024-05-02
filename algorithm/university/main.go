package main

import (
	numberstower "algorithm/university/numbersTower"
	"fmt"
)

func main() {
	arr := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}

	n := numberstower.Numberstower(arr)
	fmt.Println(n)

}
