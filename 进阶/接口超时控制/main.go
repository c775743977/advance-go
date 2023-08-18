// 控制外部接口超时，当访问外部接口时如果超过一定时间没有收到回复则直接向前端报错或者后端进行超时处理

package main

import (
	"fmt"
	"net/http"
	"time"
	"context"
)

// 模拟数据库crud操作
func readDB() string {
	time.Sleep(200 * time.Millisecond) // 模拟耗时

	return "ok"
}

// 使用context的话就需要通过管道来接受数据，以便退出阻塞
func readDB2(ch chan string)  {
	defer close(ch) // 关闭传入的管道，使其变只读

	time.Sleep(100 * time.Millisecond) // 模拟耗时

	ch <- "ok" // 写入数据
}


func handler(w http.ResponseWriter, r *http.Request) {
	var res string
	done := make(chan struct{}, 1) // 用于接收readDB()的完成信号，缓冲区建议设置1，若设置为0可能造成协程泄漏
	go func() {
		res = readDB()
		done <- struct{}{}
	}()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond): // time.After()函数会在设置的时间之后返回一个只读的管道，元素类型为time.Time记录当前时间
		res = "timeout"
	}
	// time.After设定的时间就是上限，若readDB()先执行完则通过done来解除阻塞，若达到time.After设置的时间则通过time.After来解除阻塞

	fmt.Fprintln(w, res)
}

// 使用context包中的WithTimeout实现相同功能
func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recv get request")
	// 定义100毫秒的超时context
	ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Millisecond)
	defer cancel()

	ch := make(chan string)

	var res string

	// 协程执行
	go readDB2(ch)

	// 如果到了超时时间readDB2还没向管道中发送数据则说明超时，会执行cancel()，然后会向ctx.Done()传入信息，解除阻塞
	select {
	case <- ctx.Done():
		fmt.Println(ctx.Err())
		res = "timeout"
	case res = <- ch:
	}

	fmt.Fprintln(w, res)
}

func main() {
	http.HandleFunc("/", handler)

	http.HandleFunc("/2", handler2)

	http.ListenAndServe(":8080", nil)
}