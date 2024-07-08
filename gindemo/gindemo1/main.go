package main

import (
	"github.com/gin-gonic/gin"
)

func sayhello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello golang",
	})
}
func main() {
	// 返回默认的路由引擎
	r := gin.Default()
	r.GET("/hello", sayhello)
	r.Run(":8080")
}
