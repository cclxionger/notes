package main

import (
	"fmt"
	"packaging/pack"
)

func main() {
	pack.NewPerson("6")
	pack.NewPerson("666")
	pack.NewPerson("5da")
	p := pack.NewPerson("666")
	fmt.Println(p)
	p.Name = "777"
	fmt.Println(p.Name)
	p.Setage(10)
	fmt.Println(p.Getage())

}
