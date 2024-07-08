package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//解析模板
	//根据文件目录解析
	// r.LoadHTMLFiles("allHtml/index.html")
	//根据regex解析
	r.LoadHTMLGlob("allHtml/*")
	//渲染模板
	r.GET("/index", func(c *gin.Context) {
		//gin.H相当于是map
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "<a href='https://flowus.cn/xionger/share/3e8ec5f2-dc29-410c-a9b6-813b2bc1fc7d?code=EZ5632'>这是一个链接到我的笔记的文本</a>",
		})
	})
	r.GET("/index2", func(c *gin.Context) {
		//gin.H相当于是map
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "index2",
		})
	})
	r.Run(":9090")
}
