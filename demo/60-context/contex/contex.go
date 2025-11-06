package main

import (
	"context"
	"fmt"
	"time"
)

func f1() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	//deadline := time.Now().Add(2 * time.Second) // 指定时间点
	//ctx, cancel := context.WithDeadline(context.Background(), deadline)

	//ctx, cancel := context.WithCancel(context.Background()) // 创建一个可取消的上下文
	//cancel()// main context canceled

	defer cancel()
	//go handle(ctx, 1*time.Second)
	go handle(ctx, 3*time.Second)

	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func f2() {
	type key string
	// 处理上下文的键的最佳实践：创建未导出的自定义类型（没有风险，使用其他上下文其他的包不可能覆盖这个键已设置的值）
	const myKey key = "key"
	ctx := context.WithValue(context.Background(), myKey, "myValue")
	fmt.Println(ctx.Value(myKey))
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err()) // handle context deadline exceeded
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func main() {
	f2()
	//myValue

	f1()
	//因为过期时间大于处理时间，所以我们有足够的时间处理该请求
	//process request with 1s
	//main context deadline exceeded

	//如果我们将处理请求时间增加至 3s，整个程序都会因为上下文的过期而被中止
	//handle context deadline exceeded
	//main context deadline exceeded
	//多个 Goroutine 同时订阅 ctx.Done() 管道中的消息，一旦接收到取消信号就立刻停止当前正在执行的工作。

}
