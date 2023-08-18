// 关于空结构体struct{}{}的应用
package main

import "fmt"
import "unsafe"
import "time"

type EST struct {

}

func main() {
	var b EST
	var c struct{}
	// b和c其实是等价的

	fmt.Printf("b address %p size %d\n", &b, unsafe.Sizeof(b))
	fmt.Printf("c address %p size %d\n", &c, unsafe.Sizeof(c))
	// b和c的大小都是0，并且地址都是一样的

	// 空结构体模拟java和python中的Set
	// Set其实就是没有value的map，里面的元素不能重复
	students := make(map[string]struct{}, 10)
	students["Tom"] = struct{}{}
	students["Alex"] = EST{}
	fmt.Println(len(students))
	for k, _ := range students {
		fmt.Println(k)
	}

	// 空结构体使进程阻塞，达到WaitGroup的效果
	ch := make(chan struct{}, 0)
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("协程工作完成")
		ch <- struct{}{}
	}()
	<-ch
	fmt.Println("主进程工作完成")
}