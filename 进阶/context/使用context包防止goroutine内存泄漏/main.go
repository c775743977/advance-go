package main

import (
	"fmt"
	"context"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    ch := func(ctx context.Context) <-chan int {
        ch := make(chan int)
        go func() {
            for i := 0; ; i++ {
                select {
                case <- ctx.Done():
                    return
                case ch <- i:
                }
            }
        } ()
        return ch
    }(ctx)

    for v := range ch {
        fmt.Println(v)
        if v == 5 {
            cancel()
            break
        }
    }
}

// 下面的 for 循环停止取数据时，就用 cancel 函数，让另一个协程停止写数据。
// 如果下面 for 已停止读取数据，上面 for 循环还在写入，就会造成内存泄漏。