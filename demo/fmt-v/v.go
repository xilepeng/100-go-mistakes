package main

import "fmt"

type student struct {
	name string
	id   int
}

func main() {
	s := student{"Mojo", 123456}
	fmt.Printf("%%v的格式  = %v\n", s)
	fmt.Printf("%%+v的格式 = %+v\n", s)
	fmt.Printf("%%#v的格式 = %#v\n", s)
}

//%v的格式  = {Mojo 123456}
//%+v的格式 = {name:Mojo id:123456}
//%#v的格式 = main.student{name:"Mojo", id:123456}

//%v 值默认格式
//%+v 字段名称：字段值
//%#v 结构体名称{字段名称：字段值}
