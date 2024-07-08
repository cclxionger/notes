package main

import "golang.org/x/tour/wc"

func WordCount(s string) map[string]int {
	// 如果使用strings.Fields

	// m := make(map[string]int)
	// strArr := strings.Fields(s)
	// for _, word := range strArr {
	// 	m[word]++
	// }
	// return m

	// 如果不使用strings.Fields
	// 使用strings.SplitN
	// m := make(map[string]int)
	// strArr := strings.SplitN(s, " ", -1) // 确定每个字段中间都只有一个空格
	// for _, word := range strArr {
	// 	m[word]++
	// }
	// return m

	// 如果都不使用模块
	m := make(map[string]int)
	var word string
	var isWord bool
	isWord = false
	for _, v := range s {
		if v == ' ' { // 如果是空格
			if isWord == true { // 之前是单词，现在遇到空格了，说明这个单词结束了
				m[word]++
				word = ""
				isWord = false
			}
		} else { // 不是空格
			isWord = true
			word += string(v)
		}
	}
	// 如果最后一个单词没有空格，那么结束退出循环了，但是并没有第一个if条件，把word加入map，所以最后还要把最后一个单词加入计算
	if isWord {
		m[word]++
	}
	return m
}

func main() {
	wc.Test(WordCount)

}
