package main

import (
	"bufio"
	"fmt"
	"os"
)

// isMarker returns true if the text of 4 characters provided is a marker.
func isMarker(text string) bool {

	// Simply check that all letters are different
	return (text[0] != text[1]) &&
		(text[0] != text[2]) &&
		(text[0] != text[3]) &&
		(text[1] != text[2]) &&
		(text[1] != text[3]) &&
		(text[2] != text[3])
}

// isMessage returns true if the text of 14 characters provided is a message.
func isMessage(text string) bool {
	characters := make(map[rune]struct{}, 0)

	for _, c := range text {
		if _, ok := characters[c]; ok == false {
			characters[c] = struct{}{}
		} else {
			// character exists
			return false
		}
	}
	return true
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	// Read the entire file since it has no newlines.
	scanner.Scan()

	text := scanner.Text()
	position := 0

	for i := 0; i < len(text)-4; i++ {
		if isMarker(text[i : i+4]) {
			position = i + 4
			break
		}
	}

	fmt.Printf("The marker is at position %d\n", position)

	for i := 0; i < len(text)-14; i++ {
		if isMessage(text[i : i+14]) {
			position = i + 14
			break
		}
	}

	fmt.Printf("The message is at position %d\n", position)

}
