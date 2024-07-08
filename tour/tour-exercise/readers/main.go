package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: 为 MyReader 添加一个 Read([]byte) (int, error) 方法。
func (mr MyReader) Read(readB []byte) (int, error) {
	for i, _ := range readB {
		readB[i] = 'A'
	}
	return len(readB), nil
}
func main() {
	reader.Validate(MyReader{})
}
