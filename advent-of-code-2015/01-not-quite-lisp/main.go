package main

import (
	"bufio"
	"fmt"
	"os"
)

// determineFloor determines the floor based on a parentheses input.
func determineFloor(input string) int {
	// Starting floor
	floor := 0

	for _, p := range input {
		if p == '(' {
			floor++
		} else {
			floor--
		}
	}
	return floor
}

// determineBasementEntrance determines the position in the input where the
// floor changes to -1.
func determineBasementEntrance(input string) int {
	floor := 0
	for idx, p := range input {
		if p == '(' {
			floor++
		} else {
			floor--
		}
		if floor == -1 {
			return idx + 1
		}
	}
	return -1
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	// There is only one line in the input.
	scanner.Scan()
	line := scanner.Text()

	fmt.Printf("The instructions take santa to floor %d\n", determineFloor(line))
	fmt.Printf("Santa enters the basement at position %d", determineBasementEntrance(line))
}
