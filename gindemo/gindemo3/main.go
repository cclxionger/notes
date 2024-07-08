package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Delims("{[{", "}]}")
	//gin框架中给模板添加自定义函数
	//模式化年月日
	r.SetFuncMap(template.FuncMap{
		"title": func(t time.Time) string {
			year, month, day := t.Date()
			return fmt.Sprintf("%d年%02d月%02d日", year, month, day)
		},
	})
	//解析模板
	r.LoadHTMLFiles("./date.html")
	//渲染模板
	r.GET("/date", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "date.html", gin.H{
			"now": time.Date(2024, 05, 14, 0, 0, 0, 0, time.UTC),
		})
	})
	r.Run(":8080")
}
