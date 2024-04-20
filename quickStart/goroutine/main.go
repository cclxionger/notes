package main


import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go doSothing(1, &wg)
	go doSothing(2, &wg)
	go doSothing(3, &wg) //这样的话什么都不会输出，因为goroutine还没执行完，main函数已经执行了
	// time.Sleep(4*time.Second)//这样并不一定输出，需要调整秒
	wg.Wait() //等待输出

	for i := 0; i < 3; i++ {
		go func(v int) {
			fmt.Println(v)
		}(i) //如果v传参，就会得到 三个2
	}
	time.Sleep(1*time.Second)

}
func doSothing(jobID int, wg *sync.WaitGroup) {
	log.Printf("start job(%d)\n", jobID)
	time.Sleep(3 * time.Second)
	log.Printf("end job(%d)\n", jobID)
	wg.Done()

}
