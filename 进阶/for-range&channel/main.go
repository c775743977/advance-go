package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 5)
	ch<-1
	ch<-2
	ch<-3
	ch<-4
	ch<-5
	for k := range ch {
		if (k == 3) {
			break
		}
		fmt.Println(k)
	}

	fmt.Println(<-ch)
}