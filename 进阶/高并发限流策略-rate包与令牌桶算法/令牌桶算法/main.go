package main

import (
	"fmt"
	"time"
	"context"
)

// 令牌桶的大致原理就是两个数值，一个是最大数量b，一个是频率r。桶里最多只能存放b个令牌（初始为满），每次任务调用都必须消耗一个令牌
// 令牌会以1/r(秒)的速度补充， 如r为5，则1/5即0.2秒补充一个。
// 根据实际情况设置b和r的值
// 比如要限制一段时间内的并发量，就把b设置为1，r设置为理想的并发频率
// 比如要限制某一时刻突然爆发的大量并发，则把b设置为理想的最高并发量即可

type Limiter struct {
	b int
	r float64
	ch chan struct{}
	context LimiterContext
}

type LimiterContext struct {
	ctx context.Context
	Stop context.CancelFunc
}

// func NewLimiter(rr int, bb int) *Limiter {
// 	limiter := &Limiter{
// 		b: bb,
// 		r: 1.0/float64(rr),
// 		ch: make(chan struct{}, bb),
// 	}

// 	// 初始化时装满桶
// 	for i := 0;i < bb; i++ {
// 		limiter.ch <- struct{}{}
// 	}

// 	// 提供令牌
// 	go func() {
// 		// 最好有通知协程停止的方式，否则造成协程泄漏，需要改进
// 		for {
// 			// 频率r要换算成毫秒，不能直接直接换成秒，因为Duration是取整型
// 			time.Sleep(time.Duration(limiter.r * 1000) * time.Millisecond)
// 			if len(limiter.ch) < limiter.b {
// 				limiter.ch <- struct{}{}
// 			}
// 		}
// 	}()

// 	return limiter
// }

// 引入context包
func NewLimiter(ctx context.Context, rr int, bb int) *Limiter {
	limiter := &Limiter{
		b: bb,
		r: 1.0/float64(rr),
		ch: make(chan struct{}, bb),
	}

	limiter.context.ctx, limiter.context.Stop = context.WithCancel(ctx)

	// 初始化时装满桶
	for i := 0;i < bb; i++ {
		limiter.ch <- struct{}{}
	}

	// 提供令牌
	go func() {
		// 最好有通知协程停止的方式，否则造成协程泄漏，需要改进
		for {
			select {
			case <-limiter.context.ctx.Done():
				fmt.Println("令牌桶已暂停")
				return
			// 每过limiter.r秒就补充一个令牌，由于channel的特性，当桶满时就会阻塞，等待有令牌取出时再解除阻塞
			case <-time.After(time.Duration(limiter.r * 1000) * time.Millisecond):
				limiter.ch <- struct{}{}
			}
		}
	}()

	return limiter
}

func(this *Limiter) Wait() {
	<- this.ch
	return
}

// limit为频率，即每秒补充的令牌数量
func(this *Limiter) SetNewLimit(limit int) {
	this.r = 1.0/float64(limit)
}

func(this *Limiter) Cancel() {
	this.context.Stop()
}

func main() {
	limiter := NewLimiter(context.Background(), 10, 5)

	go func() {
		var count int = 0
		for {
			fmt.Println(count, time.Now())
			count++
			if count == 20 {
				limiter.SetNewLimit(5)
				fmt.Println("开始限速")
			}
			// 最后输出count为54，并且每秒输出10次
			limiter.Wait()
		}
	}()
	time.Sleep(5 * time.Second)
}