package main

import "fmt"

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}
	current := len(s) - 1
	if s[current] == ' ' {
		for s[current] == ' ' && current >= 0 {
			current--
		}
	}
	// This means the string only has whitespaces.
	if current == -1 {
		return 0
	}

	// Starting from here we can talk about non-whitespace characters
	end := current

	// Lazy evaluation ensures that this doesn't try to index -1.
	for current >= 0 && s[current] != ' ' {
		current--
	}

	return end - current
}

func main() {
	fmt.Println(lengthOfLastWord("a"))
}
