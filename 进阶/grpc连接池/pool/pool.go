package pool

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"fmt"
	"time"
)

type pool struct {
	conns chan *pconn
	addrs string
	maxsize int
}

type pconn struct {
	Value *grpc.ClientConn
}

// 如果池里还有空间就直接放回去，如果池满则直接关闭
func(this *pconn) Close(p *pool) {
	select{
	case p.conns <- this:
	default: this.Value.Close()
	}
}

func(this *pool) Empty() {
	if len(this.conns) == 0 {
		fmt.Println("empty!")
	}
}

func NewPool(grpcAddrs string, maxsize int) *pool {
	myPool := &pool{
		conns: make(chan *pconn, maxsize),
		addrs: grpcAddrs,
		maxsize: maxsize,
	}
	for i := 0; i < maxsize; i++ {
		myPool.conns<-&pconn{
			Value: myPool.CreateConn(),
		}
	}
	return myPool
}

// 如果等待50ms还取不到连接则直接新建一个
func(this *pool) Get() *pconn {
	select {
	case con := <-this.conns:
		return con
	case <-time.After(50 * time.Millisecond):
		return &pconn{
			Value: this.CreateConn(),
		}
	}
}

func(this *pool) CreateConn() *grpc.ClientConn {
	conn, err := grpc.Dial(this.addrs, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return conn
}

func(this *pool) Close() {
	for k := range this.conns {
		k.Value.Close()
		if len(this.conns) == 0 {
			break
		}
	}
	close(this.conns)
}