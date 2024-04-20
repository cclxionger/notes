package pack

import "fmt"

type person struct {
	Name string
	age  int
}
func NewPerson(name string) *person { //工厂模式，相当于构造函数
	return &person{
		Name: name,
	}
}
func (p *person) Setage(age int) { //age因为是小写，相当于是 private，所以要set和get
	if age > 0 && age < 100 {
		p.age = age
	} else {
		fmt.Println("年龄不合理")
	}
}
func (p *person) Getage() int {
	return p.age
}
