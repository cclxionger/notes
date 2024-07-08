package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 进行权限验证的中间件
func tokenMiddleWare(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "123456" {
		c.Next()
		return
	}
	c.JSON(200, "权限验证失败")
	c.Abort()
}

// 进行耗时统计的中间件
func timeMiddleWare(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	since := time.Since(startTime)
	sinceSecond := since.Seconds()
	// 获取当前请求所对应的函数
	f := c.HandlerName()
	fmt.Printf("函数 %s 耗时 %.2fs\n", f, sinceSecond)
}
func main() {
	router := gin.Default()
	//进行权限验证和耗时统计的分组
	r := router.Group("/validation").Use(tokenMiddleWare, timeMiddleWare)
	{
		r.POST("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "验证通过",
			})
		})
	}
	router.Run(":8080")

}
