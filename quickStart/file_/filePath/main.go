package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("dir1//", "filename")
	//dir/file on Linux vs. dir\file on Windows
	fmt.Println(p)
	fmt.Println("Dir(p):", filepath.Dir(p))   //文件目录
	fmt.Println("Base(p):", filepath.Base(p)) //文件名
	//检查string是否为绝对路径
	fmt.Println(filepath.IsAbs("D:/aliyun"))
	filename := "config.json"
	//获取扩展名，扩展（extension）
	ext := filepath.Ext(filename)
	fmt.Println(ext)
	//TrimSuffix函数用于从字符串的末尾去除指定的后缀（如果存在的话）
	fmt.Println(strings.TrimSuffix(filename, ext)) //获取文件名的名字（不包括扩展）
	//filepath.Rel返回一个表示从basepath到targetpath的相对路径的字符串，
	//以及一个可能出现的错误（如果路径无效或无法计算相对路径）。
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
