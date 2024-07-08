package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	os.Setenv("Foo", "1")
	log.Println("foo:", os.Getenv("Foo"))
	os.Setenv("hello", "2")
	//os.Environ()返回所有环境变量的一个切片
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		log.Println(pair)
	}

}
