package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World！")
	var ( //定义多个变量
		a, b int
		name string
	)
	fmt.Println(name) //默认为"",如果下面不用到这个变量，会保存
	temp := 6         //自动推导变量类型
	fmt.Println(temp)
	fmt.Printf("%T", temp) //求对应变量的类型
	// fmt.Println("%f", &temp) //求变量的地址
	a = 1
	b = 2
	fmt.Println(a, b)
	//交换 a和b的值
	a, b = b, a
	fmt.Println(a, b)
	const URL string = "www.baidu" //常量显示定义
	const URL2 = "www.wangyiyun"   //常量隐式定义
	const q = 3.14                 //定义多个常量

	fmt.Println(URL, URL2, q)
	const (
		w = iota //从0开始
		e        //不写默认=iota，依次增加
	)
	const ( //每次都从0重新开始
		r = iota
		s
	)
	fmt.Println(w, e, r, s)
	var f float64 = float64(s) //转变数据类型
	y := float64(s)
	fmt.Printf("%.2f+%.2f\n", f, y)

	//进行除法运算，要换成float
	o := 10
	p := 3
	result := float64(o) / float64(p)
	fmt.Println(result)
	// var ccl string
	// fmt.Scan(&ccl)
	// fmt.Println(ccl)
	// var h int
	// fmt.Scanln(&h)
	// fmt.Println(h)
	// if h > 0 {
	// 	fmt.Println("你所输入的h大于0")
	// }
	// switch h {
	// case 1:
	// 	fmt.Println("h大于0并且是1")
	// 	fallthrough //有这个case是穿透的，不管下一个条件是否满足
	// case 2:
	// 	fmt.Println("h大于0并且是2")
	// default:
	// 	fmt.Println("h属于其他情况")
	// } //switch语句结束
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	str := "helloWorld"
	fmt.Println(len(str))
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}
	fmt.Println()
	// str[1] = "666"字符串某个数组不能修改，但是整体可以修改
	fmt.Println(str)
	//string不能修改
	for i, v := range str { //for range循环，遍历数组，切片...
		fmt.Print(i)
		fmt.Printf(" %c\t", v)
	}
	fmt.Println()
	fmt.Println(add(1, 3))
	dajj := []int{1, 3, 4}
	fmt.Println(dajj)
	prin(dajj)
	fmt.Println(dajj)
	//如果一个函数调用的时候，前面+了 defer，那么就会最后调用（在程序执行完），相当于栈
	fmt.Printf("%t\n", essence)
	fmt.Println(essence) //函数名字指向函数体内存地址，一种特殊类型的指针变量
	fmt.Println(essence(1, 2))
	var essence1 func(int, int) (int, int) //函数本质
	essence1 = essence                     //把地址传给essence1
	fmt.Println(essence1(1, 2))            //essence1与essence有同样的功能
	anonymous := func() {                  //匿名函数
		fmt.Println("我是匿名函数")
	}
	anonymous()      //调用
	func(a, b int) { //匿名函数 自己调用自己
		fmt.Println("我能输出")
	}(a, b) //还可以增加返回值，这里不举例了

}
func add(a, b int) int { //值传递：基础类型，array，struct
	c := a + b
	return c
}
func prin(a []int) { //引用传递，slice,map, chan，切片
	a[0] = 100
}
func essence(a, b int) (int, int) { //函数传参还可以传入参数
	return a + b, a - b
}
