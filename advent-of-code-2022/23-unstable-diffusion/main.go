package main

import (
	"bufio"
	"fmt"
	"os"
)

type Location [2]int
type Direction [2]int

var N Direction = Direction{-1, 0}
var S Direction = Direction{1, 0}
var W Direction = Direction{0, -1}
var E Direction = Direction{0, 1}
var NE Direction = Direction{-1, 1}
var NW Direction = Direction{-1, -1}
var SE Direction = Direction{1, 1}
var SW Direction = Direction{1, -1}

// add takes an initial location and a direction, and returns the new location
// after moving in that direction.
func add(d1 Location, d2 Direction) Location {
	return Location{d1[0] + d2[0], d1[1] + d2[1]}
}

// processLine reads a line from the input and adds it to the list of elves.
// We don't really care about "." on the map, the map is infinite.
func processLine(line string, row int, elves map[Location]struct{}) {
	for i := 0; i < len(line); i++ {
		if line[i] == '#' {
			elves[Location{row + 1, i + 1}] = struct{}{}
		}
	}
}

// Syntactic sugar
type moveCycle [4][3]Direction

// getMoveCycle returns the order of moves each elf will attempt to make, as
// described in the problem statement. After every round, the first three
// moves at the start of the move cycle will go at the back.
func getMoveCycle() moveCycle {
	return moveCycle{
		{N, NE, NW},
		{S, SE, SW},
		{W, NW, SW},
		{E, NE, SE},
	}
}

// rotateMoveCycle takes the first move from the move cycle and adds it aat
// the end.
func (c *moveCycle) rotateMoveCylce() {
	first := c[0]
	for i := 0; i < 3; i++ {
		c[i] = c[i+1]
	}
	c[3] = first
}

// shouldMove returns whether an elf should move or not, based on its
// neighbourds.
func shouldMove(elf Location, elves map[Location]struct{}) bool {
	allDirections := []Direction{N, NE, E, SE, S, SW, W, NW}
	for _, dir := range allDirections {
		if _, hasNeighbour := elves[add(elf, dir)]; hasNeighbour {
			return true
		}
	}
	return false
}

// move simulates one step, that is, deciding on the new locations and actually
// moving to the new locations. The way this is done is by creating a "next"
// map, in which the keys represent the new locations elves want to go to, and
// each key is associates with a list of old elves (old elf locations) that
// want to move to the new locations. If that list has a single element, we
// can make the move; otherwise all elves from that list must stay in their
// original place.
//
// move returns whether the elves still need to move or not.
func move(elves map[Location]struct{}, moveCylce moveCycle) bool {

	next := make(map[Location][]Location)
	stillNeedsToMove := false

	// Go through all elves and find where they want to move:
	for elf := range elves {

		hasProposal := true

		// If the elf shouldn't move, disregard this step.
		if !shouldMove(elf, elves) {
			continue
		}

		stillNeedsToMove = true

		for _, moves := range moveCylce {

			hasProposal = true

			for _, dir := range moves {
				wantsToGoAt := add(elf, dir)

				// If there is an elf, can't go there and do not make a proposal
				// for this move cycle.
				if _, isElf := elves[wantsToGoAt]; isElf {
					hasProposal = false
					break
				}
			}

			// There is a valid proposal in the current move cycle.
			if hasProposal {
				// The proposed dir is always the first one in the cycle.
				proposedDir := moves[0]
				nextLoc := add(elf, proposedDir)

				if _, ok := next[nextLoc]; !ok {
					// If it's just this elf wanting to go there, create a new
					// entry.
					next[nextLoc] = []Location{elf}
				} else {
					// Otherwise, add it to the list of elves that want to go
					// there.
					next[nextLoc] = append(next[nextLoc], elf)
				}
				break
			}

		}
	}

	// Actually make the moves to the new location
	for newLocation := range next {
		// This is the only case where elves actually do move.
		if len(next[newLocation]) == 1 {
			elf := next[newLocation][0]
			delete(elves, elf)
			elves[newLocation] = struct{}{}
		}
	}

	return stillNeedsToMove

}

// makeMoves makes a move for the given number of rounds, using the move
// function above.
func makeMoves(elves map[Location]struct{}, rounds int) {
	moveCycle := getMoveCycle()
	for r := 0; r < rounds; r++ {
		move(elves, moveCycle)
		moveCycle.rotateMoveCylce()
	}
}

// makeNecessaryMoves moves the elves until it's necessary; that is, until no
// elf has to move anymore.
func makeNecessaryMoves(elves map[Location]struct{}) int {
	moveCycle := getMoveCycle()
	needsToMove := true
	round := 0
	for needsToMove {
		needsToMove = move(elves, moveCycle)
		moveCycle.rotateMoveCylce()
		round++
	}
	return round
}

// findRectangleBounds finds the vertices of the rectangle with the smallest
// area containing all elves. The vertices are representing the rectangle's
// diagonal.
func findRectangleBounds(elves map[Location]struct{}) [2]Location {

	min_row := 1000
	min_col := 1000
	max_row := -1000
	max_col := -1000

	for elf := range elves {
		if elf[0] > max_row {
			max_row = elf[0]
		}
		if elf[0] < min_row {
			min_row = elf[0]
		}
		if elf[1] > max_col {
			max_col = elf[1]
		}
		if elf[1] < min_col {
			min_col = elf[1]
		}
	}

	return [2]Location{
		{min_row, min_col}, {max_row, max_col},
	}
}

// drawRectangle draws the rectangle with smallest area containing the elves.
func drawRectangle(bounds [2]Location, elves map[Location]struct{}) {
	out := ""
	for row := bounds[0][0]; row <= bounds[1][0]; row++ {
		for col := bounds[0][1]; col <= bounds[1][1]; col++ {
			loc := Location{row, col}
			if _, ok := elves[loc]; ok {
				out += "#"
			} else {
				out += "."
			}
		}
		out += "\n"
	}
	fmt.Print(out)
}

func computeEmptyGround(bounds [2]Location, elves map[Location]struct{}) int {
	empty := 0
	for row := bounds[0][0]; row <= bounds[1][0]; row++ {
		for col := bounds[0][1]; col <= bounds[1][1]; col++ {
			loc := Location{row, col}
			if _, ok := elves[loc]; !ok {
				empty += 1
			}
		}
	}
	return empty
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	elves := make(map[Location]struct{})

	scanner := bufio.NewScanner(f)
	row := 0

	for scanner.Scan() {
		processLine(scanner.Text(), row, elves)
		row++
	}

	part := 2
	if part == 1 {
		makeMoves(elves, 10)
		bounds := findRectangleBounds(elves)
		emtpy := computeEmptyGround(bounds, elves)
		//drawRectangle(bounds, elves)
		fmt.Printf("Empty ground: %d", emtpy)
	} else {
		round := makeNecessaryMoves(elves)
		fmt.Printf("No need to move on round: %d", round)
	}

}
