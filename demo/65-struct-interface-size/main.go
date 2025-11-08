package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s struct{}
	var i interface{}
	fmt.Println(unsafe.Sizeof(s)) // 0
	fmt.Println(unsafe.Sizeof(i)) // 16   32位架构占用8字节
}
