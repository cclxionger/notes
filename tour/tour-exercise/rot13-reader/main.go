package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

/*
io.Reader接口为

	type Reader interface {
		Read(p []byte) (n int, err error)
	}
*/

// 加密字母 'a' 到 'z' 对应的 ASCII 码是 97 到 122，其余不变
// 注意要传指针类型，否则不会改变，引用方法
func (r13 *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r13.r.Read(p)
	if err != nil {
		if err == io.EOF { // 如果读到结尾了
			return
		} else {
			fmt.Println(err)
			return
		}

	}
	// 不能用for range 因为range 改变是v的副本
	// for _, v := range p {
	// 	// 是字母进行加密
	// 	if v >= 'a' && v <= 'z' {
	// 		v = ((v - 'a' + 13) % 26) + 'a'
	// 	}

	// }
	for i := 0; i < n; i++ {
		v := p[i]
		if v >= 'a' && v <= 'z' {
			v = ((v - 'a' + 13) % 26) + 'a'
		}
		p[i] = v
	}
	return n, nil
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := &rot13Reader{s}
	io.Copy(os.Stdout, r)

}
