package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/loadFile", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)
		fmt.Println(file.Size / 1024) //文件的字节
		//读取文件内容
		// fileRead, _ := file.Open()
		// data, _ := io.ReadAll(fileRead)
		// fmt.Println(string(data))
		des := "./" + file.Filename
		//里面用到了缓冲流，可以下载很大的文件
		c.SaveUploadedFile(file, des)
		c.JSON(200, gin.H{
			"msg": "ok",
		})
	})
	r.POST("/loadFiles", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["files"]
		fmt.Println(files)
		for _, file := range files {
			fmt.Println(file.Filename)
			des := "./" + "desFile/" + file.Filename
			c.SaveUploadedFile(file, des)
			c.String(200, fmt.Sprintf("%v下载完成", file.Filename))
		}

	})
	r.Run(":8080")

}
