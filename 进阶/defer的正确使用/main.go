// defer可以使得代码更加灵活，并且能够及时释放资源
package main

import (
	"os"
)

func test1(ch <-chan string) error {
	for path := range ch { 
        file, err := os.Open(path) 
        if err != nil {
            return err
        }
        defer file.Close() 
        // 对文件进行操作的代码
    }
    return nil
}
// defer的调用是在函数结束时，并且多个defer时先声明的后调用，后声明的先调用
// 对于上述函数则会存在一个问题，读取到的文件名，并且打开此文件后，该文件并没有及时关闭，因为defer在等到整个函数结束时才执行
// 并且打开的第一个文件需要等到最后一个才能关闭，这样就造成了资源浪费，明明处理完一个文件就该关闭一个文件

func test2(ch <-chan string) error {
	for path := range ch {
		func(path) {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			// 对文件进行操作的代码
		}
	}
	return nil
}
// 通过这样一个匿名函数，将defer包含在函数中，就可以使得每个文件使用完毕后及时关闭
// 也可以再定义一个函数，而不是采用匿名函数的方式。