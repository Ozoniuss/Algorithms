package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ROCK = 3
	SAND = 4
)

// processLine takes a line as a string and adds all coordinates that are part
// of the formed path to the list of rocks.
func processLine(line string, rocks *map[[2]int]byte) {

	partsString := strings.Split(line, " -> ")
	parts := make([][2]int, 0)

	for _, p := range partsString {
		numbers := strings.Split(p, ",")
		numberLeft, _ := strconv.Atoi(numbers[0])
		numberRight, _ := strconv.Atoi(numbers[1])
		parts = append(parts, [2]int{numberLeft, numberRight})
	}

	for i := 0; i < len(parts)-1; i++ {
		src := parts[i]
		dst := parts[i+1]

		// Note that x is column and y is row.
		if src[0] == dst[0] {
			if src[1] >= dst[1] {
				for j := src[1]; j >= dst[1]; j-- {
					(*rocks)[[2]int{j, src[0]}] = ROCK
				}
			} else {
				for j := src[1]; j <= dst[1]; j++ {
					(*rocks)[[2]int{j, src[0]}] = ROCK
				}
			}
		} else if src[1] == dst[1] {
			if src[0] >= dst[0] {
				for j := src[0]; j >= dst[0]; j-- {
					(*rocks)[[2]int{src[1], j}] = ROCK
				}
			} else {
				for j := src[0]; j <= dst[0]; j++ {
					(*rocks)[[2]int{src[1], j}] = ROCK
				}
			}
		} else {
			panic("wtf")
		}
	}
}

// dropSand drops a unit of sand from src, respecting the rules described in the
// statement. If the sand stops, it is added to the rocks map as SAND. If it
// falls to the abyss, it is not added to the rocks map.
func dropSand(src [2]int, rocks *map[[2]int]byte, bottom int) bool {
	currentPos := src
	_, blocked := (*rocks)[currentPos]
	for !blocked {
		currentPos[0]++

		// Passed bottom, will fall forever
		if currentPos[0] > bottom {
			return false
		}
		_, blocked = (*rocks)[currentPos]
	}

	// Reached something (above the place where something already existed)
	reachedAt := currentPos
	reachedAt[0]--

	downLeft := [2]int{reachedAt[0] + 1, reachedAt[1] - 1}
	downRight := [2]int{reachedAt[0] + 1, reachedAt[1] + 1}

	if _, leftOccupied := (*rocks)[downLeft]; !leftOccupied {
		return dropSand(downLeft, rocks, bottom)
	} else if _, rightOccupied := (*rocks)[downRight]; !rightOccupied {
		return dropSand(downRight, rocks, bottom)
	} else {
		(*rocks)[reachedAt] = SAND
	}

	return true
}

// dropSandUntilTop drops a unit of sand from src, respecting the rules
// described in the statement for part 2. If the sand stops, it is added to the
// rocks map as SAND. It will not fall to the abyss.
//
// The function returns the location where the sand reached.
func dropSandUntilTop(src [2]int, rocks *map[[2]int]byte) [2]int {
	currentPos := src
	_, blocked := (*rocks)[currentPos]

	// Source is blocked
	if blocked {
		return currentPos
	}

	// Will stop at some point.
	for !blocked {
		currentPos[0]++
		_, blocked = (*rocks)[currentPos]
	}

	// Reached something (above the place where something already existed)
	reachedAt := currentPos
	reachedAt[0]--

	downLeft := [2]int{reachedAt[0] + 1, reachedAt[1] - 1}
	downRight := [2]int{reachedAt[0] + 1, reachedAt[1] + 1}

	if _, leftOccupied := (*rocks)[downLeft]; !leftOccupied {
		return dropSandUntilTop(downLeft, rocks)
	} else if _, rightOccupied := (*rocks)[downRight]; !rightOccupied {
		return dropSandUntilTop(downRight, rocks)
	} else {
		(*rocks)[reachedAt] = SAND
	}

	return reachedAt
}

func drawRocks(rocks *map[[2]int]byte) {
	row_min := 1000
	row_max := 0
	col_min := 1000
	col_max := 0

	for r := range *rocks {
		if r[0] < row_min {
			row_min = r[0]
		}
		if r[0] > row_max {
			row_max = r[0]
		}
		if r[1] < col_min {
			col_min = r[1]
		}
		if r[1] > col_max {
			col_max = r[1]
		}
	}

	out := "   "
	for col := col_min - 5; col < col_max+5; col++ {
		out += fmt.Sprint(col) + " "
	}
	out += "\n"
	for row := row_min - 5; row < row_max+5; row++ {
		if len(fmt.Sprint(row)) == 1 {
			out += fmt.Sprint(row) + "  "
		}
		if len(fmt.Sprint(row)) == 2 {
			out += fmt.Sprint(row) + " "
		}
		if len(fmt.Sprint(row)) == 3 {
			out += fmt.Sprint(row) + ""
		}
		for col := col_min - 5; col < col_max+5; col++ {

			// There's something
			if val, ok := (*rocks)[[2]int{row, col}]; ok {
				if val == ROCK {
					out += " # "
				} else {
					if val == SAND {
						out += " O "
					}
				}
			} else {
				out += " . "
			}
			out += " "
		}
		out += "\n"
	}

	f, err := os.OpenFile("map.txt", os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString(out)
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	part := 2

	// Map holding the positions of all rocks.
	rocks := make(map[[2]int]byte)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		processLine(line, &rocks)
	}

	src := [2]int{0, 500}

	if part == 1 {
		// This has been determined by visualizing the drawing
		bottom := 158
		canDrop := true
		count := 0

		for canDrop {
			canDrop = dropSand(src, &rocks, bottom)
			count++
		}

		drawRocks(&rocks)

		fmt.Printf("A total of %d sand units fell.\n", count-1)
	} else {
		bottom := 159
		count := 0

		// fill the bottom with rocks
		for i := 0; i < 1000; i++ {
			rocks[[2]int{bottom, i}] = ROCK
		}

		s := [2]int{0, 0}
		for s != src {
			s = dropSandUntilTop(src, &rocks)
			count += 1
		}

		fmt.Printf("A total of %d sand units fell.\n", count)
	}

}
