package main

import (
	"bufio"
	"fmt"
	"os"
)

// readPart reads the part number from the standard input.
func readPart() int {
	var part int
	fmt.Print("part: ")
	fmt.Scanf("%d\n", &part)
	return part
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	part := readPart()
}
