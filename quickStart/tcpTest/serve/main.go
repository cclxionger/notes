package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	//开启服务
	listen, err := net.Listen("tcp", "10.23.39.142") //第一个字符串传的是哪个类型的网络上监听连接
	//比如："tcp"、"tcp4"、"tcp6"：表示 IPv6 的 TCP 网络。。"ip"：表示原始 IP 数据包。
	if err != nil {
		fmt.Println("listen failed")
		return
	}
	for {
		Conn, err := listen.Accept() //建立连接
		if err != nil {
			fmt.Println("accept failed")
			continue
		}
		go process(Conn)
	}

}
func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read failed")
			break
		}
		recv := string(buf[:n])
		fmt.Println("接收到数据为")
		fmt.Println(recv)
		conn.Write(buf[:n])

	}
}
