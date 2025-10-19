package main

func main() {
	a := 3
	b := 2
	c := sumPtr(a, b)
	println(*c)
}

//go:noinline
func sumPtr(x, y int) *int { // 返回一个指针
	z := x + y
	return &z // 如果函数返回后某个变量被引用，则会在堆上分配该变量。
}


/*
➜  95-heap-stack git:(main) ✗ go build -gcflags "-m=2" main.go
# command-line-arguments
./main.go:11:6: cannot inline sumPtr: marked go:noinline
./main.go:3:6: can inline main with cost 78 as: func() { a := 3; b := 2; c := sumPtr(a, b); println(*c) }
./main.go:12:2: z escapes to heap:
./main.go:12:2:   flow: ~r0 = &z:
./main.go:12:2:     from &z (address-of) at ./main.go:13:9
./main.go:12:2:     from return &z (return) at ./main.go:13:2
./main.go:12:2: moved to heap: z
*/