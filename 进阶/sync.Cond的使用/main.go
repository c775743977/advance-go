package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})    // 创建cond对象，需要传入一个Mutex或者RWMutex类型
	queue := make([]interface{}, 0, 10) // 创建一个队列，长度0，容量10
	
	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()        // 对出队操作进行加锁
		queue = queue[1:] // 第一个元素出队列
		fmt.Println("Removed from queue")
		c.L.Unlock() // 对出队操作解锁
		c.Signal()   // 通知Wait的协程可以继续工作
	}
	
	for i := 0; i < 10; i++ {
		c.L.Lock() // 对入队操作进行加锁
		if  len(queue) == 2 { // 判断队列长度是否为2
			c.Wait() // 主协程，也就是main函数进行阻塞并等待
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second) // 开启一个协程进行出队操作
		c.L.Unlock()                        // 对入队操作解锁，个人认为应该放在调用出队函数之前
	}
}

// 整个流程就是实现了一个容量为2的队列，当队列满时入队操作进入阻塞并等待，一旦有元素出队便通知入队操作的协程，解除其阻塞状态并继续入队