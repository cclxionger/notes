package cformat

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

type Myformat struct {
}

// func (f *TextFormatter) Format(entry *Entry) ([]byte, error)
func (f Myformat) Format(entry *logrus.Entry) ([]byte, error) {
	b := &bytes.Buffer{}
	// 自定义时间格式
	timeFormat := entry.Time.Format("2006-01-02 15:04:05")
	// 自定义字段格式
	var data string
	for k, v := range entry.Data {
		data += fmt.Sprintf("%s : %s\t", k, v)
	}
	// 如果开启了行号和函数
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(b, "%s[%s]\t %s %s\t[%s]msg:%s\n", data, timeFormat, funcVal, fileVal, entry.Level, entry.Message)
		// 如果没有开启行号和函数
	} else {
		fmt.Fprintf(b, "%s[%s]\t[%s]%s\n", data, timeFormat, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

var log *logrus.Logger

func NewLogrus(filePath string) *logrus.Logger {
	// 新建一个实例
	myLog := logrus.New()
	// 设置格式
	myLog.SetFormatter(&Myformat{})
	// 开启返回函数名和行号
	myLog.SetReportCaller(true)
	// 打印到哪里,如果传来的是空字符串则只打印到控制台
	if filePath == "" {
		myLog.SetOutput(os.Stdout)
	} else {
		file, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		myLog.SetOutput(io.MultiWriter(file, os.Stdout))
	}
	return myLog
}
