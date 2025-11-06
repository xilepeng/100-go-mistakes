package main

import (
	"context"
	"fmt"
	"time"
)

// 取消信号
func f1(t int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1) // 创建一个可取消的上下文
	defer cancel()                       // 延迟调用取消函数
	go func() {
		time.Sleep(4 * time.Second)
		cancel() // 调用取消函数
	}()
	// 感知上下文的取消信号
	select {
	case <-c1.Done(): // 如果上下文已经完成，返回它的错误信息
		fmt.Println("f1():", c1.Err()) // context canceled
		return
		// 代表程序正常操作。意味着如果该程序执行超时就会立刻被取消。
	case r := <-time.After(time.Duration(t) * time.Second): // 持续从 ch 中读取信息
		fmt.Println("f1():", r) // f1(): 2025-11-05 10:16:37.859308875 +0800 CST m=+3.000235959
	}
	return
}

// 最后期限
func f2(t int) {
	c2 := context.Background()
	c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second) // 创建4秒后超时的长下文
	defer cancel()
	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()
	select {
	case <-c2.Done(): // 如果上下文已经完成，返回它的错误信息
		fmt.Println("f2():", c2.Err()) // f2(): context canceled
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f2()取消时间:", r) // f2(): 2025-11-05 10:16:40.862259583 +0800 CST m=+6.003186542
	}
	return
}

// 最后期限
func f3(t int) {
	c3 := context.Background()
	deadline := time.Now().Add(time.Duration(2*t) * time.Second) // 指定时间点
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()
	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()
	select {
	case <-c3.Done():
		fmt.Println("f3():", c3.Err()) // f3(): context canceled
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f3():", r) // f3(): 2025-11-05 10:16:43.863456833 +0800 CST m=+9.004384125
	}
	return
}

func main() {
	delay := 3
	//delay := 5

	fmt.Println("delay=", delay)
	f1(delay)
	f2(delay)
	f3(delay)
}

/*
time.After() 函数调用的返回值。它们代表程序正常操作。意味着如果该程序执行超时就会立刻被取消。
delay= 3
f1(): 2025-11-05 10:16:37.859308875 +0800 CST m=+3.000235959
f2(): 2025-11-05 10:16:40.862259583 +0800 CST m=+6.003186542
f3(): 2025-11-05 10:16:43.863456833 +0800 CST m=+9.004384125

c.Done() 函数调用的返回值
delay= 5
f1(): context canceled
f2(): context canceled
f3(): context canceled
*/
