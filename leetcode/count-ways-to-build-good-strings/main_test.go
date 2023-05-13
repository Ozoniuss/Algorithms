package main

import "testing"

var total int

func BenchmarkBfsNormal(b *testing.B) {
	k := 0
	for n := 0; n < b.N; n++ {
		k = countGoodStringsBfsNormal(100, 100, 10, 1)
	}
	total = k
}
func BenchmarkDfsMemo(b *testing.B) {
	k := 0
	for n := 0; n < b.N; n++ {
		k = countGoodStringsDfs(100, 100, 10, 1)
	}
	total = k
}
func BenchmarkDp(b *testing.B) {
	k := 0
	for n := 0; n < b.N; n++ {
		k = countGoodStringsDp(100, 100, 10, 1)
	}
	total = k
}
