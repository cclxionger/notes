package main

import (
	"notes/logrus-study/gin-logrus/log"
	"notes/logrus-study/gin-logrus/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(middleware.TimeConMiddleware())
	l := log.NewLog()
	router.GET("/", func(c *gin.Context) {

		l.Info("info来了")
		c.JSON(200, gin.H{"msg": "终于运行了"})
	})
	router.Run(":8080")

}
