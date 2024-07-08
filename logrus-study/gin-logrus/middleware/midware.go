package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 耗时中间件
func TimeConMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 进行下一个中间件
		c.Next()
		// 返回来这个中间件的耗时统计
		since := time.Since(startTime)
		// 单位为秒
		sinceSecond := since.Seconds()
		// 获取当前请求所对应的函数
		f := c.HandlerName()
		// 打印
		fmt.Printf("函数 %s 耗时 %.2fs\n", f, sinceSecond)
	}

}
