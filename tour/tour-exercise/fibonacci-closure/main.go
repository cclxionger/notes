package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

// 0 1 1 2 3 5 8
func fibonacci() func(i int) string {
	a, b := 0, 1
	return func(i int) string {
		if i == 1 {
			return fmt.Sprintf("%d次斐波那契数为%d", i, a)
		}
		a, b = b, a+b
		return fmt.Sprintf("%d次斐波那契数为%d", i, a)
	}
}

func main() {
	f := fibonacci()
	for i := 1; i <= 10; i++ {
		fmt.Println(f(i))
	}
}
