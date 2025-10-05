package main

import "fmt"

func no_update_copy() {
	a := [3]int{0, 0, 0}
	a_copy := []int{}
	for i, v := range a {
		a[i] = 1
		a_copy = append(a_copy, v)
	}
	fmt.Println("a_copy =", a_copy) // a_copy = [0 0 0]
	fmt.Println("a =", a)           // a = [1 1 1]
}

func main() {
	no_update_copy()
}
