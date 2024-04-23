package main

import (
	"fmt"
	"io"
)

func main() {
	// rea := strings.NewReader("Hello World!")
	// b := make([]byte, 10)
	// for {

	// 	n, err := rea.Read(b)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println(n)
	// 	fmt.Printf("%q\n", b[:n]) //如果不截取到每次的n，如果上次n-1有值
	// 	//这次n-1的地方没有值，那么不会覆盖掉

	// }

	reader := MyReader{}
	buffer := make([]byte, 1)
	count := 0
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF || n == 0 {
			break
		}
		fmt.Printf("%c\t", buffer[0])
		count++
		if count >= 10 {
			break
		}
	}

}

// exercise reader
type MyReader struct { //好像是一个结构体可以调用Read方法
}

func (r MyReader) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
}
