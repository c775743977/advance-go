package main

import (
	"sort"
	"fmt"
)

func main() {
	var slice []int
	for i := 0; i < 10; i++ {
		// var num int
		// fmt.Scan(&num)
		slice = append(slice, i+i)
	}
	fmt.Println()

	for _, k := range slice {
		fmt.Println(k)
	}

	sort.Ints(slice)
	fmt.Println()

	for _, k := range slice {
		fmt.Println(k)
	}

	fmt.Println(IsExist(slice, -1))
}


// go中使用 sort.searchXXX 方法，在排序好的切片中查找指定的方法，但是其返回是对应的查找元素不存在时，
// 待插入的位置下标(元素插入在返回下标前)。
// 可以通过封装如下函数，达到目的。
func IsExist(slice []int, num int) (int, bool) {
	index := sort.SearchInts(slice, num)
	fmt.Println("index:", index)
	if (index >= len(slice)) || (slice[index] != num) {
		return index, false
	} else {
		return index, true
	}
}