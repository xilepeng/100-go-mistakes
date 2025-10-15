package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	s1 := "Hello, World!"
	s2 := s1[:5] // 前5个字节创建字符串
	fmt.Println(s2)

	s1 = "Hêllo, World!"
	s2 = string([]rune(s1)[:5]) // 前5个rune创建字符串
	fmt.Println(s2)
}

type store struct{}

func (s store) handleLog1(log string) error {
	if len(log) < 36 {
		return errors.New("log is not correctly formatted")
	}
	uuid := log[:36] // 内存泄露
	s.store(uuid)
	// Do something
	return nil
}

// 因为字符串主要是一个指针，所以调用函数传递字符串不会导致字节的深度复制。
// 复制的字符串仍将引用相同的底层数组。-> 导致内存泄露
func (s store) handleLog2(log string) error {
	if len(log) < 36 {
		return errors.New("log is not correctly formatted")
	}
	uuid := string([]byte(log[:36])) // 深度复制，使内部的字节切片引用新的底层数组，防止内存泄露。
	s.store(uuid)
	// Do something
	return nil
}

// strings.Clone 可返回字符串的新副本
func (s store) handleLog3(log string) error {
	if len(log) < 36 {
		return errors.New("log is not correctly formatted")
	}
	uuid := string(strings.Clone(log[:36])) // 复制到新分配中，以防止内存泄露
	s.store(uuid)
	// Do something
	return nil
}

func (s store) store(uuid string) {
	// ...
}
