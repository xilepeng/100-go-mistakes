package main

import (
	"context"
	"fmt"
)

func listing1() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)

		switch i {
		default:
		case 2:
			break
		}
	}
}

func listing2() {
loop:
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)

		switch i {
		default:
		case 2:
			// 等价于 return
			break loop // Breaks out of the outer loop
		}
	}
}

func listing3(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ch:
			// Do something
		case <-ctx.Done():
			break
		}
	}
}

func listing4(ctx context.Context, ch <-chan int) {
loop:
	for {
		select {
		case <-ch:
			// Do something
		case <-ctx.Done():
			break loop// 终止 loop 标签关联循环，而不是 select
		}
	}
}

func main() {
	listing1()
}
