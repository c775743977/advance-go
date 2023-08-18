// recover可以从panic中恢复，panic会被取消，recover必须用在defer中，且必须在defer的函数中直接调用
package main

import (
	"fmt"
)

func Test(name string) int {
	// 捕获panic
	defer func() {
		err := recover() // recover()会返回错误
		if err != nil {
			fmt.Println(err)
		}
	}()
	return 10 / len(name) // 当name为空字符串时，len(name)值为0，会报panic
}

func main() {
	fmt.Print(Test(""))
}

// gin框架中就对每个handler都加入了recover()，所以使用gin时会发现很多报错不会使服务器终止