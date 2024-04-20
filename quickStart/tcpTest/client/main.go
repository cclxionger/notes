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
	conn, err := net.Dial("tcp", "10.23.39.142")
	if err != nil {
		fmt.Println("dial failed")
	}
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}
		//给服务端发消息
		_, err := conn.Write([]byte(s))
		if err != nil {
			fmt.Println("send failed")
			return
		}
		//从服务端接收回复的消息
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("read failed")
		}
		fmt.Println("收到服务端回复",string(buf[:n]))

	}

}
