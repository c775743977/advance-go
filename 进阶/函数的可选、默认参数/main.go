package main

import (
	"fmt"
)

type Connection struct {
	
}

type StuffClient *stuffClient

type stuffClient struct {
    conn    Connection
    timeout int
    retries int
}

type StuffClientOption func(*StuffClientOptions)
type StuffClientOptions struct {
    Retries int //number of times to retry the request before giving up
    Timeout int //connection timeout in seconds
}
func WithRetries(r int) StuffClientOption {
    return func(o *StuffClientOptions) {
        o.Retries = r
    }
}
func WithTimeout(t int) StuffClientOption {
    return func(o *StuffClientOptions) {
        o.Timeout = t
    }
}

var defaultStuffClientOptions = StuffClientOptions{
    Retries: 3,
    Timeout: 2,
}
func NewStuffClient(conn Connection, opts ...StuffClientOption) StuffClient {
    options := defaultStuffClientOptions
    for _, o := range opts {
        o(&options)
    }
    return &stuffClient{
        conn:    conn,
        timeout: options.Timeout,
        retries: options.Retries,
    }
}

func main() {
	x := NewStuffClient(Connection{})
	fmt.Println(x) // prints &{{} 2 3}
	x = NewStuffClient(
    Connection{},
    WithRetries(1),
	)
	fmt.Println(x) // prints &{{} 2 1}
	x = NewStuffClient(
    Connection{},
    WithRetries(1),
    WithTimeout(1),
	)
	fmt.Println(x) // prints &{{} 1 1}
}