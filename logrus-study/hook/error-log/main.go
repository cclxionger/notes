package errlog

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// 只把err文件输出到 err.log里面

/*
type Hook interface {
	Levels() []Level
	Fire(*Entry) error
}
MyHook 需要实现这俩个方法
*/
// 实现单或者多个level输出到文件
type MyHook struct {
	level  logrus.Level
	writer io.Writer
}

// NewFileHook 创建一个新的 MyHook 实例，用于将指定级别的日志写入到文件中
func NewMyHook(level logrus.Level, fileName string) (*MyHook, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &MyHook{level: level, writer: file}, nil
}

// 哪些日志等级生效
func (hook MyHook) Levels() []logrus.Level {
	return []logrus.Level{hook.level}
}

// Fire 是钩子被触发时调用的方法，负责处理日志条目
// 设置哪些加入到文件中或者输出
func (hook *MyHook) Fire(entry *logrus.Entry) error {
	// 这里我们假设你想要写入整个日志条目（包括时间戳、级别、消息和字段）
	// 因此我们需要先格式化这个条目为字符串
	// 注意：使用 entry.String() 可能不会包含所有字段，因为它依赖于默认的格式化器,如果需要加入字段，需要在它前面加入
	// 如果需要自定义格式，你需要自己构建字符串
	// 临时设置输出为 hook.writer，然后让 logrus 格式化并写入
	// 但是这样做可能会干扰全局的 logrus 输出设置，所以更推荐自己构建字符串
	entry.Data["study"] = "ck"
	entry.Data["learn"] = "lc"
	// 默认格式
	// line, err := entry.String()
	// if err != nil {
	// 	return err
	// }

	// 自定义字符串格式
	// 我想把 entry.Data 的数据放到打印每一行的最前面
	var Dataline string
	var line string
	for k, v := range entry.Data {
		Dataline += k + ": " + fmt.Sprintf("%+v", v) + "\t" // 注意：%+v 用于显示结构体和数组的字段名
	}
	line += Dataline
	line += entry.Time.Format("2006-01-02 15:04:05") + "\t[" + entry.Level.String() + "]\t" + entry.Message + "\n"
	// 将日志条目写入文件
	_, err := hook.writer.Write([]byte(line))
	return err
}
func main() {
	// 设置输出到控制台的样式,如果想要文件和控制台输出一样，需要重新设置控制台
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true},
	)
	// 设置输出等级
	logrus.SetLevel(logrus.DebugLevel)
	
	// 增加debugHook
	debugLog, _ := NewMyHook(logrus.DebugLevel, "debug.log")
	logrus.AddHook(debugLog)
	data := "ck is silly"
	logrus.Debugf("debug %s balabalabadladmakdnakd", data)
	logrus.Debugln("ck debug")
	// 增加errHook
	errLog, _ := NewMyHook(logrus.ErrorLevel, "err.log")
	logrus.AddHook(errLog)
	logrus.Errorln("ck error")
}
