package main

import (
	"fmt"
	"time"
)

func merge1(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		for v := range ch1 { // 从 ch1 接收，然后从 ch2 接收，ch1 关闭之前不会收到 ch2 的消息
			ch <- v
		}
		for v := range ch2 {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

// <-chan 只读通道 chan<- 只写通道
func merge2(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		for {
			select { // 同时接收 ch1 ch2 的消息
			case v := <-ch1:
				ch <- v // 问题：如果ch1 关闭就不能工作了
			case v := <-ch2:
				ch <- v
			}
		}
		close(ch) // 问题：不可访问
	}()

	return ch
}

func merge3(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)
	ch1Closed := false
	ch2Closed := false

	go func() {
		for { // 问题：当两个 channel 之一被关闭时，for 循环将持续循环知道 channel 有数据，
			// 这意味着即使在另一个 channel 中没有收到新消息，for 循环也会继续循环
			// 浪费 CPU 周期，必须避免这种情况发生
			select {
			case v, open := <-ch1:
				if !open { // ch1 被关闭时处理
					ch1Closed = true
					break
				}
				ch <- v
			case v, open := <-ch2:
				if !open { // // ch1 被关闭时处理
					ch2Closed = true
					break
				}
				ch <- v
			}

			if ch1Closed && ch2Closed { // ch1 ch2 都被关闭，则关闭ch并返回
				close(ch)
				return
			}
		}
	}()

	return ch
}

// 从 nil channel 接收消息将永远阻塞，在 channel 关闭后将此 channel 分配给 nil
// 可以使用 nil channel 实现一个优雅的状态机，该状态机降从 select 语句中删除一个案例
func merge4(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		for ch1 != nil || ch2 != nil { // 如果至少一个 channel 不是 nil 则继续
			select {
			case v, open := <-ch1:
				if !open {
					ch1 = nil // 一旦 ch1 被关闭，将其赋值为 nil channel
					break
				}
				ch <- v
			case v, open := <-ch2:
				if !open {
					ch2 = nil // 一旦 ch2 被关闭，将其赋值为 nil channel
					break
				}
				ch <- v
			}
		}
		close(ch)
	}()

	return ch
}

func PrintChan(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		//time.Sleep(1 * time.Second)
		ch1 <- 1
		close(ch1)
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 2
		close(ch2)
	}()
	ch := merge1(ch1, ch2)
	//ch := merge2(ch1, ch2)
	//ch := merge3(ch1, ch2)
	//ch := merge4(ch1, ch2)
	PrintChan(ch)
}
