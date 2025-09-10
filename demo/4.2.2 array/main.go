package main

import "fmt"

func main() {
	a := [3]int{0, 0, 0}
	for i, v := range a {
		a[0] = 1
		if i == 0 {
			fmt.Printf("a_copy[0]=%d \n", v)
		}
	}
}
