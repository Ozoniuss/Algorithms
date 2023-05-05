package main

import (
	"sort"
	"testing"
)

var a []byte

func BenchmarkReverseString(b *testing.B) {
	var str = readString()
	b.ResetTimer()
	for n := 0; n <= b.N; n++ {
		reverseString(str)
		a = str
	}
}

func BenchmarkReverseStringReinitialize(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		b.StopTimer()
		var str = readString()
		b.StartTimer()
		reverseString(str)
		a = str
	}
}

func BenchmarkReverseStringInPlace(b *testing.B) {
	var str = readString()
	b.ResetTimer()
	for n := 0; n <= b.N; n++ {
		reverseStringInPlace(str)
		a = str
	}
}

func BenchmarkSliceStable(b *testing.B) {
	var str = readString()
	b.ResetTimer()
	for n := 0; n <= b.N; n++ {
		sort.SliceStable(str, func(a, b int) bool {
			return a > b
		})
		a = str
	}
}

// func BenchmarkReverseStringReinitialize(b *testing.B) {
// 	for n := 0; n <= b.N; n++ {
// 		b.StopTimer()
// 		var reverse = []byte("abcjsbfaskjfbdlfneslfbnesilfnielsflisenleilnfelinrd")
// 		b.StartTimer()
// 		reverseString(reverse)
// 	}
// }
