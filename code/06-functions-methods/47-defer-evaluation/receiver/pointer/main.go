package main

import "fmt"

func main() {
	s := &Struct{id: "foo"}
	defer s.print() // s会被立即计算
	s.id = "bar"    // 更新s.id（可见）
}

type Struct struct {
	id string
}

func (s *Struct) print() {
	fmt.Println(s.id)// 结果是 bar
}
