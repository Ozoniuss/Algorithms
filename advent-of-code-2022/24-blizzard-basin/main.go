package main

import (
	"bufio"
	"fmt"
	"os"
)

type Location [2]int
type Direction [2]int

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDistance(l1, l2 Location) int {
	return abs(l1[0]-l2[0]) + abs(l1[1]-l2[1])
}

// Each blizzard has a location and a direction. For example, ">" has the
// "right" direction.
type Blizzard struct {
	loc Location
	dir Direction
}

var (
	LEFT  Direction = [2]int{0, -1}
	RIGHT Direction = [2]int{0, +1}
	UP    Direction = [2]int{-1, 0}
	DOWN  Direction = [2]int{1, 0}
)

// add takes an initial location and a direction, and returns the new location
// after moving in that direction.
func add(d1 Location, d2 Direction) Location {
	return Location{d1[0] + d2[0], d1[1] + d2[1]}
}

// addWithBorder takes an initial location and a direction, and returns the new
// location after moving in that direction. If the new location is on the
// borders, it returns the new location accross the map, simulating the wrap
// around the map.
func addWithBorder(d1 Location, d2 Direction, borders [2]Location) Location {
	newLocation := add(d1, d2)

	upperLeft := borders[0]
	lowerRight := borders[1]

	// Reached the top border row, going up
	if newLocation[0] == upperLeft[0] {
		// Move it just above the bottommost row
		newLocation[0] = lowerRight[0] - 1
	} else
	// Reached the bottom border row, going down
	if newLocation[0] == lowerRight[0] {
		// Move it just below the topmost row
		newLocation[0] = upperLeft[0] + 1
	} else
	// Reached the left border column, going left
	if newLocation[1] == upperLeft[1] {
		// Move it just right of the rightmost column
		newLocation[1] = lowerRight[1] - 1
	} else
	// Reached the right border column, going right
	if newLocation[1] == lowerRight[1] {
		// Move it just left of the leftmost column
		newLocation[1] = upperLeft[1] + 1
	}
	return newLocation
}

// isOnBorder returns whether the given location is on the border, excluding the
// starting and ending locations.
func isOnBorder(loc, start, end Location, border [2]Location) bool {
	if loc == end {
		return false
	}

	// Note that less/greater or equal are used to cover edge cases.

	// Is on row border
	if loc[0] <= border[0][0] || loc[0] >= border[1][0] {
		return true
	}

	// Is on column border
	if loc[1] <= border[0][1] || loc[1] >= border[1][1] {
		return true
	}

	return false
}

// processLine reads a line from the input and adds the blizzards from that
// line.
func processLine(line string, row int, blizzards *[]Blizzard) {
	for col, c := range line {
		if c == '^' {
			*blizzards = append(*blizzards, Blizzard{loc: Location{row + 1, col + 1}, dir: UP})
		} else if c == '<' {
			*blizzards = append(*blizzards, Blizzard{loc: Location{row + 1, col + 1}, dir: LEFT})
		} else if c == '>' {
			*blizzards = append(*blizzards, Blizzard{loc: Location{row + 1, col + 1}, dir: RIGHT})
		} else if c == 'v' {
			*blizzards = append(*blizzards, Blizzard{loc: Location{row + 1, col + 1}, dir: DOWN})
		}
	}
}

// getBlizzardsMap returns a map with the pozitions of all blizzards, recording
// every position at most once.
func getBlizzardsMap(blizzards []Blizzard) map[Location]struct{} {
	blizzardsMap := make(map[Location]struct{})
	for _, b := range blizzards {
		blizzardsMap[b.loc] = struct{}{}
	}
	return blizzardsMap
}

// moveBlizzards goes through the list of blizzards and moves them to their new
// location. It returns a new Blizzard vector with the new positions.
//
// Note that there are no blizzards facing up or down on the columns with the
// starting or ending positions. This was determined by examining the input and
// the example. I didn't find it mentioned in the problem statement, but at
// least avoids the confusion of moving in that case.
func moveBlizzards(blizzards []Blizzard, borders [2]Location) []Blizzard {
	nextBlizzards := make([]Blizzard, 0, len(blizzards))
	for _, b := range blizzards {
		nextb := Blizzard{dir: b.dir, loc: addWithBorder(b.loc, b.dir, borders)}
		nextBlizzards = append(nextBlizzards, nextb)
	}
	return nextBlizzards
}

// startAndEndLocations returns the starting and ending location, based on the
// borders.
func startAndEndLocations(borders [2]Location) (start Location, end Location) {
	start = Location{borders[0][0], borders[0][1] + 1}
	end = Location{borders[1][0], borders[1][1] - 1}
	return
}

// getAvailableLocations returns the available locations to move from, given
// a map of blizzards and the borders. It's possible to stay in the same
// location.
func getAvailableLocations(current, start, end Location, blizzardMap map[Location]struct{}, borders [2]Location) []Location {

	available := []Location{}
	// You have to move if there's a blizzard coming to your position.
	// Otherwise, can stay.
	if _, ok := blizzardMap[current]; !ok {
		available = append(available, current)
	}

	up := add(current, UP)
	down := add(current, DOWN)
	left := add(current, LEFT)
	right := add(current, RIGHT)

	possibilities := []Location{up, down, left, right}

	for _, pos := range possibilities {
		// We only care about locations that are not covered by blizzards.
		if _, ok := blizzardMap[pos]; !ok {
			if !isOnBorder(pos, start, end, borders) {
				available = append(available, pos)
			}
		}
	}
	return available
}

type state struct {
	location  Location
	blizzards []Blizzard
	step      int
}

// I don't care if I reached the same location through different paths. I only
// have to explore it once.
type hash struct {
	location Location
	step     int
}

// bfs generates all possible moves, stopping when you reached the goal. It
// returns the blizzard positions for part 2, so that the algorithm can be
// reexecuted by starting from the other places.
func bfs(start, end Location, blizzards []Blizzard, borders [2]Location) (int, []Blizzard) {
	initialState := state{
		location:  start,
		blizzards: blizzards,
		step:      0,
	}
	hashes := make(map[hash]struct{})

	q := []state{initialState}
	hashes[hash{
		location: start,
		step:     0,
	}] = struct{}{}

	for {
		if len(q) == 0 {
			panic("something bad happened")
		}
		current := q[0]
		q = q[1:]

		// Remove the current hash
		currentHash := hash{
			location: current.location,
			step:     current.step,
		}
		delete(hashes, currentHash)

		fmt.Print(current.step, current.location, " ")

		// if current.step == 3 {
		// 	return 3
		// }

		nextBlizzards := moveBlizzards(current.blizzards, borders)
		neighbourds := getAvailableLocations(current.location, start, end, getBlizzardsMap(nextBlizzards), borders)
		for _, n := range neighbourds {

			if n == end {
				return current.step + 1, nextBlizzards
			}

			hash := hash{
				location: n,
				step:     current.step + 1,
			}

			// Add to queue only if we didn't encounter this hash before
			if _, ok := hashes[hash]; !ok {
				// I was able to add this optimization once I ran the algorithm
				// once and I knew how much it took to run it. It does help,
				// but not as much as you would expect.
				if manhattanDistance(n, end) < 310 {
					q = append(q, state{
						location:  n,
						blizzards: nextBlizzards,
						step:      current.step + 1,
					})
					hashes[hash] = struct{}{}
				}
			}
		}
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	blizzards := make([]Blizzard, 0)

	borders := [2]Location{}
	borders[0] = Location{1, 1}
	row := 0

	for scanner.Scan() {
		line := scanner.Text()
		processLine(line, row, &blizzards)
		if row == 0 {
			borders[1][1] = len(line)
		}
		row++
	}

	borders[1][0] = row
	start, end := startAndEndLocations(borders)

	exitsAfter1, blizzards1 := bfs(start, end, blizzards, borders)
	exitsAfter2, blizzards2 := bfs(end, start, blizzards1, borders)
	exitsAfter3, _ := bfs(start, end, blizzards2, borders)

	fmt.Printf("Exited after %d %d %d steps.", exitsAfter1, exitsAfter2, exitsAfter3)

}
