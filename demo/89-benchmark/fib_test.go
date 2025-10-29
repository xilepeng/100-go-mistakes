package main

import "testing"

var res int

// func BenchmarkFib10(b *testing.B) {
// 	var r int
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		r = Fib(10)
// 	}
// 	res = r
// }

func BenchmarkFibFast10(b *testing.B) {
	var r int
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		r = FibFast(10)
	}
	res = r
}
