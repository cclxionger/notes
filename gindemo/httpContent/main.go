package main

import (
	"fmt"
	"net/http"
	"os"
)

func sayhello(w http.ResponseWriter, r *http.Request) {

	// _, err := fmt.Fprintln(w, "<h1> hello golang!")
	// if err != nil {
	// 	fmt.Println("sayhello err ", err)
	// }
	//如果让网页显示文件内容
	byt, _ := os.ReadFile("./hello.txt")
	_, err := fmt.Fprintln(w, string(byt))
	if err != nil {
		fmt.Println("sayhello err ", err)
	}

}
func main() {
	http.HandleFunc("/hello", sayhello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("serve failed,err:", err)
		return
	}
}
