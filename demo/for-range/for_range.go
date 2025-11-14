package main

import (
	"fmt"
	"time"
)

func forIterateAddr() {
	for i := 0; i < 3; i++ {
		fmt.Printf("&i=%p\n", &i)
	}
}

// go 1.22 后，i 的地址是变化的
// i0 i1 i2,超出作用域自动回收

//&i=0x14000182020
//&i=0x14000182040
//&i=0x14000182048

func rangeIterateAddr() {
	s := []int{1, 2, 3}
	for i, v_copy := range s {
		fmt.Printf("&i=%p,&v_copy=%p\n", &i, &v_copy)
	}
}

// &i=0x1400010a068,&v_copy=0x1400010a060
// &i=0x1400010a078,&v_copy=0x1400010a070
// &i=0x1400010a088,&v_copy=0x1400010a080

// 闭包绑定了一个变量，导致变量超出了它的作用域
func listing1() {
	s := []int{1, 2, 3}
	for _, v_copy := range s {
		go func() { // 闭包是一个函数值，它不会捕获 goroutine 创建时的值，而是所有 goroutine 都引用完全相同的变量i，
			fmt.Print(v_copy) // 因此，自 goroutine 启动以来 i 的地址可能已被修改
		}()
	}
}

// 不以特定顺序打印 123（因为不能保证创建的第一个 goroutine 会首先执行完成）
// 312 132 123 213

func listing2() {
	s := []int{1, 2, 3}
	for _, v_copy := range s {
		val := v_copy // 为每次迭代创建一个新变量
		go func() {
			fmt.Print(val)
		}()
	}
}

// 不以特定顺序打印 123（因为不能保证创建的第一个 goroutine 会首先执行完成）
// 312 132 123 213

func listing3() {
	s := []int{1, 2, 3}
	for _, v_copy := range s {
		go func(val int) { // 执行带参数函数
			fmt.Print(val)
		}(v_copy) // 调用函数并传递 i 的当前值
	}
}

// 不以特定顺序打印 123（因为不能保证创建的第一个 goroutine 会首先执行完成）
// 312 132 123 213
func main() {
	forIterateAddr()
	rangeIterateAddr()
	// listing1() // 312 132 123 213
	// listing2() // 312 132 123 213
	// listing3()
	time.Sleep(1 * time.Second)
}
