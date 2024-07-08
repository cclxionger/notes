package main

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

/*
Regex   Meaning
(?i)	关闭大小写敏感
.	匹配任何单一字符
?	匹配前面的元素一次要么根本不匹配。
+	匹配前面的元素一次或多次。
*	匹配前一个元素的零次或多次。
^	匹配字符串中的起始位置。
$	匹配字符串中的结束位置。
|	交替运算符
\s	匹配空格字符。
\w	匹配一个单词字符；等同于 [a-zA-Z_0-9]
*/
func main() {
	//MatchString返回的是bool和err
	//直接用 MatchString 匹配字符串会影响性能
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	//regexp.Compile返回的是r  *regexp.Regexp和err，r.MatchString 返回的是bool，和上面一样
	//Compile 函数解析正则表达式，如果成功，则返回可用于匹配文本的 Regexp 对象。编译的正则表达式产生更快的代码。假如正则表达式非法，那么 Compile() 方法回返回 error ,
	//而 MustCompile() 编译非法正则表达式时不会返回 error ，而是返回 panic ，程序会崩溃
	r, _ := regexp.Compile("p([a-z]+)ch")
	match = r.MatchString("peach")
	fmt.Println(match)
	//FindString() 用来返回第一个匹配的结果。
	//如果没有匹配的字符串，那么它回返回一个空的字符串
	fmt.Println(r.FindString("peach punch"))
	//注意FindAllString和FindString()区别，这个可以后面参数是-1返回全部的，否则返回指定的，如果不够指定的，只能返回最多的
	fmt.Println(r.FindAllString("peach punch", -1))
	//寻找第一个匹配的索引,返回的是一维切片
	a := r.FindStringIndex("peach punch")
	fmt.Println(a)
	nedStr := "peach punch"
	//寻找匹配的全部索引，返回的是二维切片
	//前闭后开
	b := r.FindAllStringIndex(nedStr, -1)
	//这个for range 可能有点难懂，j代表每一行的对应的值,返回的是n行2列数组，所以直接使用j[0]和j[1]
	for _, j := range b {
		match := nedStr[j[0]:j[1]]
		fmt.Printf("%s at %d:%d\n", match, j[0], j[1])
	}
	//看不懂for range 看for  一样的
	for i := 0; i < len(b); i++ {
		resStr := nedStr[b[i][0]:b[i][1]]
		fmt.Printf("第%d次匹配到字符是%s,对应索引是%d-%d\n", i, resStr, b[i][0], b[i][1])
	}
	//Split 函数
	var data = `1,2,3`
	//反引号（```）用于定义原始字符串字面量。原始字符串包含在两个反引号之间，
	//并且其中的所有字符都按字面意思处理，不进行转义
	//使用双引号（需要转义换行符）,使用反引号（无需转义），``会保存原来的，"会对\n之类的进行转移，如果要匹配\n，那么必须使用``
	sum := 0
	re, _ := regexp.Compile(",")
	vals := re.Split(data, -1)
	for _, val := range vals {
		n, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		sum += n
	}
	fmt.Println(sum)
	//匹配到了之后转换成对应的字符串
	fmt.Println(r.ReplaceAllString("a peach", "fruit"))
	//匹配到了之后转换成对应的函数
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
	//当然也可以使用ReplaceAllStringFunc，对应 strings.ToLower)
	fmt.Println(r.ReplaceAllStringFunc(string(in), strings.ToLower))
	//Go 正则表达式捕获组(字符串是个数组),使用 FindStringSubmatch 函数
	//圆括号 () 用于创建捕获组,这个例子中相当于匹配[a-z]
	//r, _ := regexp.Compile("p([a-z]+)ch")
	//返回的索引的字符串，索引0为匹配到的整个字符串，从1开始对应第一个()括号对应的内容
	websit := []string{"peach", "punch"}
	for _, web := range websit {
		resStr := r.FindStringSubmatch(web)
		fmt.Println(resStr)

	}

}
