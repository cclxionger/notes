package format

import (
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

type MyFormat struct {
}

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (format *MyFormat) Format(entry *logrus.Entry) ([]byte, error) {
	b := &bytes.Buffer{}
	// 自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	// 自定义字段格式
	var data string
	for k, v := range entry.Data {
		data += fmt.Sprintf("%s : %s\t", k, v)
	}
	// fmt.Sprintf：格式化字符串并返回结果字符串，不写入任何 I/O 流。
	// fmt.Fprintf：将格式化的数据写入到指定的 I/O 流中（如文件、网络连接或内存缓冲区
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		// 自定义输出格式
		fmt.Fprintf(b, "%s[%s]\t%s %s\t[%s] %s", data, timestamp, funcVal, fileVal, entry.Level, entry.Message)
	} else {
		fmt.Fprintf(b, "%s[%s]\t[%s] %s", data, timestamp, entry.Level, entry.Message)
	}
	b.WriteByte('\n')
	return b.Bytes(), nil
}

var log *logrus.Logger

// func init() {
// 	log = NewLog()
// }
func NewLog() *logrus.Logger {
	// 新建一个实例
	myLog := logrus.New()
	// 打印到哪里
	myLog.SetOutput(os.Stdout)
	// 开启返回函数名和行号
	myLog.SetReportCaller(true)
	// 设置格式
	myLog.SetFormatter(&MyFormat{})

	return myLog
}
func main() {
	log.WithFields(logrus.Fields{
		"author":  "lcc",
		"brother": "lc",
	}).Info("infott")
	// 这样打印不出来有字段的Info
	log.WithField("sb", "666")
	log.Infoln("info log")
}
