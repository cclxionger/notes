package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/form",func(ctx *gin.Context) {
		message := ctx.PostForm("message")
		// messageAnonymous := ctx.DefaultPostForm("message","messageAnonymous ")
		nick := ctx.DefaultPostForm("nick","anonymous")
		ctx.JSON(200,gin.H{
			"status" : "posted",
			"message":message,
			"nick":nick,
		})
	})
	r.Run(":8080")
}