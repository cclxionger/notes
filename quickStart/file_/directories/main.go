package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//创建文件夹（目录）
	err := os.Mkdir("SUB", 0755)
	examineErr(err)
	/*
		os.Remove：用于删除指定的文件或（空的）目录。
		如果尝试删除一个非空的目录，它会返回一个错误。
		os.RemoveAll：用于删除指定的文件或目录，包括其所有子目录和文件。
		它会递归地删除目录及其内容。
	*/
	defer os.RemoveAll("SUB")
	//读取指定目录的文件夹或者文件
	dir := ".." // 要读取的目录
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("无法读取目录 %s: %v", dir, err)
	}
	for _, entry := range entries {
		fmt.Printf("%s: %t\n", entry.Name(), entry.IsDir())
	}
	//filepath.WalkDir可以visit is called for every file or directory
	// io.fs.DirEntry  是一个接口，里面有Name() string，IsDir() bool
	err = filepath.WalkDir("..", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(" ", path, d.IsDir())
		return nil
	})
}
func examineErr(err error) {
	if err != nil {
		fmt.Println("出错了", err)

	}
}
