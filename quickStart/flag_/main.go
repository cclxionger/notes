package main

import (
	"flag"
	"fmt"
)

func main() {
	//但是这种不太方便，例子：mysql 登录的时候 -u - r -p 顺序是任意的
	//所以要使用flag包进行解析
	var user string
	var pwd string
	var port string
	//&use 就是命令行-u后面要输入的参数
	//"u", 就是-u指定参数
	// ""默认值
	//最后一个参数的 进行说明
	flag.Parse()
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&port, "port", "3306", "密码，默认为3306")
	fmt.Println("用户名字是", user)
	fmt.Println("用户密码是", pwd)
	fmt.Println("用户端口是", port)
}
