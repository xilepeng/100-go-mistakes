package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 请在 goroutine 之前调用 Add 方法
func listing1() {
	wg := sync.WaitGroup{}
	var v uint64
	// Add 应在父 goroutine 中启动，并在 goroutine 启动前完成
	// 创建3个goroutine
	for i := 0; i < 3; i++ {
		go func() {
			wg.Add(1) // 不在父 goroutine 中调用，不能保证我们已经想等待组指示我们要在调用 wg.Wait() 之前等待3个 goroutine
			atomic.AddUint64(&v, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(v)
}

func listing2() {
	wg := sync.WaitGroup{}
	var v uint64

	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			atomic.AddUint64(&v, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(v)
}

func listing3() {
	wg := sync.WaitGroup{}
	var v uint64

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			atomic.AddUint64(&v, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(v)
}

func main() {
	listing1() // 0 没有同步，得到一个不确定的值
	listing2() // 3
	listing3() // 3
}
