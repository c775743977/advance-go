package main

import (
	"fmt"
	"sync"
	"runtime"
	_ "time"
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
		// 加上sleep可以观察到输出结果就是递增的
		// time.Sleep(time.Second)
	}
	wg.Wait()
}

// 这种现象的原因在于闭包共享外部的变量i，注意到，每次调用go就会启动一个goroutine，这需要一定时间；
// 但是，启动的goroutine与循环变量递增不是在同一个goroutine，可以把i认为处于主goroutine中。
// 启动一个goroutine的速度远小于循环执行的速度，所以即使是第一个goroutine刚起启动时，
// 外层的循环也执行到了最后一步了。由于所有的goroutine共享i，而且这个i会在最后一个使用它的goroutine结束后被销毁，
// 所以最后的输出结果都是最后一步的i==5。(我的测试结果不一定全是5，主要理解i是共享的)
