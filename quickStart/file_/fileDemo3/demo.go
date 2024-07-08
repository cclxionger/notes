package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// //文件小的时候，一次性读，一次性写
	// fileReadPath := "d:/output.txt"
	// fileWritePath := "d:/text2.txt"
	// data, err := os.ReadFile(fileReadPath)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = os.WriteFile(fileWritePath, data, 0666)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//当文件特别大的时候，边读边写

	fileReadPath := "d:/virus.txt"
	fileWritePath := "d:/virus25.txt"
	sourceFile, err := os.Open(fileReadPath)
	if err != nil {
		fmt.Println(err)
	}
	defer sourceFile.Close()
	aimFile, err := os.OpenFile(fileWritePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer aimFile.Close()
	// 在Go语言中，io.Reader 是一个接口，
	// 而 *os.File 是 os.File 类型的指针，
	// 并且 os.File 类型实现了 io.Reader 接口。
	// 因此，*os.File 可以被当作 io.Reader 来使用，
	// 因为它们之间存在接口和实现的关系。
	//进行流式读取和写入
	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(aimFile)
	//定义一个缓冲区，4096个字节
	buffer := make([]byte, 4096)
	//边读边写
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		writer.Write(buffer[:n])
	}
	writer.Flush()
	fmt.Println("文件复制完成")

}
