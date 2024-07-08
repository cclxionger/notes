package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:age`
}
type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func UserList(c *gin.Context) {
	users := []UserInfo{
		{Name: "ck", Age: 19},
		{Name: "lc", Age: 20},
	}
	c.JSON(200, Response{0, users, "请求成功"})

}
func main() {
	router := gin.Default()
	r1 := router.Group("/api")
	//分组之后 调用的时候，需要/api/group1
	{
		r1.GET("/userInfo", UserList)
	}
	r2 := router.Group("/update")
	{
		r2.GET("/userInfo", UserList)
	}
	router.Run(":8080")
}
