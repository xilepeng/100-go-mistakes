package main

import (
	"strings"
)

// 连接5个以内的字符串，使用 += 运算符代码可读性更高
func concat1(values []string) string {
	s := ""
	for _, value := range values {
		s += value
	}
	return s
}

func concat2(values []string) string {
	sb := strings.Builder{}
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}

// 从性能方面考虑，连接5个以上的字符串时，使用strings.Builder 方案会更快。
// 如果预先知道将来字符串的字节数，应该使用 Grow 方法预先分配内部字节切片
func concat3(values []string) string {
	total := 0
	for i := 0; i < len(values); i++ { // 迭代字符串切片以计算总字节数
		total += len(values[i])
	}
	// total= 5
	sb := strings.Builder{}
	sb.Grow(total)
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}

// v3 版本是迄今为止最高效的：比 v1 快 99%，比 v2 快 78%

func main() {
	s := []string{"hello"} // len(s)= 1

	concat1(s)
	concat2(s)
	concat3(s)
}
