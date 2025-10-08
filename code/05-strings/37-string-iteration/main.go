package main

import "fmt"

func main() {
	s := "hêllo"
	for i := range s { // 迭代 rune 的每个起始索引，不会迭代每个 rune
		fmt.Printf("position %d: %c\n", i, s[i])
	}
	fmt.Printf("len=%d\n", len(s))

	for i, r := range s { // 返回 rune 的起始索引和 rune 本身
		fmt.Printf("position %d: %c\n", i, r)
	}

	runes := []rune(s)        //·将字符串转换为 rune 切片
	for i, r := range runes { // 返回 rune 索引和 rune 本身
		fmt.Printf("position %d: %c\n", i, r)
	}

	s2 := "hello"
	fmt.Printf("%c\n", rune(s2[4]))
}

func getIthRune(largeString string, i int) rune {
	for idx, v := range largeString {
		if idx == i {
			return v
		}
	}
	return -1
}
