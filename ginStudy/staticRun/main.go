package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.StaticFile("/static/dae.png", "ginStudy/static/dae.png")
	r.Run(":8080")
}
