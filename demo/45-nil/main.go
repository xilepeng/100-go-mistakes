package main

import "fmt"

type Foo struct{}

func (foo *Foo) Bar() string {
	return "bar"
}

func main() {
	var foo *Foo
	fmt.Println("foo=",foo==nil) // nil
	fmt.Println("foo.Bar()=",foo.Bar()) // bar
}


