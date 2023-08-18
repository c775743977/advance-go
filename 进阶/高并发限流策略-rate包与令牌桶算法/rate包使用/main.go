package main

import (
	"fmt"
	"time"
	"context"
	_ "sync"

	"golang.org/x/time/rate"
)

// rate包中主要是其中Limiter类型，能够平滑的限制频率（接口的访问请求等） 使用方法有allow, wait, reserve
// rate包主要依靠令牌桶算法实现

// wait是最常用的
func use_wait() {
	limiter := rate.NewLimiter(10, 5) // 10为频率，即每1/10秒提供一个令牌，5为最大容量，即总共能够有5个令牌存在
	go func() {
		var count int = 0
		for {
			fmt.Println(count, time.Now())
			count++
			// 使用wait函数后，没有拿到令牌的时候就只能阻塞。可以分析最开始有5个令牌满的，此后每提供一个令牌就会被消耗一个
			// 设置的频率是1/10即0.1秒，所以在sleep这5秒内，每秒有10个令牌，所以一共输出55个 count最后等于54
			// 如果不加wait这里的值就会在5秒内快速增加
			err := limiter.Wait(context.Background())
			// Wait(ctx) == WaitN(ctx, 1) WaitN函数中的N指的是执行一次消耗N个令牌，令牌不够的时候就只能阻塞等待
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}()
	time.Sleep(5 * time.Second)
}

func use_allow() {
	limiter := rate.NewLimiter(10, 5) 
	go func() {
		var count int = 0
		for {
			count++
			// Allow()函数不会阻塞，但是在令牌不足的时候会返回false
			// 比如这里的代码在返回true时打印count，根据上面讲的桶里初始有5个令牌，如果Sleep时间为101毫秒，
			// 那么正好提供一个新令牌，所以count会输出6次，但是第六次输出时数值已经非常大了，因为这个过程中count一直在累加
			// 说明Allow没有阻塞程序进行
			if limiter.Allow() {
				fmt.Println(count)
			}
		}
	}()
	time.Sleep(101 * time.Millisecond)
}

func use_reserve() {
    limiter := rate.NewLimiter(10, 5)
    // var wg = sync.WaitGroup{}
    // for {
    //     wg.Add(1)
    //     go func() {
    //         defer wg.Done()
    //         r := limiter.ReserveN(time.Now(), 10)
    //         if r.OK() {
    //             fmt.Printf("r1 delay %s\n", r.Delay())
    //         } else {
    //             fmt.Println("not ok")
    //         }
    //     }()
    // }

	go func() {
		var count int = 0
		for {
			// Reserve方法会返回结果告诉你每个待执行的阻塞任务需要等待多久
			count++
			fmt.Println(count)
		    r := limiter.ReserveN(time.Now(), 1)
            if r.OK() {
                fmt.Printf("r1 delay %s\n", r.Delay())
            } else {
                fmt.Println("not ok")
            }	
		}
	}()
	time.Sleep(200 * time.Millisecond)
    // wg.Wait()
}

func main() {
	// use_wait()

	// use_allow()

	use_reserve()
}