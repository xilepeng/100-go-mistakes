package main

import (
	"fmt"
)

func value_copy() {
	s := []int{0, 1, 2}
	for _, s_copy := range s {
		s_copy += 10
	}
	fmt.Println("原始切片 s =", s) //  s= [0 1 2]
}

func update_slice() {
	s := []int{0, 1, 2}
	for i := range s {
		s[i] += 10
	}
	fmt.Println("使用索引变量访问切片中的元素 s =", s) // s = [10 11 12]

	for i := 0; i < len(s); i++ {
		s[i] += 10
	}
	fmt.Println("使用传统for访问切片 s =", s) // s = [20 21 22]
}

type accounts struct {
	balance float32
}

func pointer_slice() {
	a := []*accounts{{balance: 0}, {balance: 1}, {balance: 2}}
	for _, a := range a {
		a.balance += 10
		fmt.Println("更新后的切片元素 a.balance =", a.balance)
	}
}

func main() {
	value_copy()
	update_slice()
	pointer_slice()
}

/*
原始切片 s = [0 1 2]
使用索引变量访问切片中的元素 s = [10 11 12]
使用传统for访问切片 s = [20 21 22]
更新后的切片 a.balance = 10
更新后的切片 a.balance = 11
更新后的切片 a.balance = 12
*/
