package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		b, a = a, b
	}

	out := make([]byte, len(a)+1, len(a)+1)
	diff := len(a) - len(b)

	var cr byte = 0
	var x byte

	// Note that character '0' is ASCII 48, and character '1' is ASCII 49.
	for i := len(a) - 1; i >= 0; i-- {
		if i-diff >= 0 {
			x = (a[i] - 48 + b[i-diff] - 48 + cr)
		} else {
			x = (a[i] - 48 + cr)
		}
		out[i+1] = x%2 + 48
		cr = x / 2
	}
	if cr == 1 {
		out[0] = '1'
		return string(out)
	} else {
		return string(out[1:])
	}
}

func main() {
	fmt.Println(addBinary("10", "110")) // 1000
}
