package main

import (
	"net"

	"grpc_pool/server/pbfile"
	
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()

	pbfile.RegisterHelloServer(server, &pbfile.HelloService{})

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	server.Serve(listener)
}