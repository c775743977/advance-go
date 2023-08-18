// WithTimeout类似于之前做过的超时控制的封装

package main

import (
	"context"
	"fmt"
	"time"
)

func test1() {
	ctx, cancel := context.WithTimeout(context.Background(), 500 * time.Millisecond)

	defer cancel()

	time.Sleep(1 * time.Millisecond)
	select {
	case <- ctx.Done(): // 一旦超时或者执行cancel函数，就会关闭Done()函数返回的管道，从而解除阻塞
		fmt.Println(ctx.Err()) // Err()可以捕获到关闭原因
	case <- time.After(500 * time.Millisecond):
		fmt.Println("over")
	}
}

func test2() {
	// 当context有继承并且继承的context再次设置了一个WithTimeout，那么以最先超时的那个设置为准
	ctx, cancel := context.WithTimeout(context.Background(), 200 * time.Millisecond)
	t1 := time.Now() // ctx的存活时间
	defer cancel()

	time.Sleep(100 * time.Millisecond)

	ctx2, cancel2 := context.WithTimeout(ctx, 500 * time.Millisecond)
	t2 := time.Now() // ctx2的存活时间
	defer cancel2()

	select {
	case <- ctx2.Done():
		fmt.Printf("t1:%v\tt2:%v\n", time.Since(t1), time.Since(t2)) // 输出t1:200, t2:100，ctx2设置500才超时，但是因为ctx的200到了，所以ctx2也跟着结束了
	}
}

func test3() {
	ctx, cancel := context.WithCancel(context.Background()) //WithCancel就是调用cancel()时，ctx结束
	t := time.Now()

	go func() {
		time.Sleep(500 * time.Millisecond)
		cancel()
	}()

	select {
	case <- ctx.Done():
		fmt.Println("t:", time.Since(t))
	}
}

func main() {
	test1()
}