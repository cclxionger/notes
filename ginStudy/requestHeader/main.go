package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//获取请求头
	r.GET("/", func(c *gin.Context) {
		//有Get字母大小写不区分  单词与单词之间用 - 连接，只返回第一个value
		//可以传重复的key
		fmt.Println(c.GetHeader("User-Agent"))
		fmt.Println(c.Request.Header.Get("user-Agent"))
		//没有Get要区分大小写，返回的是一个切片
		fmt.Println(c.Request.Header["user-Agent"])
	})
	//爬虫和用户区别对待
	r.GET("/p", func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")
		if strings.Contains(userAgent, "python") {
			c.JSON(0, gin.H{"data": ""})
			return
		}
		c.JSON(0, gin.H{"data": "这是相应用户的数据"})
	})
	//设置响应头
	r.GET("/h", func(c *gin.Context) {
		c.Header("Content-type", "application/text;charset=utf-8")
		c.JSON(0, gin.H{"data": "响应头"})
	})
	r.Run(":8080")
}
