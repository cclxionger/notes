package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//与服务端建立连接
	//10.23.160.16是我的ip4
	conn, err := net.Dial("tcp", "10.23.160.16:8888")

	if err != nil {
		fmt.Println("dial failed")
		return
	}
	// defer conn.Close()
	fmt.Println("conn:", conn, "客户端ip为", conn.RemoteAddr())
	//os.Stdin表示标准输入 终端
	//从终端读取一行用户输入内容，发送给服务端
	reader := bufio.NewReader(os.Stdin)
	for {

		line, err := reader.ReadString('\n')
		for err != nil {
			fmt.Println("readString err=", err)
		}
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出")
			break
		}
		_, err = conn.Write([]byte(line + "\n"))

		if err != nil {
			fmt.Println("conn write err= ", err)
		}
		// fmt.Printf("客户端发送了%d个字节数据\n", n)
	}
	// for {
	// 	s, _ := input.ReadString('\n')
	// 	s = strings.TrimSpace(s)
	// 	if strings.ToUpper(s) == "Q" {
	// 		return
	// 	}
	// 	//给服务端发消息
	// 	_, err := conn.Write([]byte(s))
	// 	if err != nil {
	// 		fmt.Println("send failed")
	// 		return
	// 	}
	// 	//从服务端接收回复的消息
	// 	var buf [1024]byte
	// 	n, err := conn.Read(buf[:])
	// 	if err != nil {
	// 		fmt.Println("read failed")
	// 	}
	// 	fmt.Println("收到服务端回复", string(buf[:n]))

	// }

}

// func main() {
// 	resp, err := http.Get("https://gobyexample.com")
// 	if err != nil {
// 		panic(err)
// 	}
// 	//后面要关闭这个网页主体
// 	defer resp.Body.Close()
// 	//输出这个打开的这个网页的状态
// 	fmt.Println("Response status:", resp.Status)
// 	in := bufio.NewScanner(resp.Body)
// 	//Print the first 5 lines of the response body.
// 	for i := 0; in.Scan() && i < 5; i++ {
// 		content := in.Text()
// 		fmt.Println(content)
// 	}
// 	if err := in.Err(); err != nil {
// 		panic(err)
// 	}
// }
