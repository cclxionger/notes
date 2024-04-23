package main

import (
	"fmt"
	errorp "packageTest/errorp"
	"packageTest/stringp"
)

func main() {
	a := stringp.Person{"ccl Dent", 42}
	z := stringp.Person{"ck Beeblebrox", 9001}
	fmt.Println(a)
	fmt.Println(z)
	ipArr := stringp.IP{Arr: []int{1, 2, 3, 4, 5}} //结构体字面量
	fmt.Println(ipArr.String())                    //也可以写成ipArr，底层实现
	c := errorp.ErrSqrt(-6)
	fmt.Println(c)
}
