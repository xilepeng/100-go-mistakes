package main

import "testing"

var global string

func BenchmarkConcatV1(b *testing.B) {
	var local string
	s := getInput()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = concat1(s)
	}
	global = local
}

func BenchmarkConcatV2(b *testing.B) {
	var local string
	s := getInput()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = concat2(s)
	}
	global = local
}

func BenchmarkConcatV3(b *testing.B) {
	var local string
	s := getInput()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = concat3(s)
	}
	global = local
}

func getInput() []string {
	n := 1_000
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = string(make([]byte, 1_000))
	}
	return s
}

/*
➜  39-string-concat git:(main) ✗ go test -bench=.
BenchmarkConcatV1-10                  31          36779341 ns/op
BenchmarkConcatV2-10                2841            408536 ns/op
BenchmarkConcatV3-10               16041             73615 ns/op
PASS
ok      github.com/xilepeng/100-go-mistakes/05-strings/39-string-concat     6.337s

*/
