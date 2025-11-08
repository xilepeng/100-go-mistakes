package main

import (
	"fmt"
	"time"
)

func main() {
	messageCh := make(chan int, 10)
	disconnectCh := make(chan struct{})

	//go listing1(messageCh, disconnectCh)
	go listing2(messageCh, disconnectCh)

	for i := 0; i < 10; i++ {
		messageCh <- i
	}
	disconnectCh <- struct{}{}
	time.Sleep(10 * time.Millisecond)
}

// switch 的 case 语句是依赖顺序的
// select 的 case 语句是随机选择的：防止可能的饥饿
func listing1(messageCh <-chan int, disconnectCh chan struct{}) {
	for {
		select {
		case v := <-messageCh:
			fmt.Println(v)
		case <-disconnectCh:
			fmt.Println("disconnection, return")
			return
		}
	}
}

// case 语句随机执行

// 如果想在连接中断导致返回前接收所有消息，解决方案：
// 从 messageCh 或 disconectCh 接收
// 如果收到断开链接消息

// 读取 messageCh 中的所有现有消息
// 然后返回
func listing2(messageCh <-chan int, disconnectCh chan struct{}) {
	for {
		select {
		case v := <-messageCh:
			fmt.Println(v) // 0
		case <-disconnectCh:
			for {
				select {
				case v := <-messageCh: // 读取剩余的消息
					fmt.Println(v) // 123456789
				default: // 仅当其他情况都不匹配时，才会选择在 select 语句中使用 default
					fmt.Println("disconnection, return")
					return // 意味着只有收到 messageCh 中的所有剩余消息后才会返回
				}
			}
		}
	}
}

// 0123456789
// disconnection, return
