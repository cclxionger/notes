package main

import (
	"fmt"
	"sort"
)

func main() {
	//第一种声明方式
	var a map[string]string
	a = make(map[string]string) //在使用map前面必须要进行make进行分配空间
	//第三种 声明方式
	a["no1"] = "松江"
	fmt.Println(a) //并且输出无序，不按照key，也不按照value
	//第二种 声明方式
	b := make(map[string]string)
	b["no1"] = "松江"
	fmt.Println(b)
	c := map[string]string{
		"1": "theshy", //记得在最后写一个逗号
	}
	fmt.Println(c)
	//删除一个
	delete(c, "1")
	fmt.Println(c)
	//如果要把一个map里面的所有都删除，遍历key  或者直接make一个新的空间
	b["2"] = "c"
	fmt.Println(b)
	b = make(map[string]string, 5)
	fmt.Println(b)
	//查找
	value, flag := a["no1"] //第一个参数为key对应的value，第二个返回一个bool
	if flag == true {
		fmt.Print("有,")
		fmt.Print("值为")
		fmt.Println(value)
	}
	//遍历  不能用for循环，只能用for range  好像新版直接是按照添加 元素位置进行遍历
	a["no2"] = "ning"
	a["no3"] = "dada"
	a["no4"] = "zhendeyouxvme"
	for k, v := range a {
		fmt.Printf("k=%v,v=%v\n", k, v)
	}
	//也可以把map中的key添加到切片中进行遍历
	keys := []string{}
	for k, _ := range a {
		keys = append(keys, k)
	}
	fmt.Println(keys) //无序
	sort.Strings(keys)
	fmt.Println(keys) //排序后
	for _, k := range keys {
		fmt.Print(k)
		fmt.Printf("对应的value为%v\n", a[k])
	}
	fmt.Println()
	//map切片  切片类型是map
	// var mapSlice = make([]map[string]string, 3)
	// for index, value := range mapSlice {
	// 	fmt.Printf("index:%d value:%v\n", index, value)
	// }
	// mapSlice[0] = make(map[string]string)
	// mapSlice[0]["name"] = "小王子"
	// mapSlice[0]["password"] = "123456"
	// mapSlice[0]["address"] = "海淀区"
	// for index, value := range mapSlice {
	// 	fmt.Printf("index:%d value:%v\n", index, value)
	// }


}
