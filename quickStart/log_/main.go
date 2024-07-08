package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("666")
	/*
		log.Ldate：包含日期（默认为 2006/01/02 的格式）
		log.Ltime：包含时间（默认为 15:04:05 的格式）
		log.Lmicroseconds：包含微秒级的时间（例如 15:04:05.000000）
		log.Llongfile：包含完整的文件名和行号（例如 /a/b/c/d.go:23）
		log.Lshortfile：仅包含文件名和行号（例如 d.go:23）
		log.LstdFlags：初始的默认标志（Ldate | Ltime）
	*/
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Ltime)
	log.Println("666")
	log.SetFlags(log.Llongfile)
	log.Println("666")
	//log.New返回的是*log.Logger
	mylog := log.New(os.Stdout, "tied:", log.LstdFlags) //在输出的时候增加一个前缀tied:
	mylog.Println("666")
	//把前缀改为ohmy:
	mylog.SetPrefix("ohmy:")
	mylog.Println("666")
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)
	buflog.Println("hello打都锕dada")
	//这条消息并没有直接输出到控制台或文件中。
	//它只是被写入到了 buf 缓冲区。如果你想查看这条消息，你需要从 buf 缓冲区中读取它
	fmt.Print(buf.String())
	//这三行目前看不懂
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

}
