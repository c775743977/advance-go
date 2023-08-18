package main

import (
	"context"

	"grpc_pool/client/pbfile"
	"grpc_pool/pool"
	"testing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
  )

  func BenchmarkGrpcClient(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
			defer conn.Close()
			if err != nil {
				panic(err)
			}
			client := pbfile.NewHelloClient(conn)
			_, err = client.SayHello(context.Background(), &pbfile.Helloreq{
				Mes: "cdl",
			})
			if err != nil {
				panic(err)
			}
		}()
	}
  }

  func BenchmarkGrpcPool(b *testing.B) {
	p := pool.NewPool("localhost:8000")
	defer p.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {	
		conn := p.Get()
		defer conn.Close(p)
		client := pbfile.NewHelloClient(conn.Value)
		_, err := client.SayHello(context.Background(), &pbfile.Helloreq{
			Mes: "cdl",
		})
		if err != nil {
			panic(err)
		}
	}
  }