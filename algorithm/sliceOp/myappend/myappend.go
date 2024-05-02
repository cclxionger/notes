package myappend

import "fmt"

func MyAppend() {  // s1和s2分别从s中切出一部分，然后s1修改，不改变s2
	s := []int{1, 2, 3, 4, 5}
	s1 := s[:3]
	s2 := s[3:]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println("俩个切片长度和容量分别为")
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
	s1 = append(s1[:len(s1):len(s1)], 7)//核心在这样，append的时候，【里面传参容量要和长度相同】
	fmt.Println(s1)
	fmt.Println(s2)
}
