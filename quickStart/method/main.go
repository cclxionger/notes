package main

import "fmt"

func main() {
	var p1 Person
	p1.age = 10
	fmt.Println(p1.age) //值传递，不会影响
	fmt.Println(p1.age) //值传递，不会影响
	fmt.Println(p1.age) //值传递，不会影响
	fmt.Println(p1.age) //值传递，不会影响

}

type Person struct {
	age int
}

func (p Person) test() {
	p.age = 6
	fmt.Println(p.age)
}
