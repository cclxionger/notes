package main

import (
	"fmt"
	"io"
	"notes/logrus-study/color"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {

	data := "我有红颜色啦"
	fmt.Println(color.Forecolor("红色", data))
	// logrus也是支持颜色输出的
	// 打印 INFO[0000]  我有红颜色啦，看起来似乎有点怪，那是因为没有设置输出格式
	logrus.Println(color.Forecolor("红色", data))
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true},
	)
	// 打印 INFO[2024-05-25 09:36:40]  我有红颜色啦
	logrus.Println(color.Forecolor("红色", data))
	// 当然还有其他一些设置
	/*
		ForceColors：是否强制使用颜色输出。
		DisableColors：是否禁用颜色输出。
		ForceQuote：是否强制引用所有值。
		DisableQuote：是否禁用引用所有值。
		DisableTimestamp：是否禁用时间戳记录。
		FullTimestamp：是否在连接到 TTY 时输出完整的时间戳。
		TimestampFormat：用于输出完整时间戳的时间戳格式。
	*/

	// 正片开始，日志输出到文件里面
	file, _ := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// logrus.SetOutpu 只能设置一个 io.Writer 参数
	logrus.SetOutput(file)
	// 输出到文件和屏幕上 , io.MultiWriter可以传多个 io.Writer 参数
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))
	// 设置显示行号
	logrus.SetReportCaller(true)
	// INFO[2024-05-25 09:46:28]  我是绿色，我还写入到文件里面了
	logrus.Println(color.Forecolor("绿色", "我是绿色，我还写入到文件里面了"))

}
