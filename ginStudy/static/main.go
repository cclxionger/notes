package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func _redirect(c *gin.Context) {
	c.Redirect(301, "https://www.baidu.com")
}
func main() {
	r := gin.Default()
	//在golang中，没有相对文件的路径，只有相对项目的路径
	//StaticFS第二个参数是文件的目录，访问的时候是:8080/staticTest/dae.png
	r.StaticFS("/staticTest", http.Dir("staticTest"))
	// 访问的时候是:8080/jitui.png
	r.StaticFile("/jitui.png", "jitui.png")
	//重定向
	r.GET("/baidu", _redirect)
	r.Run(":8080")
}
