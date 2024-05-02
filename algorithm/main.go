package main

import (
	basic "algorithm/basic"
	"fmt"
)

func main() {
	a := basic.Constructor()
	a.AddAtHead(1)
	a.AddAtHead(2)
	a.AddAtTail(0)
	
	a.DeleteAtIndex(1)
	a.AddAtIndex(0,10)
	fmt.Println(a.Get(0))
	fmt.Println(a.Get(1))
	fmt.Println(a.Get(2))
	
}
