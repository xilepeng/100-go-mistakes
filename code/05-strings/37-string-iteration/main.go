package main

import "fmt"

func main() {
	s := "hêllo"       // 5 个rune
	for i := range s { // 迭代 rune 的每个起始索引，不会迭代每个 rune
		fmt.Printf("rune 的每个起始索引 %d: %c\n", i, s[i])
	}
	fmt.Printf("len 返回字符串中的字节数，而不是 rune 数,ê不是单个字节编码的，它需要2个字节，所以len=%d\n", len(s))

	for i, r := range s { // 返回 rune 的起始索引和 rune 本身
		fmt.Printf("rune 的起始索引和 rune 本身 %d: %c\n", i, r)
	}

	// 如果想访问字符串的第i个 rune
	runes := []rune(s)        //·将字符串转换为 rune 切片
	for i, r := range runes { // 迭代 rune 索引和 rune 本身
		fmt.Printf("迭代 rune 索引和 rune 本身 %d: %c\n", i, r)
	}

	s2 := "hello"
	fmt.Printf("%c\n", rune(s2[4])) // 0

// rune 的每个起始索引 0: h
// rune 的每个起始索引 1: Ã
// rune 的每个起始索引 3: l
// rune 的每个起始索引 4: l
// rune 的每个起始索引 5: o
// len 返回字符串中的字节数，而不是 rune 数,ê不是单个字节编码的，它需要2个字节，所以len=6

// rune 的起始索引和 rune 本身 0: h
// rune 的起始索引和 rune 本身 1: ê
// rune 的起始索引和 rune 本身 3: l
// rune 的起始索引和 rune 本身 4: l
// rune 的起始索引和 rune 本身 5: o

// 迭代 rune 索引和 rune 本身 0: h
// 迭代 rune 索引和 rune 本身 1: ê
// 迭代 rune 索引和 rune 本身 2: l
// 迭代 rune 索引和 rune 本身 3: l
// 迭代 rune 索引和 rune 本身 4: o
// o

}

func getIthRune(largeString string, i int) rune {
	for idx, v := range largeString {
		if idx == i {
			return v
		}
	}
	return -1
}
