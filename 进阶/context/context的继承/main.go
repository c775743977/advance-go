// context可以继承，并且context里可以存放key-value

package main

import (
	"context"
	"fmt"
)

func step1(ctx context.Context) context.Context {
	// context.WithValue会继承传入的context，然后添加上kv键值对，返回一个新的、继承了传入context的、添加新键值对的context.Context
	return context.WithValue(ctx, "name", "cdl")
}

func step2(ctx context.Context) context.Context {
	return context.WithValue(ctx, "age", 25)
}

func step3(ctx context.Context) context.Context {
	return context.WithValue(ctx, "num", 10001)
}

func printKV(ctx context.Context) {
	fmt.Printf("name %s\n", ctx.Value("name"))
	fmt.Printf("age %d\n", ctx.Value("age"))
	fmt.Printf("num %d\n", ctx.Value("num"))
}

func main() {
	ctx1 := context.TODO() // context.TODO()与context.Background等效，都是返回一个空的context.Context

	ctx2 := step1(ctx1)

	ctx3 := step2(ctx2)

	ctx4 := step3(ctx3)

	printKV(ctx4)
}