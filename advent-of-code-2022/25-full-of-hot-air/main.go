package main

import (
	"bufio"
	"fmt"
	"os"
)

var ShafuValues = map[byte]int{
	'2': 2,
	'1': 1,
	'0': 0,
	'-': -1,
	'=': -2,
}

var ShafuStrings = map[int]byte{
	2:  '2',
	1:  '1',
	0:  '0',
	-1: '-',
	-2: '=',
}

// div returns the quotient and remainder when dividing a number by 5, with the
// condition that the remainder must be between -2 and 2.
//
// The quotient remainder theorem must still be satisfied:
// dividend = quotient * 5 + remainder
//
// Using this new div operation, one can follow the exact same method of
// converting to base 5 to convert a number to Shafu, and remainers of -2 and -1
// will be represented with '=' and '-', respectively. This operation will take
// care automatically of adding one digit when the leading digit is "negative".
//
// For example, if converting 3, the operation returns 1 as the quotient and
// -2 (=) as the remainder, so in the convertion process this will be converted
// to '1='.
func div(dividend int) (quotient int, remainder int) {
	// Cool math trick
	quotient = (dividend + 2) / 5
	remainder = (dividend+2)%5 - 2
	return
}

// decValueToShafu takes a decimal value and returns its SHAFU representation.
func decValueToShafu(number int) string {
	shafu := ""
	for number > 2 {
		quotient, remainder := div(number)
		shafu = fmt.Sprintf("%c", ShafuStrings[remainder]) + shafu
		number = quotient
	}
	// Always positive for now.
	shafu = fmt.Sprint(number) + shafu
	return shafu
}

// shafuToDec takes a SHAFU number and returns its decimal value.
func shafuToDecValue(number string) int {
	value := 0
	multiplier := 1
	for i := len(number) - 1; i >= 0; i-- {
		value += ShafuValues[number[i]] * multiplier
		multiplier *= 5
	}
	return value
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	total := 0

	for scanner.Scan() {
		number := scanner.Text()
		total += shafuToDecValue(number)
	}

	fmt.Printf("Total is %d\n", total)
	fmt.Printf("Total in shafu is %s\n", decValueToShafu(total))
}
