package main

import (
	"bytes"
	"fmt"
	"text/template"
)

var Template = `
| 时间 	    | 第一值班人 		| 第二值班人 |
| :----		| :----     	       | :----     |
{{- range . }}
| {{ .Date}}   {{range .Persons}}          | {{.}}   {{end}} |
{{- end -}}
`

type RowUnit struct {
	Date    string
	Persons []string
}
//当然不使用模板当然也可以，使用 fmt.Sprintf() 同样可以实现
func main() {
	t, err := template.New("duty").Parse(Template)
	if err != nil {
		return
	}
	users := []RowUnit{
		{"20230505",
			[]string{"张三", "李四"},
		},
	}
	text := new(bytes.Buffer)
	err = t.Execute(text, users)
	if err != nil {
		return
	}
	fmt.Println(text)

}
