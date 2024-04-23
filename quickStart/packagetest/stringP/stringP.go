package stringp

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	type Stringer interface {
	    String() string
	}
*/
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string { //实现Stringer//相当于java中的toString
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)

}

type IP struct {
	Arr []int
}

// 进行练习
func (ip IP) String() string {
	temp := make([]string, len(ip.Arr))

	for i, value := range ip.Arr {
		temp[i] = strconv.Itoa(value) // 使用 strconv.Itoa 将整数转换为字符串
		if i != len(ip.Arr)-1 {
			temp[i] += "." // 在非最后一个元素后添加点
		}
	}

	result := strings.Join(temp, "") // 利用 strings.Join 将 temp 转换成字符串类型

	return result // 直接返回结果字符串
}
