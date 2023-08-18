package main

import (
	"fmt"
)

func Test01() {
	for i := 0;  i < 5; i++ {
		defer func() {
			fmt.Println(i) // 显而易见结果肯定全是5
		} ()
	}
}

func Test02() {
	for i := 0; i < 5; i++ {
		i := i
		defer func() {
			fmt.Println(i) // i := i中的左值i会被视为在闭包中申请的变量，所以闭包中的函数会优先使用闭包内部的变量
		} ()
	}
}

func main() {
	Test01()

	Test02()
}