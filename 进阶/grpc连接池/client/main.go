package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"context"
	"fmt"
	"grpc_pool/client/pbfile"
	"grpc_pool/pool"
	"time"
	"sync"
	"math/rand"
)

func WithClient() {
	begin := time.Now()

	wg := &sync.WaitGroup{}
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
			defer conn.Close()
			if err != nil {
				panic(err)
			}
			client := pbfile.NewHelloClient(conn)
			res, err := client.SayHello(context.Background(), &pbfile.Helloreq{
				Mes: "cdl" + fmt.Sprint(i),
			})
			if err != nil {
				panic(err)
			}
			fmt.Println(res)
		}()
	}
	wg.Wait()

	fmt.Println(time.Since(begin))
}

func WithPool() {
	p := pool.NewPool("localhost:8000")
	defer p.Close()

	begin := time.Now()
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000 ; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			conn := p.Get()
			p.Len()
			defer conn.Close(p)
		
			client := pbfile.NewHelloClient(conn.Value)
			time.Sleep(100 * time.Millisecond)
			res, err := client.SayHello(context.Background(), &pbfile.Helloreq{
				Mes: "cdl"+fmt.Sprint(i),
			})
			if err != nil {
				panic(err)
			}
			fmt.Println(res)
		}()
	}
	wg.Wait()
	fmt.Println(time.Since(begin))

}

func main() {
	rand.Seed(time.Now().UnixNano())
	WithPool()

	// WithClient()
}