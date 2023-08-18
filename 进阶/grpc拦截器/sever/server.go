package main

import (
	pb "grpc-interceptor/sever/pbfile"

 	"google.golang.org/grpc"

	"net"
	"fmt"
	"context"
	"time"
)

var (
	limit = make(chan struct{}, 10) // 用于限制并发，最多10个
)

// 拦截器函数参数完全照搬文档的UnaryServerInterceptor
// info就是实例化的server对象
// handler就是编写的grpc服务，需要调用服务一样调用，req就是传入的请求，resp就是返回的响应
func LimitAndTimeInterceptor(ctx context.Context, req interface{}, 
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 并发+1
		limit <- struct{}{}
		// 计时
		start := time.Now()
		resp, err = handler(context.Background(), req)
		fmt.Println("concurrency:", len(limit))
		time.Sleep(100 * time.Millisecond) // 为了使数据更直观
		<- limit
		fmt.Println("time:", time.Since(start))
		return resp, err
	}

func main() {
	// 需要在初始化时传入写好的拦截器函数
	server := grpc.NewServer(grpc.UnaryInterceptor(LimitAndTimeInterceptor))

	pb.RegisterHelloServiceServer(server, &pb.HelloService{})

	listener, err := net.Listen("tcp",":8888")
	defer listener.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	server.Serve(listener)
}