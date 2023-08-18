package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		i := i       // 注意这里的同名变量覆盖
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}
