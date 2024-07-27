package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := make([]int, 2)
	a[0] = 1
	a[1] = 1
	a = append(a, 1)
	a = append(a, 1)
	a = append(a, 1)
	a = append(a, 1)
	a = append(a, 1)

	fmt.Println(a)
	fmt.Println(unsafe.Alignof(a))
	fmt.Println(666)
}
