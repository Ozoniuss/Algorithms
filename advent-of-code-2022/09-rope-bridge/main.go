package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Syntactic sugar.
type Location [2]int

// Direction is intended to represent the values added to each coordinate, if
// the current location is moved in that direction.
type Direction [2]int

// add takes an initial location and a direction, and returns the new location
// after moving in that direction.
func add(d1 Location, d2 Direction) Location {
	return Location{d1[0] + d2[0], d1[1] + d2[1]}
}

// dirFromString converts a direction from string to Direction.
func dirFromString(dir string) Direction {
	if dir == "U" {
		return UP
	} else if dir == "D" {
		return DOWN
	} else if dir == "L" {
		return LEFT
	} else if dir == "R" {
		return RIGHT
	}
	return EMPTY_DIR
}

var (
	UP        = Direction{-1, 0}
	DOWN      = Direction{1, 0}
	LEFT      = Direction{0, -1}
	RIGHT     = Direction{0, 1}
	EMPTY_DIR = Direction{0, 0}
)

// abs returns the absolute value of an integer.
func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

// manhattanDistance returns the manhattan distance between two points.
func manhattanDistance(p1, p2 Location) int {
	return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
}

/* Part 1 */

// move takes the current position of the head and the tail and the direction
// of the move, and will return the new position of the head and tail after the
// move. This function is designed for the two-length rope at part (a).
func move(direction Direction, head, tail Location) (Location, Location) {

	original := head
	head = add(head, direction)

	// If the tail needs to move, it will always go in the head's original
	// location.
	if needsMove(head, tail) {
		tail = original
	}

	return head, tail
}

// needsMove returns wether the tail needs to be moved as well, assuming that
// head was moved to the provided location, and tail had not been moved yet
// from the previous location.
func needsMove(head, tail Location) bool {
	// In this case it's obvious that the tail needs to move.
	if manhattanDistance(head, tail) > 2 {
		return true
	} else if manhattanDistance(tail, head) == 2 {
		// If new head and old tail are located diagonally, no need to move.
		if abs(head[0]-tail[0]) == 1 && abs(head[1]-tail[1]) == 1 {
			return false
		} else {
			// Otherwise, new head and old tail are on a line, and there's at
			// least one empty spot between them.
			return true
		}
	} else {
		return false
	}
}

/* Part 2 */

// moveIfNeedsMove makes a move for a knot in the rope, based on where the
// previous knot went. This location of the previous knot is sufficient in order
// to determine where the knot has to move.
func moveIfNeedsMove(prev, current Location) Location {
	if needsMove(prev, current) {
		// Note that if moving the tail manually, there's some additional
		// details to consider at part 2. At part 1, since diagonal motions are
		// not allowed, the motion described below would not be possible, but
		// it's possible to reach the scenarios below for the longer rope.
		//
		//            . . .           . . P
		//            . P .   --- >   . . .
		//            C . .           C . .
		//
		//
		//            . . .           P . .
		//            . P .   --- >   . . .
		//            C . .           C . .
		//
		// In the first case, that's an additional motion to consider, if
		// moving the tail manually, since it must move diagonally. That is
		// covered by moving the tail to the previous knot's original location,
		// but the second case is not, because it would place the current knot
		// just below the previous one and not in the center of the grid.
		//
		// Thus, I decided to handle these cases manually.
		if manhattanDistance(prev, current) == 4 {
			// Diagonal case illustrated in first drawing.
			current[0] = (prev[0] + current[0]) / 2
			current[1] = (prev[1] + current[1]) / 2
		} else if prev[0]-current[0] == 2 || prev[0]-current[0] == -2 {
			// Move the knot behind the previous one, on the same row.
			current[0] = (prev[0] + current[0]) / 2
			current[1] = prev[1]
		} else if prev[1]-current[1] == 2 || prev[1]-current[1] == -2 {
			// Move the knot behind the previous one, on the same column.
			current[0] = prev[0]
			current[1] = (prev[1] + current[1]) / 2
		}
	}
	// Return the new position of the current knot.
	return current
}

// moveRope moves the entire rope in one direction and returns the new
// positions of the rope.
func moveRope(direction Direction, rope [10]Location) [10]Location {

	// Move the head.
	newRope := [10]Location{}
	newRope[0] = add(rope[0], direction)

	// Move all other knots based on the head.
	for i := 0; i <= 8; i++ {
		newRope[i+1] = moveIfNeedsMove(newRope[i], rope[i+1])
	}

	return newRope
}

// processLine parses the input of the line, returning the direction and the
// number of moves in that direction.
func processLine(line string) (Direction, int) {
	parts := strings.Split(line, " ")
	direction := parts[0]
	noMoves, _ := strconv.ParseInt(parts[1], 10, 32)
	return dirFromString(direction), int(noMoves)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	// The locations visited by the tail
	visited := make(map[[2]int]struct{})
	visited[[2]int{0, 0}] = struct{}{}

	var part int
	fmt.Print("part:")
	fmt.Scanf("%d\n", &part)

	if part == 1 {
		//Initial head and tail positions at 0,0.
		var head, tail Location

		for scanner.Scan() {
			line := scanner.Text()
			direction, noMoves := processLine(line)
			for i := 0; i < noMoves; i++ {
				head, tail = move(direction, head, tail)
				visited[tail] = struct{}{}
			}
		}

		fmt.Printf("The tail travelled through %d places.\n", len(visited))
	}

	if part == 2 {

		// Initialize rope with all positions at 0,0.
		rope := [10]Location{}

		for scanner.Scan() {
			line := scanner.Text()
			direction, noMoves := processLine(line)
			for i := 0; i < noMoves; i++ {
				rope = moveRope(direction, rope)
				visited[rope[len(rope)-1]] = struct{}{}
			}
		}

		fmt.Printf("The tail travelled through %d places.\n", len(visited))
	}

}
