# Day 9: Rope Bridge

> View the [problem statement](https://adventofcode.com/2022/day/9) and [my original submission](https://github.com/Ozoniuss/Algorithms/commit/01ee0638f0ae8a1e5d1edca251de2083fed9823e) at the provided links.

I found this problem to be very similar to simulating a snake game with slightly different movements. The possible motions at part 1 were very nicely described, and since part 2 had just a longer "snake" I thought this would get done quickly. I did however miss some of the subtleties of part 2.

To represent the rope, I stored the position of each knot in a 2-dimensional array. For a better coding experience, I defined the following syntactic sugar:

```go
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

// rope at part 1
var head, tail Location

// rope at part 2
rope := [10]Location{}

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
```

Next, the idea I went with at part 1 was the following: once the head moves, figure out its position relative to the tail, and move the tail if needed. Since those cases were explained clearly in the problem statement, I found an easy implementation using the Manhattan Distance:

```go
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
```

Once this function was defined, it was just a matter of actually moving the head at every step before finishing the problem.