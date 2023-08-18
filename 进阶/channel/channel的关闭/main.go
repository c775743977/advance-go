package main

import (
	"fmt"
	"time"
)

func test1() {
	ch := make(chan int, 10)

	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()

	// 报错，因为管道关闭后不能再进行写操作，但是可以进行读
	for i := 1; i <= 10; i++ {
		ch <- i
		if i == 5 {
			close(ch)
		}
	}

	time.Sleep(time.Second)
}

func test2() {
	ch := make(chan int, 10)

	for i := 1; i <= 5; i++ {
		ch <- i
	}

	close(ch)

	for k := range ch {
		fmt.Println(k) // 读操作正常
	}

	time.Sleep(time.Second)
}

func test3() {
	ch := make(chan int, 10)

	for i := 1; i <= 5; i++ {
		ch <- i
	}

	close(ch)

	go func() {
		for {
			fmt.Println(<-ch) // 读操作正常，但是后面只能读到0
		}
	}()

	time.Sleep(time.Millisecond)
}

func main() {
	test3()
}