package main

import (
	"encoding/json"
	"fmt"
)

// 可视化 相当于public 在其他包可以导入这个包，ID和Name大写属性大写
type student struct {
	ID   int    `json:"id"` //结构体标签 后面对应的是字符类型
	Name string `json:"name"`
}
type class struct {
	Title    string
	Students []student
}

// student的构造函数
func NewStudent(id int, name string) student {
	return student{
		ID:   id,
		Name: name,
	}
}

/*
因为slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意
正确的做法是在方法中使用传入的slice的拷贝进行结构体赋值。
func (p *Person) SetDreams(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}
*/

func main() {
	//创建一个班级变量
	c3 := class{
		Title:    "计科03",
		Students: make([]student, 0, 20),
	}
	for i := 0; i < 10; i++ {
		stu := NewStudent(i, fmt.Sprintf("stu%02d", i))
		c3.Students = append(c3.Students, stu)
	}
	fmt.Printf("%v\n", c3)
	//json序列化：go数据-》json格式字符串
	data, err := json.Marshal(c3)
	if err != nil {
		fmt.Println("出错了", err)
		return
	}
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)
	fmt.Println()
	//反序列化
	jsonStr := `{"Title":"计科03","Students":[{"ID":0,"Name":"stu00"},
	{"ID":1,"Name":"stu01"},{"ID":2,"Name":"stu02"},{"ID":3,"Name":"stu03"},
	{"ID":4,"Name":"stu04"},{"ID":5,"Name":"stu05"},{"ID":6,"Name":"stu06"},
	{"ID":7,"Name":"stu07"},{"ID":8,"Name":"stu08"},{"ID":9,"Name":"stu09"}]}`
	var c2 class
	err = json.Unmarshal([]byte(jsonStr), &c2) //得用指针进行修改
	fmt.Println(jsonStr)
	if err != nil {
		fmt.Println("反序列化出错了")
		return
	}
	fmt.Printf("%#v\n", c2)

}
