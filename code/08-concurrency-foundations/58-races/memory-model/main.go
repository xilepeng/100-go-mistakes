package main

import "fmt"

// 创建一个 goroutine 一定 happens before 执行这个 goroutine
func listing1() {
	i := 0
	go func() {
		i++
	}()
}

// goroutine 的退出不能保证 happens before 其他事件，所以下面的代码有数据竞争
func listing2() {
	i := 0
	go func() {
		i++
	}()
	fmt.Println(i)
}

// 执行顺序：变量加1<从channel发送<从channel读取<变量读取
func listing3() {
	i := 0
	ch := make(chan struct{})
	go func() {
		<-ch           // 3
		fmt.Println(i) // 4
	}()
	i++              // 1
	ch <- struct{}{} // 2
}

func listing4() {
	i := 0
	ch := make(chan struct{})
	go func() {
		<-ch
		fmt.Println(i)
	}()
	i++
	close(ch)
}

func listing5() {
	i := 0
	ch := make(chan struct{}, 1)//channel是缓冲的，会导致数据竞争
	go func() {
		i = 1
		<-ch
	}()
	ch <- struct{}{}
	fmt.Println(i)//对变量i的读和写可能同时发生
}

func listing6() {
	i := 0
	ch := make(chan struct{})//非缓冲channel，无数据竞争
	go func() {
		i = 1
		<-ch
	}()
	ch <- struct{}{}
	fmt.Println(i)
}
