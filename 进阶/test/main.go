package main

import (
    "fmt"
    "time"
    _"sync"
)

func main() {
    ch := make(chan struct{}, 15)
    go func() {
        for{
            if len(ch) > 17 {
                fmt.Println("out")
                <-ch
            }
            time.Sleep(100 * time.Millisecond)
            // <-ch
            fmt.Println("out")
        }
    }()

    for i := 0; i < 25; i++ {
        fmt.Println(len(ch))
        ch<-struct{}{}
    }
}