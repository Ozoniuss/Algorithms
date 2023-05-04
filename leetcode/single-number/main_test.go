package main

import "testing"

var outnum = 0

func BenchmarkSingleNum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		outnum = singleNumber(allnums)
	}
}
func BenchmarkSingleNumInt32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		outnum = singleNumberInt32(allnums)
	}
}
