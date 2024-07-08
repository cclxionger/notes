package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 查询参数
func _query(c *gin.Context) {
	//网页输入http://127.0.0.1/query?id=2&user=ck&user=cl
	user := c.Query("user")
	//不知道为什么c.Get 一直返回nil，false
	fmt.Println(c.Get("user"))
	fmt.Println(c.QueryArray("user")) //返回[ck,cl]
	fmt.Println(user)                 //返回ck

}

// 动态参数
func _param(c *gin.Context) {
	//网页输入http://127.0.0.1/param/da
	fmt.Println(c.Param("user_id")) //得到da
	//网页输入http://127.0.0.1/param/da/bc
	fmt.Println(c.Param("book_id")) //得到bc
}

// 表单参数
func _form(c *gin.Context) {
	//网页输入http://127.0.0.1/form
	//form-data传入 user,ab user,cd
	fmt.Println(c.PostForm("user"))      //得到ab
	fmt.Println(c.PostFormArray("user")) //得到[ab,cd]
	//form-data没有传入 age
	//默认age为18
	fmt.Println(c.DefaultPostForm("age", "18"))
	fmt.Println(c.PostForm("age")) //得到18
	fmt.Println(c.MultipartForm()) //得到结构体&{map[user:[ab cd]] map[]} <nil>
}

// 原始参数
func _raw(c *gin.Context) {
	body, _ := c.GetRawData()
	fmt.Println(body)                        //[117 115 101 114 61 97 98 38 117 115 101 114 61 99 100]
	fmt.Println(string(body))                //user=ab&user=dadacd
	fmt.Println(c.GetHeader("Content-Type")) //application/x-www-form-urlencoded
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":

		// json解析到结构体
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"` //传入的age必须是int类型，不能是string类型，否则解析不了
		}
		var user User
		err := json.Unmarshal(body, &user)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(user)
	}
}
func main() {
	r := gin.Default()
	// r.GET("/query", _query)
	// r.GET("/param/:user_id", _param)
	// r.GET("/param/:user_id/:book_id", _param)
	// r.GET("/form", _form)
	r.GET("/raw", _raw)
	r.Run(":80")
}
