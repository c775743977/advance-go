package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	go func() {
		for {
			log.Printf("当前 goroutine 数量为：%d\n", runtime.NumGoroutine())
			time.Sleep(1 * time.Second)
		}
	}()
	for i := 0; i < 10; i++ {
		ch := make(chan time.Time, 4)
		go func() {
			select {
			case <-ch:
			case <-time.After(1 * time.Second):
			}
		}()
		go func() {
			t := <-time.After(3 * time.Second)
			ch <- t
		}()
		log.Println("create a goroutine")
		time.Sleep(1 * time.Second)
	}
	for {
	}
}
