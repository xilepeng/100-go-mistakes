package main

import "fmt"

func listing1() {
	a := [3]int{0, 1, 2}
	for i, v := range a {
		a[2] = 10
		if i == 2 {
			fmt.Println("打印数组拷贝，循环不会更新数组拷贝", v)//2
		}
	}
}

func listing2() {
	a := [3]int{0, 1, 2}
	for i := range a {
		a[2] = 10
		if i == 2 {
			fmt.Println("使用索引访问原始数组", a[2]) //10
		}
	}
}

func listing3() {
	a := [3]int{0, 1, 2}
	for i, v := range &a {
		a[2] = 10
		if i == 2 {
			fmt.Println("使用数组指针访问原始数组", v) //10
		}
	}
}
func main() {
	listing1()
	listing2()
	listing3()
}
