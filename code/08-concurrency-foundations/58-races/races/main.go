// package races

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func listing1() {
	i := 0

	go func() {
		i++
	}()

	go func() {
		i++
	}()
	fmt.Println(i) // 0
}

/*
➜ go run -race main.go

Found 1 data race(s)
*/

func listing2() {
	var i int64

	go func() {
		atomic.AddInt64(&i, 1) // 原子操作加 1
	}()

	go func() {
		atomic.AddInt64(&i, 1)
	}()
}

// 原子操作不能被中断，避免同时进行两次访问

// 互斥锁确保最多只有一个 goroutine 访问临界区
func listing3() {
	i := 0
	mutex := sync.Mutex{}

	go func() {
		mutex.Lock()
		i++
		mutex.Unlock()
	}()

	go func() {
		mutex.Lock()
		i++
		mutex.Unlock()
	}()
}

// 跨 goroutine 通信
func listing4() {
	i := 0
	ch := make(chan int)

	go func() {
		ch <- 1 // 通知 goroutine 进行加 1 操作
	}()

	go func() {
		ch <- 1
	}()

	i += <-ch // 加上从 channel 中读取的值
	i += <-ch
}

// 无数据竞争，但有竞争条件。当行为取决于无法控制的事件顺序或时间时，就会出现竞争条件。
// 互斥锁保证无数据竞争，但无确定结果
func listing5() {
	i := 0
	mutex := sync.Mutex{}

	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		i = 1
	}()

	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		i = 2
	}()

	_ = i
}

func main() {
	listing1()
	listing2()
	listing3()
	listing4()
	listing5()
}
