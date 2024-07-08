package main // 声明这个Go源文件属于main包，这意味着它包含了程序的入口点。

import (
	"fmt" // 导入fmt包，它包含了格式化输入输出的函数。
)

// 定义了一个名为fibonacci的函数，它接收两个通道参数：c用于发送斐波那契数列的数字，quit用于接收停止信号。
func fibonacci(c, quit chan int) {
	x, y := 0, 1 // 初始化斐波那契数列的前两个数字。

	for { // 无限循环，直到收到quit通道的信号。
		select { // 使用select语句来等待多个通信操作中的一个发生。
		case c <- x: // 如果通道c可以接收数据，则发送当前斐波那契数x。
			x, y = y, x+y // 更新x和y的值以计算下一个斐波那契数。

		case <-quit: // 如果quit通道有数据可读，则接收该数据。
			fmt.Println("quit") // 打印"quit"表示接收到停止信号。
			return              // 退出函数。
		}
	}
}

func main() { // 定义main函数，它是程序的入口点。
	c := make(chan int, 5)    // 创建一个整数类型的通道c，用于发送斐波那契数列的数字。
	quit := make(chan int, 5) // 创建一个整数类型的通道quit，用于发送停止信号。

	go func() { // 使用go关键字启动一个新的goroutine来并发执行匿名函数。
		for i := 0; i < 5; i++ { // 循环5次，每次从通道c接收一个斐波那契数并打印。
			fmt.Println(<-c) // 从通道c接收数据并打印。
		}
		quit <- 6 // 当接收到5个斐波那契数后，向quit通道发送一个值（这里发送6，但实际上值不重要，重要的是发送操作本身）。
	}() // 匿名函数定义结束，goroutine开始执行。

	fibonacci(c, quit) // 调用fibonacci函数，并传入之前创建的通道c和quit。
	
} // main函数结束，程序执行完毕。
