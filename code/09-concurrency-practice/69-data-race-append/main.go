package main

import "fmt"

// 没有数据竞争
func listing1() {
	s := make([]int, 1)

	go func() {
		s1 := append(s, 1)
		fmt.Println(s1)
	}()

	go func() {
		s2 := append(s, 1)
		fmt.Println(s2)
	}()
}

// 没有数据竞争
// 制作切片副本，在副本上使用append,防止数据竞争，
// 两个 goroutine 都在隔离数据上工作
func listing2() {
	s := make([]int, 0, 1)

	go func() {
		sCopy := make([]int, len(s), cap(s))
		copy(sCopy, s)

		s1 := append(sCopy, 1)
		fmt.Println(s1)
	}()

	go func() {
		sCopy := make([]int, len(s), cap(s))
		copy(sCopy, s)

		s2 := append(sCopy, 1)
		fmt.Println(s2)
	}()
}

// 为什么没检测到数据竞争？
func listingRace() {
	s := make([]int, 0, 1)
	go func() {
		s1 := append(s, 1)
		fmt.Println(s1)
	}()

	go func() {
		s2 := append(s, 1)
		fmt.Println(s2)
	}()
}

func main() {
	//listing1()
	//listing2()
	listingRace()
}

//go run main.go -race
