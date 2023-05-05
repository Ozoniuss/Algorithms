package main

import (
	"fmt"
	"os"
)

func readString() []byte {
	f, _ := os.Open("longstring.txt")
	defer f.Close()
	buf := make([]byte, 50000, 50000)
	f.Read(buf)
	return buf
}

//go:noinline
func reverseString(s []byte) {
	var aux byte
	var length = len(s)
	for i := 0; i < length/2; i++ {
		aux = s[i]
		s[i] = s[length-1-i]
		s[length-1-i] = aux
	}
}

// No auxiliary variable.
//
//go:noinline
func reverseStringInPlace(s []byte) {
	var length = len(s)
	for i := 0; i < length/2; i++ {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
}

func main() {
	a := readString()
	reverseString(a)
	fmt.Println(string(a))
}
