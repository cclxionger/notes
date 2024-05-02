package main

import (
	"fmt"
	"packaging/pack"
)

func main() {
	p := pack.NewPerson("666")
	fmt.Println(p)
	p.Name = "777"
	fmt.Println(p.Name)
	p.Setage(10)
	fmt.Println(p.Getage())

	
}
