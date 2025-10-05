package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Init
	n := 1_000_000
	m := make(map[int]*[128]byte)
	printAlloc()

	// Add elements
	for i := 0; i < n; i++ {
		m[i] = randBytes()
	}
	printAlloc()

	// Remove elements
	for i := 0; i < n; i++ {
		delete(m, i)
	}

	// End
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(m)
}

func randBytes() *[128]byte {
	return &[128]byte{}
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}


/*
m := make(map[int][128]byte)

0 MB    // 分配 m 后
317 MB  // 添加 100 万个元素后
288 MB  // 删除 100 万个元素后


m := make(map[int]*[128]byte)

0 MB
170 MB
36 MB
*/