package main

import "github.com/gin-gonic/gin"

type loginForm struct {
	User     string `form:"user" binding: "required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/login", func(ctx *gin.Context) {
		var form loginForm
		if ctx.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				ctx.JSON(200, gin.H{"status": "you can logged in"})
			} else {
				ctx.JSON(401,gin.H{"status": "unauthorized"})
			}
		}
	})
	r.Run(":8080")
}
