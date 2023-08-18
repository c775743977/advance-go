package main

import (
	pb "grpc-interceptor/client/pbfile"

	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	limit = make(chan struct{}, 10)
	ctx = context.Background()
)

func LimitInterceptor(ctx context.Context, method string, req, reply interface{}, 
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		limit <- struct{}{}
		fmt.Println("concurrency:", len(limit))
		err := invoker(ctx, method, req, reply, cc)
		<- limit
		return err
	}

func TimeInterceptor(ctx context.Context, method string, req, reply interface{}, 
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		begin := time.Now()

		err := invoker(ctx, method, req, reply, cc)

		fmt.Println("time:", time.Since(begin))
		return err
	}

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()), 
	grpc.WithChainUnaryInterceptor(LimitInterceptor, TimeInterceptor))
	defer conn.Close()
	
	if err != nil {
		fmt.Println(err)
		return
	}

	client := pb.NewHelloServiceClient(conn)

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res, _ := client.SayHello(context.Background(), &pb.HelloRequest{
				Name: "tom",
			})
			
			fmt.Println(res)
		}()
	}
	wg.Wait()

}