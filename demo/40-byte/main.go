package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := []byte{'a', 'b', 'c'}
	s := string(b)
	b[1] = 'x'                   //修改切片元素，字符串不可改变
	fmt.Println("b=", string(b)) //b= axc
	fmt.Println("字符串不变性：s=", s)  // abc
	byte_res := bytes.TrimPrefix(b, []byte{'a'})
	fmt.Println("去掉前缀", string(byte_res)) // xc

	// b= axc
	// 字符串不变性：s= abc
	// 去掉前缀 xc
}
