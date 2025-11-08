package main

import (
	"fmt"
)

func listing0() {
	for i := 0; i < 3; i++ {
		fmt.Printf("&i=%p\n", &i)
	}
}

// go 1.22 后，i 的地址是变化的
// i0 i1 i2,超出作用域自动回收
//&i=0x14000182020
//&i=0x14000182040
//&i=0x14000182048

func listing1() {
	s := []int{0, 1, 2}
	for i, val_copy := range s {
		fmt.Printf("i=%p,val_copy=%p\n", &i, &val_copy)
	}
	fmt.Println("原始切片 s =", s) //  s= [0 1 2]
}

//i=0x1400010a028,val_copy=0x1400010a020
//i=0x1400010a058,val_copy=0x1400010a050
//i=0x1400010a068,val_copy=0x1400010a060

func listing2() {
	s := []int{1, 2, 3}
	for index := range s {
		fmt.Printf("index=%d ", index)
	}
}

//index=0 index=1 index=2

func value_copy() {
	s := []int{0, 1, 2}
	for _, s_copy := range s {
		s_copy += 10
	}
	fmt.Println("原始切片 s =", s)
}

//  s= [0 1 2]

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

// 使用索引变量访问切片中的元素 s = [10 11 12]
// 使用传统for访问切片 s = [20 21 22]
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

//更新后的切片 a.balance = 10
//更新后的切片 a.balance = 11
//更新后的切片 a.balance = 12

func main() {
	listing0()
	//listing1()
	//listing2()
	//value_copy()
	//update_slice()
	//pointer_slice()
}
