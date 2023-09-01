package main

import (
	"fmt"
)

// Speedran solution (note this fails for utf-8 encoded strings in which there)
// are characters encoded with more than 1 byte.
func reverseString(s string) string {
	if len(s) == 1 {
		return s
	}
	return string(s[len(s)-1]) + reverseString(s[:len(s)-1])
}

func main() {
	fmt.Println("gnirtsym" == reverseString("mystring"))
	fmt.Println("gnirtsÔ" == reverseString("Ôstring"))
}
