package main

import (
	"fmt"
)

func main() {
	//将i进行断言
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)
	s, ok := i.(string) //判断i的类型是否为string
	fmt.Println(s, ok)  //如果是，返回是true
	f, ok := i.(float32)
	fmt.Println(f, ok) //flase
	// d := i.(float32) //报错，要有俩个接受值
	// fmt.Println(d)

	//类型选择 是一种按顺序从几个类型断言中选择分支的结构。
	//类型选择中的声明与类型断言 i.(T) 的语法相同，
	//只是具体类型 T 被替换成了关键字 type。

	switch v := i.(type) {
	case int:
		fmt.Printf("二倍的 %v 是 %v\n", v, v*2)
	case string:
		fmt.Printf("%q 长度为 %v 字节\n", v, len(v))

	default:
		fmt.Printf("我不知道类型 %T!\n", v)
	}
	
	

}
