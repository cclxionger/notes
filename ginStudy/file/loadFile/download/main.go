package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/download", func(c *gin.Context) {
		// 表示是文件流，唤起浏览器下载，一般设置了这个，就要设置文件名
		c.Header("Content-Type", "application/octet-stream")
		// 用来指定下载下来的文件名
		c.Header("Content-Disposition", "attachment; filename="+"牛逼.png")
		// 表示传输过程中的编码形式，乱码问题可能就是因为它
		c.Header("Content-Transfer-Encoding", "binary")
		//要下载文件的地址
		c.File("../upload/desFile/欠的钱.txt")
	})
	r.Run(":8080")
}
