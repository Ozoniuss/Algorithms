package main

import "fmt"

/* My actual timed solution */

func medium(a, b, c int) int {
	// The reasoning here is that if the two sides of the multiplication are of
	// opposite signs, it means that the number on the left in each parenthesis
	// is between the two numbers on the right in the parenthesis.
	if (a-b)*(a-c) < 0 {
		return a
	} else if (b-a)*(b-c) < 0 {
		return b
	} else {
		return c
	}
}

/* end */

func main() {
	var a, b, c int

	fmt.Scanf("%d %d %d", &a, &b, &c)
	fmt.Printf("medium of %d, %d, %d: %d", a, b, c, medium(a, b, c))
}
