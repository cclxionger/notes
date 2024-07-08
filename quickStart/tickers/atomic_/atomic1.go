package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Atomic() {
	var result1 atomic.Int64 //使用原子，避免出现错误
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 100; c++ {
				result1.Add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(result1.Load())
}
