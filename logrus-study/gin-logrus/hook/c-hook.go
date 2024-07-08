package chook // 你的包名

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

// MyHook 是一个logrus钩子，它将指定级别的日志记录到文件中
type MyHook struct {
	Writer   io.Writer
	Filename string
}

var MyLever logrus.Level

// NewMyHook 创建一个新的MyHook实例
func NewMyHook(level logrus.Level, filename string) (*MyHook, error) {
	MyLever = level
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &MyHook{
		Writer:   file,
		Filename: filename,
	}, nil
}

// Fire 是当钩子被触发时调用的方法
func (hook *MyHook) Fire(entry *logrus.Entry) error {
	b := &bytes.Buffer{}
	// 自定义时间格式
	timeFormat := entry.Time.Format("2006-01-02 15:04:05")
	// 自定义字段格式
	var data string
	for k, v := range entry.Data {
		data += fmt.Sprintf("%s=%v\t", k, v)
	}
	// 如果开启了行号和函数
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(b, "%s[%s]\t %s %s\t[%s] msg:%s\n", data, timeFormat, funcVal, fileVal, entry.Level, entry.Message)
	} else {
		fmt.Fprintf(b, "%s[%s]\t[%s] msg:%s\n", data, timeFormat, entry.Level, entry.Message)
	}

	// 写入日志到文件
	_, err := b.WriteTo(hook.Writer)
	return err
}

// Close 关闭钩子使用的文件句柄
func (hook *MyHook) Close() error {
	return hook.Writer.(*os.File).Close()
}

// Levels 定义了哪些日志级别应该被处理
func (hook *MyHook) Levels() []logrus.Level {
	// 这里返回你希望处理的日志级别，例如只处理Error级别
	return []logrus.Level{MyLever}
}
