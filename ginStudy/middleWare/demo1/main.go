package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
	Age  int
}

func ware1(c *gin.Context) {
	c.Set("user", User{"小晨", 19})
	log.Println("我是中间件1 input")
	c.Next() //中间件放行，如果没有这个，会1 input, 1 out
	log.Println("我是中间件1 out")

}
func ware2(c *gin.Context) {
	log.Println("我是中间件2 input")
	// c.Abort() //不执行后来的 中间件3
	c.Next()
	log.Println("我是中间件2 out")
}
func ware3(c *gin.Context) {
	log.Println("我是中间件3 input")
	c.Next()
	log.Println("我是中间件3 out")
}
func global(c *gin.Context) {
	log.Println("我是全局中间件")
}
func main() {
	r := gin.Default()
	r.Use(global) //这个是第一个输出,作为全局
	r.GET("/", ware1, ware2, ware3, func(c *gin.Context) {
		userType, ok := c.Get("user")
		if ok != false {
			user := userType.(User)
			log.Println(user.Name, "年龄是", user.Age)
			log.Println()
			
		}
	})
	r.Run(":8080")
}
