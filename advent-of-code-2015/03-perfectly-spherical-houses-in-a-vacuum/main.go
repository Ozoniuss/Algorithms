package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Ozoniuss/Algorithms/aocommon/utils/twod"
)

// dirFromString takes a string representing a direction as input, and returns
// the corresponding direction.
func dirFromString(dirstr byte) twod.Direction {
	switch dirstr {
	case '<':
		return twod.W
	case '>':
		return twod.E
	case '^':
		return twod.N
	case 'v':
		return twod.S
	default:
		return twod.Direction{}
	}
}

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
	if part == 1 {

		visited := make(map[twod.Location]struct{})
		current := twod.ORIGIN
		visited[current] = struct{}{}

		// Read the input.
		scanner.Scan()
		line := scanner.Text()
		for i := 0; i < len(line); i++ {

			// Find the direction and move in that direction.
			dir := dirFromString(line[i])
			current = twod.Move(current, dir)

			// If the current direction was not visited, mark it as visited.
			if _, ok := visited[current]; !ok {
				visited[current] = struct{}{}
			}
		}

		fmt.Printf("Santa visited %d locations", len(visited))
	} else if part == 2 {
		visited := make(map[twod.Location]struct{})
		currentSanta := twod.ORIGIN
		currentRobo := twod.ORIGIN
		visited[twod.ORIGIN] = struct{}{}

		// Read the input.
		scanner.Scan()
		line := scanner.Text()
		for i := 0; i < len(line); i++ {

			// Find the direction and move in that direction.
			dir := dirFromString(line[i])

			var current twod.Location
			if i%2 == 0 {
				current = twod.Move(currentSanta, dir)
				currentSanta = current
			} else {
				current = twod.Move(currentRobo, dir)
				currentRobo = current
			}
			// If the current direction was not visited, mark it as visited.
			if _, ok := visited[current]; !ok {
				visited[current] = struct{}{}
			}
		}
		fmt.Printf("Santa and RoboSanta visited %d locations", len(visited))
	}
}
