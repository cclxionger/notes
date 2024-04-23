package main

import (
	"fmt"
	"math_func/fibonacci"
)

func main() {
	//不知道为什么不能这样写
	// for i := 0; i < 10; i++ {
	// 	a := fibonacci.Fib(i)
	// 	fmt.Printf("%d\t",a)
	// }
	fmt.Println(fibonacci.Fib1(5))
	fmt.Println(fibonacci.Fib2(5))

}
