package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	log1 := logrus.WithField("project", "study")
	log1.Debugln("Debugln")
	log1.Infoln("Infoln")
	log1.Warnln("Warnln")
	log1.Errorln("Errorln")
	log1.Println("Println")
	log2 := logrus.WithFields(logrus.Fields{
		"ck": 19,
		"lc": 20,
	},
	)
	log2.Println("多个子段打印")
	//显示样式设置为json类型
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus2 := logrus.WithField("project", "study")
	logrus2.Println("666")
	file, _ := os.OpenFile("info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	//  同时写文件和屏幕
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))
	logrus.Println("我真的没有颜色啊 ")

}
