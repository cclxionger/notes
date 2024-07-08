package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//第一种写入文件方法  对应读的第四种方法
	//使用os.OpenFile打开文件，然后使用bufio.NewWriter.WriteString写入
	filePath := "D:/output.txt"
	/*
		const (
		// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
		O_RDONLY int = syscall.O_RDONLY // open the file read-only.
		O_WRONLY int = syscall.O_WRONLY // open the file write-only.
		O_RDWR   int = syscall.O_RDWR   // open the file read-write.
		// The remaining values may be or'ed in to control behavior.
		O_APPEND int = syscall.O_APPEND // append data to the file when writing.
		O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
		O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
		O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
		O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
	)*/
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666) //FileMode在windows系统下不起作用
	//不能这样打开文件，因为    这样不让写
	// file, err := os.Open("D:/output.txt")
	if err != nil {

		fmt.Println(err)
	}
	defer file.Close()
	str := "hello world哈哈哈\n"
	writer := bufio.NewWriter(file)
	writer.WriteString(str)
	//使用缓存写之后必须得writer.Flush()
	writer.Flush()

	//第二种方法
	//使用os.WriteFile方法，这个方法里面调用了OpenFile
	// 注意这个方法不能没有Append ，不能追加
	d1 := []byte(str)
	err = os.WriteFile("D:/output2.txt", d1, 000)
	examineErr(err)
	//第三种方法
	//使用os.Create创建这个文件夹，然后使用os.Create.Write进行写入
	file, err = os.Create("D:/output3.txt")
	examineErr(err)
	n3, err := file.Write(d1)
	examineErr(err)
	fmt.Println("写入了字节数量为", n3)
	//第四种方法
	//调用*File.WriteString方法
	//这个*File哪里来的呢，可以从os.Create返回值来，也可以从os.OpenFile返回值来
	//注意WriteString是追加的，前面应该是没有os.O_APPEND
	//换言之，想要追加，只能用第一个方法（增添os.O_APPEND）和第四个方法
	n4, err := file.WriteString("6666")
	examineErr(err)
	fmt.Println("写入了字节数量为", n4)
}

func examineErr(err error) {
	if err != nil {
		fmt.Println("出错了", err)

	}
}
