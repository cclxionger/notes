package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func examineErr(err error) {
	if err != nil {
		fmt.Println("出错了", err)

	}
}
func main() {

	//第一种方式读取文件
	//一次性读取文件，适合文件较小的情况下
	content, err := os.ReadFile("d:/text2.txt")

	contStr := string(content)
	fmt.Printf("%v\n", contStr)
	//第二种方式读取文件先打开
	//打开文件
	file, err := os.Open("D:/text2.txt")
	examineErr(err)
	defer file.Close()
	fmt.Println(file) //file是个指针
	//利用bufio.NewReader(file)读取文件
	reader := bufio.NewReader(file)
	/*
		Peek 方法是 bufio.Reader 的一个方法，
		它允许你查看缓冲区中的下一个 n 个字节，
		而不实际从输入中消耗（读取并移动偏移量）这些字节。
		这在你需要基于接下来的几个字节做出决定但又不想改变读取位置时非常有用。
	*/
	b4, err := reader.Peek(10)
	examineErr(err)
	fmt.Println()
	fmt.Printf("10 bytes: %s\n", string(b4))
	for {
		str, err := reader.ReadString('\n') //每次读取每行
		fmt.Print(str)                      //每次输出str，自动换行
		if err == io.EOF {
			break //读完了
		}

	}
	fmt.Println()
	//第三种方法
	// 在使用 file.Read 之前重置文件的位置
	_, err = file.Seek(0, io.SeekStart)
	//如果只读取五个字节
	b1 := make([]byte, 5)
	n1, err := file.Read(b1)
	examineErr(err)
	fmt.Printf("%dbytes : %s\n", n1, string(b1[:n1]))

	//每次读取五个字节  但是这样如果有汉字会出现乱码
	for {
		n1, err = file.Read(b1)
		if err != nil {
			if err == io.EOF {
				break
			}
			examineErr(err)

		}
		fmt.Printf("%dbytes : %s\n", n1, string(b1[:n1]))
	}
	/*
		func (f *File) Seek(offset int64, whence int) (ret int64, err error)

		offset：相对于whence的偏移量，以字节为单位。
		whence：指定偏移量的起始位置，它可以是以下三个常量之一：
		io.SeekStart：从文件开始处计算偏移量（0为文件的开始）。
		io.SeekCurrent：从当前位置计算偏移量。
		io.SeekEnd：从文件末尾计算偏移量（负数表示从文件末尾往回数）
	*/
	// _, err = file.Seek(0, io.SeekStart)
	_, err = file.Seek(-4, io.SeekEnd)
	examineErr(err)
	n1, err = file.Read(b1)
	examineErr(err)
	fmt.Println()
	fmt.Printf("%dbytes : %s\n", n1, string(b1[:n1]))
	_, err = file.Seek(0, io.SeekStart)
	//n3是读取的返回的最少字节，如果读取少于3个字节，会返回错误（不是EOF）
	//第三个参数一定要比b1的长度小，否则一定会返回错误
	n3, err := io.ReadAtLeast(file, b1, 3)
	fmt.Println()
	examineErr(err)
	fmt.Printf("%d bytes : %s\n", n3, string(b1))
}
