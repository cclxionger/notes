package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string //结构体如果小写，在html里面使用，页面会显示不出来如果是map小写的话可以
	Age    int
}

func sayhllo(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t, err := template.ParseFiles("./hello.html")
	if err != nil {
		fmt.Println("parse template failed", err)
		return
	}
	//渲染模板
	u1 := User{
		"ck",
		"男",
		18,
	}
	u2 := User{
		"lc",
		"女",
		20,
	}
	// hobbyList := []string{
	// 	"篮球",
	// 	"足球",
	// 	"双色球",
	// }
	err = t.Execute(w, map[string]interface{}{
		"u1": u1,
		"u2": u2,
		// "hobby": hobbyList,
	})
	if err != nil {
		fmt.Println("render template failed", err)
		return
	}
}
func main() {
	http.HandleFunc("/hello", sayhllo)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http serve failed,", err)
		return
	}
}
