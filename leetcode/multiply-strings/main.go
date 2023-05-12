package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	var out = make([]byte, len(num1)+len(num2), len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		d1 := num1[i] - '0'
		for j := len(num2) - 1; j >= 0; j-- {
			d2 := num2[j] - '0'

			// Based on carry too.
			out[i+j+1] += d1 * d2

			// Take the result mod 10, and add the carry to the next
			// digit.
			out[i+j] += out[i+j+1] / 10
			// Always take the result mod 10
			out[i+j+1] = out[i+j+1] % 10
		}
	}

	// Note that at most 1 digit can be 0, place multiplication between
	// x 10..0 and x 10..0 * 10
	if out[0] == 0 {
		out = out[1:]
	}

	for i := range out {
		out[i] += '0'
	}

	return string(out)

}

func main() {
	fmt.Println(multiply("123", "456"))
}
