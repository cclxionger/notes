package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ArticleModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// 得到文章列表
func _getArticles(c *gin.Context) {
	articleList := []ArticleModel{
		{"Go语言入门", "这篇文章是《Go语言入门》"},
	}
	response := Response{
		0, articleList, "ok",
	}
	c.JSON(200, response)
}

// 得到文章详情
func _getDetail(c *gin.Context) {
	//获取param中的id
	fmt.Println(c.Param("id"))
	articleList := []ArticleModel{
		{"Go语言入门", "这篇文章是《Go语言入门》"},
		{"python语言入门", "这篇文章是《python语言入门》"},
		{"JavaScript语言入门", "这篇文章是《JavaScript语言入门》"},
	}
	response := Response{
		0, articleList, "ok",
	}
	c.JSON(200, response)
}

// 增加文章
func _post(c *gin.Context) {
	var articleList ArticleModel
	body, _ := c.GetRawData()
	contentTpye := c.GetHeader("Content-Type")
	if contentTpye == "application/json" {
		err := json.Unmarshal(body, &articleList)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	c.JSON(200, Response{0, articleList, "添加成功"})
}

//http://127.0.0.1:8080/static/dae.png

// 修改文章
func _put(c *gin.Context) {

}

// 删除文章
func _delete(c *gin.Context) {
	fmt.Println(c.Param("id"))
	c.JSON(200, Response{200, map[string]string{}, "删除成功"})
}
func main() {
	r := gin.Default()
	r.GET("/articles", _getArticles)
	r.GET("/articles/:id", _getDetail)
	r.POST("/articles", _post)
	r.DELETE("articles/:id", _delete)
	r.Run(":8080")
}
