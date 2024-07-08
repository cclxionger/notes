package main

import (
	"fmt"
	"strings"
)

var m = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	path := make([]byte, 0)
	if strings.Contains(digits, "0") || strings.Contains(digits, "1") || digits == "" {
		return res
	}
	dfs(digits, 0, path, &res)
	return res

}
func dfs(digits string, index int, path []byte, res *[]string) { //这个index代表当前正在处理digits里面的某个索引，比如digits是23，index就是0或者1
	if len(path) == len(digits) {
		temp := string(path)
		*res = append(*res, temp)
		return
	}
	digit := int(digits[index] - '0') //获取到digits中的2和3
	letter := m[digit-2]
	// 这个i每次都从0开始，不是从index开始
	for i := 0; i < len(letter); i++ {
		path = append(path, letter[i])
		dfs(digits, index+1, path, res)
		path = path[:len(path)-1]
	}
}
func main() {
	fmt.Println(letterCombinations("23"))
}
