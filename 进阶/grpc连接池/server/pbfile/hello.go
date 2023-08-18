package pbfile

import (
	"context"
)

type HelloService struct {

}

func(this *HelloService) mustEmbedUnimplementedHelloServer() {}

func(this *HelloService) SayHello(ctx context.Context, req *Helloreq) (*Hellores, error) {
	return &Hellores{
		Reply: "hello" + req.Mes,
	}, nil
}