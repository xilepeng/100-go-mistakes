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

//go:noinline
func sumValue(x, y int) int {
	// Go 运行时创建一个新的栈帧作为当前 goroutine 栈的一部分。
	// x、y、z 在当前栈帧中一起被分配
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
func sumPtr(x, y int) *int { // 返回一个指针
	z := x + y
	return &z //如果函数返回后某个变量被引用，则会在堆上分配该变量。
}


func listing3() {
	a := 3
	b := 2
	c := sum(&a, &b) // 向下分享留在栈上
	println(c)
}

//go:noinline
func sum(x, y *int) int {
	return *x + *y
}

func main() {
	listing1()
	listing2()
	listing3()
}

/*

➜  95-stack-heap git:(main) ✗ go build -gcflags "-m=2" main.go
# command-line-arguments
./main.go:15:6: cannot inline sumValue: marked go:noinline
./main.go:4:6: can inline listing1 with cost 77 as: func() { a := 3; b := 2; c := sumValue(a, b); println(c) }
./main.go:32:6: cannot inline sumPtr: marked go:noinline
./main.go:23:6: can inline listing2 with cost 78 as: func() { a := 3; b := 2; c := sumPtr(a, b); println(*c) }
./main.go:41:6: can inline main with cost 79 as: func() { listing1() }
./main.go:42:10: inlining call to listing1
./main.go:33:2: z escapes to heap:
./main.go:33:2:   flow: ~r0 = &z:
./main.go:33:2:     from &z (address-of) at ./main.go:34:9
./main.go:33:2:     from return &z (return) at ./main.go:34:2
./main.go:33:2: moved to heap: z
*/
