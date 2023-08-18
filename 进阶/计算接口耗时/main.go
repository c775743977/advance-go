package main

import (
	"fmt"
	"time"
)

func handler(a int, b int) string {
	now := time.Now()
	// 为什么不直接写defer fmt.Printf("handler耗时:%d ms\n", time.Since(now).Milliseconds())
	// 因为执行到defer这一行时会先将语句中的变量计算出来，然后在函数结束时再执行语句，所以此时的耗时就为程序执行到defer这一行的耗时
	defer func() {
		fmt.Printf("handler耗时:%d ms\n", time.Since(now).Milliseconds())
	}()
	if a > b {
		time.Sleep(100 * time.Millisecond)
		return "ok"
	} else {
		time.Sleep(200 * time.Millisecond)
		return "ok"
	}
}

func main() {
	handler(3, 4)
}