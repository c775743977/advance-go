package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var (
	concurrence       int32                     // 用于查看当前并发量
	concurrence_limit = make(chan struct{}, 10) // 用于限制当前最大并发为10
)

// 模拟数据库io操作
func readDB() string {
	// atomic包的作用
	atomic.AddInt32(&concurrence, 1)
	fmt.Println("当前并发量:", concurrence)
	time.Sleep(200 * time.Millisecond)
	atomic.AddInt32(&concurrence, -1)
	return "ok"
}

// 模拟一个操作
func handler() {
	concurrence_limit <- struct{}{}
	readDB()
	<-concurrence_limit
}

func main() {
	for i := 0; i < 100; i++ {
		go handler()
	}
	time.Sleep(3 * time.Second)
}
