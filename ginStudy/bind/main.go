package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name" form:"name" uri:"name" binding:"required,min=5" msg:"名字不能为空,最小长度为5"`
	Age  int    `json:"age" form:"age" uri:"age" binding:"gt=18" msg:"年龄需要大于18岁"`
	Date string `json:"date" binding:"datetime=2006-01-02"`
}

func main() {
	r := gin.Default()
	//ShouldBindJSON对应json
	r.POST("/json", func(c *gin.Context) {
		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{
				"msg": "你错了",
			})
			return
		}
		c.JSON(200, gin.H{
			"msg":  "ok",
			"date": user,
		})

	})
	///ShouldBindQuery对应form
	//浏览器输入form?name=ck&age=18
	r.POST("/form", func(c *gin.Context) {
		var user User
		err := c.ShouldBindQuery(&user)
		if err != nil {
			c.JSON(200, gin.H{
				"msg": "你错了",
			})
			return
		}
		c.JSON(200, gin.H{
			"msg": "ok",
		})

	})
	//c.ShouldBindUri对应uri
	// 浏览器输入/uri/fengfeng/21/男
	r.POST("/uri/:name/:age", func(c *gin.Context) {
		var user User
		err := c.ShouldBindUri(&user)
		if err != nil {
			c.JSON(200, gin.H{
				"msg": "你错了",
			})
			return
		}
		c.JSON(200, gin.H{
			"msg": "ok",
		})

	})
	//c.ShouldBind对应form-data、x-www-form-urlencode
	r.Run(":8080")
}
