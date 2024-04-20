package main

import "fmt"

func main() {
	//创建方式1
	var sd1 student
	fmt.Println(sd1)
	sd1.name = "ccl"
	sd1.age = 19
	fmt.Println(sd1)
	//创建方法2
	sd2 := student{"ck", 20}
	fmt.Println(sd2)
	//指针创建方法3
	var sd3 *student = new(student)
	(*sd3).name = "lc" //一定要+一个括号，因为.的优先级比*高，因为指针赋值，因此标准的通过指针赋值
	sd3.age = 19       //为了方便也可以这样赋值，底层做取值处理
	fmt.Println(sd3)
	//指针创建方法4
	var std4 *student = &student{} //这个大括号里面，可以直接赋值
	std4.name = "amy"              //同上面
	std4.age = 6
	fmt.Println(std4)

	var i inter = 0
	fmt.Println(i)

}

type student struct { //结构体是值传递
	name string
	age  int
}
type inter int //type后面不仅可以接struct，还可以接基本类型
