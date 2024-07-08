package main

import "github.com/sirupsen/logrus"

func main() {
	// time="2024-05-25T08:59:19+08:00" level=info msg="我是Println"
	logrus.Println("我是Println")
	// 4
	println(logrus.GetLevel())
	// Info等级就是4
	logrus.Info("我是Info")
	// 不会打印出来，默认等级到4就不打印了，DebugLevel和TraceLevel
	logrus.Debugln("我是Debugln")
	//如果想要打印，需要设置等级
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debugln("我是Debugln")
	// 设置特定字段
	// 设置单个字段
	log1 := logrus.WithField("project", "study")
	log1.Errorln("我有单个字段")
	// 设置多个字段
	log2 := logrus.WithFields(logrus.Fields{
		"func": "main",
		"name": "ck",
	})
	log2.Println("我有多个子段")
	// 显示样式 Text和Json，默认是Test样式，可以设置为Json样式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Println("我是Json样式")

}
