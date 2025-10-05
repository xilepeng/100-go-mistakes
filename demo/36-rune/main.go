package main

import (
	"fmt"
)

func main() {
	s1 := "hello"
	fmt.Println("字节数量: len(s1) = ", len(s1))

	s2 := "徐"
	fmt.Println("字节数量: len(s2) = ", len(s2))
}


// ➜  36 rune git:(main) ✗ go run main.go
// 字节数量: len(s1) =  5
// 字节数量: len(s2) =  3
