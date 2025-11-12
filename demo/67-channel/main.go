package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	//ch1 := make(chan int)
	ch2 := make(chan int, 0) // 非缓冲 channel (同步 channel),发送者将阻塞，直到接受者从 channel 接收到数据
	defer close(ch2)
	go func() {
		ch2 <- 2
	}()
	go func() {
		message2 := <-ch2
		fmt.Println("message2=", message2)
	}()

	ch3 := make(chan int, 2) // 缓冲 channel 已满，它将阻塞，直到接受者 goroutine 收到消息
	defer close(ch3)
	ch3 <- 1 // 非阻塞
	ch3 <- 2 // 阻塞(无法解除阻塞的通道将死锁)
	go printer(ch3)
	wg.Wait()
}
func printer(ch chan int) {
	for ch != nil {
		select {
		case v, open := <-ch:
			if !open {
				ch = nil
				break
			}
			fmt.Println(v)
		}
	}
}
