package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
// 函数参数传递是瞬时的，就是说每个goroutine开启时传入的参数就是当时i的值