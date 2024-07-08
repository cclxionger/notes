package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"name"`
	Age  int
}

func PareseJson(obj any) {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	fmt.Println(v)
	n := v.NumField()
	for i := 0; i < n; i++ {
		//获取每个子段的变量名和标签
		fmt.Println(t.Field(i).Name, t.Field(i).Tag)
		if t.Field(i).Tag.Get("json") == "name" {//如果没有json = name,会返回一个""
			fmt.Println(666)
		}
		//获取每个子段的值
		fmt.Println(v.Field(i))
	}
}
func main() {
	student := Student{
		"ck", 19,
	}
	PareseJson(student)

}
