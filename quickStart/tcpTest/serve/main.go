package main

import (
	"fmt"
	"net"
)

func main() {
	//开启服务
	//tcp表示使用网络协议是tcp
	//0.0.0.0：8888表示在本地监听8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888") //第一个字符串传的是哪个类型的网络上监听连接
	//比如："tcp"、"tcp4"、"tcp6"：表示 IPv6 的 TCP 网络。。"ip"：表示原始 IP 数据包。
	if err != nil {
		fmt.Println("listen failed", err)
		return
	}
	defer listen.Close()
	for {
		//listen.Accept()会一直等待客户端建立链接
		conn, err := listen.Accept() //建立连接
		if err != nil {
			fmt.Println("accept failed")

		} else {
			fmt.Println("Coon=", conn)
			//准备一个协程，为客户端服务
			go process(conn)
		}

	}
}
func process(conn net.Conn) {
	defer conn.Close()
	//循环读取接收客户端发送的数据
	for {
		buf := make([]byte, 1024)
		fmt.Println("服务端等待客户端发送消息,客户端为", conn.RemoteAddr())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务端readerr", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}

}

// }
// func process(conn net.Conn) {
// 	defer conn.Close()
// 	for {
// 		reader := bufio.NewReader(conn)
// 		var buf [128]byte
// 		n, err := reader.Read(buf[:])
// 		if err != nil {
// 			fmt.Println("read failed")
// 			break
// 		}
// 		recv := string(buf[:n])
// 		fmt.Println("接收到数据为")
// 		fmt.Println(recv)
// 		conn.Write(buf[:n])

// 	}
// }
