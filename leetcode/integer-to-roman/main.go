package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	var d1, d2, d3, d4 int
	d4 = num % 10
	num = num / 10

	d3 = num % 10
	num = num / 10

	d2 = num % 10
	num = num / 10

	d1 = num % 10
	num = num / 10

	out := &strings.Builder{}
	if d1 > 0 {
		for range d1 {
			out.WriteByte('M')
		}
	}

	if d2 >= 1 && d2 <= 3 {
		for range d2 {
			out.WriteByte('C')
		}
	}
	if d2 == 4 {
		out.WriteByte('C')
		out.WriteByte('D')
	}
	if d2 >= 5 && d2 <= 8 {
		out.WriteByte('D')
		for i := 5; i < d2; i++ {
			out.WriteByte('C')
		}
	}
	if d2 == 9 {
		out.WriteByte('C')
		out.WriteByte('M')
	}

	if d3 >= 1 && d3 <= 3 {
		for range d3 {
			out.WriteByte('X')
		}
	}
	if d3 == 4 {
		out.WriteByte('X')
		out.WriteByte('L')
	}
	if d3 >= 5 && d3 <= 8 {
		out.WriteByte('L')
		for i := 5; i < d3; i++ {
			out.WriteByte('X')
		}
	}
	if d3 == 9 {
		out.WriteByte('X')
		out.WriteByte('C')
	}

	if d4 >= 1 && d4 <= 3 {
		for range d4 {
			out.WriteByte('I')
		}
	}
	if d4 == 4 {
		out.WriteByte('I')
		out.WriteByte('V')
	}
	if d4 >= 5 && d4 <= 8 {
		out.WriteByte('V')
		for i := 5; i < d4; i++ {
			out.WriteByte('I')
		}
	}
	if d4 == 9 {
		out.WriteByte('I')
		out.WriteByte('X')
	}

	return out.String()
}

func main() {
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(8))
	fmt.Println(intToRoman(19))
	fmt.Println(intToRoman(60))
}
