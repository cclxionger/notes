package main

import "fmt"

func main() {
	//数组
	var a [3]int
	a[0] = 1
	fmt.Println(a[0])
	//切片Slice
	months2 := []int{1, 2, 3}
	months := []int{1: 1, 2: 2, 3: 3, 4: 4, 5: 6} //没有写0，默认是定义的类型，这个是int
	fmt.Println(months2)
	fmt.Println(months)
	spring := months[1:5] //前开后闭 ， 从1到4
	fmt.Println(spring)
	fmt.Println(spring[:3]) //从0到2
	reverse(months2)
	fmt.Println(months2)
	reverse(months2[:2]) //0-1索引反转
	fmt.Println(months2)
	/*
		内置的make函数创建一个指定元素类型、长度和容量的slice。
		容量部分可以省略，在这种情况下，容量将等于长度。
		make([]T, len)
		make([]T, len, cap) // same as make([]T, cap)[:len]
	*/
	s1 := make([]int, 5, 8)
	fmt.Println(s1)
	/*
		如果你需要测试一个slice是否是空的，使用len(s) == 0来判断，而不应该用s == nil来判断。
		除了和nil相等比较外，一个nil值的slice的行为和其它任意0长度的slice一样；
		例如reverse(nil)也是安全的。
		除了文档已经明确说明的地方，所有的Go语言函数应该以相同的方式对待nil值的slice和0长度的slice。*/
	s2 := make([]int, 0, 0)
	fmt.Println(s2)
	fmt.Println(len(s2) == 0)
	fmt.Println(s2 == nil)
	s2 = append(s2, 2, 3)
	fmt.Println(s2) 

}
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
