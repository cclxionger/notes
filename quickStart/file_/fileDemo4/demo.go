package main

import (
	"fmt"
	"io"
	"os"
)

// 拷贝文件
func main() {
	srcPath := "D:/222.jpg"
	desPath := "D:/333.jpg"
	//先通过路径打开文件
	srcFile, err := os.Open(srcPath)
	defer srcFile.Close()
	if err != nil {
		fmt.Println("open", err)
	}

	desFile, err := os.Create(desPath)
	if err != nil {
		fmt.Println("create", err)
	}
	defer desFile.Close()
	
	//io.Copy 这个方法也是边读边写
	io.Copy(desFile, srcFile)
	fmt.Println("copy")

}
