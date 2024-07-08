package log

import (
	cformat "notes/logrus-study/gin-logrus/format"
	chook "notes/logrus-study/gin-logrus/hook"

	"github.com/sirupsen/logrus"
)

func NewLog() *logrus.Entry {

	log := cformat.NewLogrus("")
	l := log.WithFields(logrus.Fields{
		"author":  "cl",
	})
	errLog, _ := chook.NewMyHook(logrus.ErrorLevel, "./log/errLog.log")
	l.Logger.AddHook(errLog)

	infoLog, _ := chook.NewMyHook(logrus.InfoLevel, "./log/infoLog.log")
	l.Logger.AddHook(infoLog)

	l.Logger.SetLevel(logrus.DebugLevel)
	debugLog, _ := chook.NewMyHook(logrus.DebugLevel, "./log/debugLog.log")
	l.Logger.AddHook(debugLog)

	warnLog, _ := chook.NewMyHook(logrus.WarnLevel, "./log/warnLog.log")
	l.Logger.AddHook(warnLog)

	return l
}
