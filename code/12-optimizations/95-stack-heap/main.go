package main

// 栈
func listing1() {
	// 创建一个栈帧，变量a,b被分配到这个栈帧所在的栈中。
	// 存储的所有变量都是有效的地址，这意味着它们可以被引用和访问。
	a := 3
	b := 2

	c := sumValue(a, b)
	println(c) //强制在堆上分配
}

// 
//go:noinline
func sumValue(x, y int) int {
	z := x + y
	return z
}

// 堆
func listing2() {
	a := 3
	b := 2

	c := sumPtr(a, b)
	println(*c)
}

//go:noinline
func sumPtr(x, y int) *int {
	z := x + y
	return &z
}

func listing3() {
	a := 3
	b := 2
	c := sum(&a, &b)
	println(c)
}

//go:noinline
func sum(x, y *int) int {
	return *x + *y
}
