package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	//os.Args返回的是 一个切片
	// //切片第一个是文件的名字 ，其余是你输入的（以空格分割）
	fmt.Println(argsWithProg)
	//但是这种不太方便，例子：mysql 登录的时候 -u - r -p 顺序是任意的
	//所以要使用flag包进行解析
	var user string
	var pwd string
	var port string
	//&use 就是命令行-u后面要输入的参数
	//"u", 就是-u指定参数
	// ""默认值
	//最后一个参数的 进行说明

	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&port, "port", "3306", "密码，默认为3306")
	flag.Parse()
	//应该在调用flag.Parse()之后再访问标志变量的值，
	//因为flag.Parse()负责解析命令行参数并更新标志变量的值
	fmt.Println("用户名字是", user)
	fmt.Println("用户密码是", pwd)
	fmt.Println("用户端口是", port)
	//lag.Args()函数用于获取命令行参数中未被解析为标志（flags）的部分
	fmt.Println("tail:", flag.Args())
	numbPtr := flag.Int("numb", 42, "an int")
	fmt.Println("numb:", *numbPtr)

}
