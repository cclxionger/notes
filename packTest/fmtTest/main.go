package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//Errorf
	err := fmt.Errorf("user ccl  not found") //like define Error personally again
	fmt.Println(err.Error())
	fmt.Println()
	//Fprint  Fprinf  Fprinln  		similar
	n, err := fmt.Fprint(os.Stdout, "a") //os.Stdout meanings Standard Output
	// The n and err return values from Fprint are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
	fmt.Print("\n", n, " bytes written.\n")
	fmt.Println()

	//Fscanf
	/*
		fmt.Fscanf 的设计初衷是按照提供的格式字符串一次读取一个或多个值，
		并将它们分别赋给对应的变量
	*/
	var (
		i int
		b bool
		s string
	)
	r := strings.NewReader("5 true gophers false,6")
	n, err = fmt.Fscanf(r, "%d %t %s", &i, &b, &s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	}
	fmt.Println(i, b, s)
	fmt.Println(n) //i\b\s numbers  （max = 3）
	fmt.Println()
	//Fscanln  读取每一行的内容
	mq := `dmr 1771 1.61803398875
	ken 271828 3.14159` //换行并且没用使用\n 所以使用反引号``
	r = strings.NewReader(mq)
	var a string
	var bd int
	var c float64
	for {
		n, err := fmt.Fscanln(r, &a, &bd, &c)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%d: %s, %d, %f\n", n, a, bd, c)
	}
	fmt.Println()

	/*
			fmt.Scan
		输入多个时候，用空格或者换行分割
			fmt.Scanf
		输入多个时候，只能用空格分隔
			fmt.Scanln
		输入多个时候，只能用换行分割（并且一行里面不能有空格）
	*/

	/*
			fmt.Sprint
		按照fmt.Sprint格式重新制定输出格式，相当于实现了java中的toString方法
		fmt.Sprinf和fmt.Sprinln同理
	*/
	//	fmt.Sscanf/fmt.Sscan/fmt.Sscanln
	//使用 fmt.Sscanf 可以方便地从字符串中提取信息，并将其转换为程序能够处理的变量类型。
	//这在处理从文件、网络或其他数据源读取的文本数据时特别有用。
	//n, err = fmt.Sscanf("Kim is 22 years old", "%s is %d years old", &name, &age)

	//		实现GoString（）方法
	/*如果结构体里面包含另外一个结构体，想要正确输出格式，需要实现GoString()方法
	并且在GoString（）方法中使用 fmt.Sprint /fmt.Sprinf格式返回
	还需要在主方法中使用Printf进行输出， 格式化占位符是 %#v，如果是%v或者%+v的话，第二个结构体会返回地址
	*/
	// 		String（）方法类似
	/*
		用途：Stringer主要用于自定义类型的格式化输出，以提供用户友好的字符串表示；
		而GoString()主要用于获取值的Go语法表示，以辅助调试和测试。
		使用场景：Stringer通常用于控制类型在格式化输出时的表现；
		而GoString()则更多用于内部调试和反射操作。
	*/

}
