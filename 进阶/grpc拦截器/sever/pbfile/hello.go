package pbfile

import (
	"context"
)

type HelloService struct {

}

func(this *HelloService) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{
		Age: 18,
	}, nil
}