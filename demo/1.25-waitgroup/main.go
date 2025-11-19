package main

import (
	"fmt"
	"sync"
)

func main() {
	Good()
	Bad()
}

// Go 调用 f 输入一个新的 goroutine，并将该任务添加到 [WaitGroup]。
// 当 f 返回时，任务从 WaitGroup 中移除。
func Good() {
	var wg sync.WaitGroup
	wg.Go(func() {
		// 自动 Add 和 Done
		fmt.Println("Hello Go 1.25")
	})
	wg.Wait()
}

// ➜  1.25-waitgroup git:(main) ✗ go vet
// WaitGroup.Add called from inside new goroutine
func Bad() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		//wg.Add(1) // WaitGroup.Add called from inside new goroutine
		defer wg.Done()
		//do work
	}()
	wg.Wait()
}
