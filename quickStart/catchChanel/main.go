package main

import (
	"errors"
	"fmt"
	"time"
)

func Receive(receiver chan<- string, str string) {
	receiver <- str
}
func Send(sender <-chan string, receiver chan<- string) string {
	str := <-sender
	receiver <- str
	return str

}
func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	Receive(ch1, "666")
	Send(ch1, ch2)
	fmt.Println(<-ch2)

	ch := make(chan string)
	go func() { //必须要有这个
		ch <- "str"
	}()
	fmt.Println(<-ch)

	ch3 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch3 <- i
	}
	res := 0
	close(ch3) //要把ch3的管道关闭掉，才能获取ch3管道的范围
	for v := range ch3 {
		res += v
	}
	fmt.Println(res)
	//
	chWork := make(chan string)
	go worker(chWork) //这样是不能输出的
	<-chWork          //必须要 读取chWork实现同步， 前提是chWork里面要有数据，才能读取

	//select语句
	chStr := make(chan string)
	go strWork(chStr)
	chInt := make(chan int)
	go intWork(chInt) //这样是没有输出结果的
	// select {          //这样输出的顺序是不一定的，需要有一个计时
	// case <-chStr:
	// 	fmt.Println("get value from strWork")
	// case <-chInt:
	// 	fmt.Println("get value from intWork")
	// } //这样会输出string\n"get value from strWork" //如果想输出俩个，需要用一个for循环
	for i := 0; i < 2; i++ { //这个2就是要输出俩个
		select {
		case <-chStr:
			fmt.Println("get value from strWork")
		case <-chInt:
			fmt.Println("get value from intWork")
		}
	}
	//异常处理
	fmt.Println("before panic")
	panicTest() //恢复程序，正在运行下面，如果没有recover，则不会运行下面的
	fmt.Println("after panic")

	//斐波那契数列用管道和goroutine写
	c := make(chan int, 5)
	go fibonacci(cap(c), c) //斐波那契数列存在管道里面了，然后关闭管道，才能对管道进行forr循环
	for i := range c {
		fmt.Println(i)
	}

}
func strWork(ch chan string) {
	time.Sleep(1 * time.Second)
	fmt.Println("string")
	ch <- "666"
}
func intWork(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("int")
	ch <- 2
}
func worker(ch chan string) {
	fmt.Println("get into worker")
	ch <- "str"
}

func panicTest() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("recover panic")
		}
	}()
	panic(errors.New("this is a panic"))
}

func fibonacci(n int, c chan int) {
	var x, y = 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
