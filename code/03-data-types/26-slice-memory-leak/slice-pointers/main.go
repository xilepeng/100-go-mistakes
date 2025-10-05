package main

import (
	"fmt"
	"runtime"
)

type Foo struct {
	v []byte
}

func main() {
	foos := make([]Foo, 1_000)
	printAlloc() // 128 KB

	for i := 0; i < len(foos); i++ { // 为每个元素分配 1 MB 字节
		foos[i] = Foo{
			v: make([]byte, 1024*1024),
		}
	}
	printAlloc() // 1024141 KB

	two := keepFirstTwoElementsOnly(foos)
	runtime.GC() // 运行GC以强制清理堆
	printAlloc() // 1024141 KB
	runtime.KeepAlive(two)
}

// 泄露切片剩余元素
func keepFirstTwoElementsOnly(foos []Foo) []Foo {
	return foos[:2]
}

// 创建切片副本可确保不泄露切片剩余元素
func keepFirstTwoElementsOnlyCopy(foos []Foo) []Foo {
	res := make([]Foo, 2)
	copy(res, foos)
	return res
}

// 如果想保持 1000 个元素的容量，那就将切片的剩余元素显式标记为 nil
func keepFirstTwoElementsOnlyMarkNil(foos []Foo) []Foo {
	for i := 2; i < len(foos); i++ {
		foos[i].v = nil
	}
	return foos[:2]
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
