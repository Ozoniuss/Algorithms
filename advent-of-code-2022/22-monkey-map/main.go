package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// processLines reads a line from the input file, parses it and adds the
// coordinates of the tiles from that line to the board.
func processLine(line string, lineNumber int, board map[[2]int]byte) {
	if line == "" {
		return
	}
	for columnNumber := 0; columnNumber < len(line); columnNumber++ {
		if tile := line[columnNumber]; tile != ' ' {
			// The numbering starts from one in the problem.
			board[[2]int{lineNumber + 1, columnNumber + 1}] = tile
		}
	}
}

/*
	Note that because of the movement that happens when you reach the end of a
	row or column, there is no horizontal or vertical line on the map that
	contains two different sections of tiles. Something like the drawing below
	is not possible.

	#...#   ....#
	..#..   .....
	......#......
*/

// findBorder returns two maps representing the tiles that are either at the
// borders of a row or at the borders of a column. If one reaches those borders,
// these maps will help identify the new location you "wrap" to.
func findBorders(rows, cols int, board map[[2]int]byte) (
	map[int][2][2]int, map[int][2][2]int) {

	colBorders := make(map[int][2][2]int)
	rowBorders := make(map[int][2][2]int)

	// We know the total number of rows.
	for row := 1; row <= rows; row++ {
		col := 1
		borders := [2][2]int{}

		// Find the leftmost tile
		_, tile := board[[2]int{row, col}]
		for !tile {
			col++
			_, tile = board[[2]int{row, col}]
		}
		borders[0] = [2]int{row, col}

		// Find the rightmost tile
		_, tile = board[[2]int{row, col}]
		for tile {
			col++
			_, tile = board[[2]int{row, col}]
		}
		borders[1] = [2]int{row, col - 1}

		rowBorders[row] = borders
	}

	// We have also computed the total number of columns.
	for col := 1; col <= cols; col++ {
		row := 1
		borders := [2][2]int{}

		// Find the upmost tile
		_, tile := board[[2]int{row, col}]
		for !tile {
			row++
			_, tile = board[[2]int{row, col}]
		}
		borders[0] = [2]int{row, col}

		// Find the downmost tile
		_, tile = board[[2]int{row, col}]
		for tile {
			row++
			_, tile = board[[2]int{row, col}]
		}
		borders[1] = [2]int{row - 1, col}

		colBorders[col] = borders

	}
	return rowBorders, colBorders
}

// findCubeMapping returns all locations where you can go off the map
// (including a direction), as well as the new location once you wrap along
// the cube's edges.
//
// Honestly I'm hardcoding this because this task has no nice idea whatsoever
// and is just fucking retarded.
func findCubeMapping() map[[2][2]int][2][2]int {

	// Note that in the mappings you need the direction as well, because of the
	// corners. If you reach the edge through a corner, it can go to two
	// different places depending on the direction.
	mappings := make(map[[2][2]int][2][2]int)

	// Edge pairs helps building the overlapping edges of the cube. This
	// means that if you start on the first edge, you end up on the second
	// edge that's paired with it. The coordinate and direction helps to add
	// all locations belonging to that edge.
	//
	// This makes more sense in the lower for loop.
	edgePairs := [][6][2]int{
		{{1, 51}, DOWN, {150, 1}, UP},
		{{1, 51}, RIGHT, {151, 1}, DOWN},
		{{1, 101}, RIGHT, {200, 1}, RIGHT},
		{{1, 150}, DOWN, {150, 100}, UP},
		{{50, 101}, RIGHT, {51, 100}, DOWN},
		{{100, 51}, UP, {101, 50}, LEFT},

		{{150, 51}, RIGHT, {151, 50}, DOWN},
	}

	// directionChange keeps track of the direction change from the first edge
	// to the second edge, and from the second edge to the first edge.
	//
	// For example, for the first pair of matching edges, if the player goes
	// of that edge going left, they end up coming from the second edge facing
	// right. And if they go off the second edge facing left, they end up on
	// the first edge facing right.
	directionChange := [][4][2]int{
		{LEFT, RIGHT, LEFT, RIGHT},
		{UP, RIGHT, LEFT, DOWN},
		{UP, UP, DOWN, DOWN},
		{RIGHT, LEFT, RIGHT, LEFT},
		{DOWN, LEFT, RIGHT, UP},
		{LEFT, DOWN, UP, RIGHT},
		{DOWN, LEFT, RIGHT, UP},
	}

	for idx := 0; idx < len(edgePairs); idx++ {

		// Get the starting location of each of the paired edges, and start
		// adding all locations of those edges.
		currentEdge1 := edgePairs[idx][0]
		dir1 := edgePairs[idx][1]
		currentEdge2 := edgePairs[idx][2]
		dir2 := edgePairs[idx][3]

		// Match every location of the two edges
		for i := 0; i < 50; i++ {

			// Each location on an edge is mapped to exactly one location on the
			// other edge.
			mappings[[2][2]int{currentEdge1, directionChange[idx][0]}] = [2][2]int{currentEdge2, directionChange[idx][1]}
			mappings[[2][2]int{currentEdge2, directionChange[idx][2]}] = [2][2]int{currentEdge1, directionChange[idx][3]}

			// Continue with all locations
			currentEdge1 = [2]int{currentEdge1[0] + dir1[0], currentEdge1[1] + dir1[1]}
			currentEdge2 = [2]int{currentEdge2[0] + dir2[0], currentEdge2[1] + dir2[1]}
		}
	}

	return mappings
}

var (
	LEFT  = [2]int{0, -1}
	RIGHT = [2]int{0, +1}
	UP    = [2]int{-1, 0}
	DOWN  = [2]int{1, 0}
)

// turnLeft takes a turn left from the current direction and returns the new
// direction.
func turnLeft(direction [2]int) [2]int {
	// Think of it as multiplying by complex number -i.
	return [2]int{-direction[1], direction[0]}
}

// turnRight takes a turn right from the current direction and returns the new
// direction.
func turnRight(direction [2]int) [2]int {
	// Think of it as multiplying by complex number -i.
	return [2]int{direction[1], -direction[0]}
}

// parsePassword takes a password as an input string and converts it to an array
// of directions representing your direction at each move.
//
// Note: this fails at part (b).

func parsePassword(password string) [][2]int {

	directions := make([][2]int, 0)
	current := RIGHT

	re := regexp.MustCompile("([0-9]+)|R|L")
	tokens := re.FindAllString(password, -1)

	// Note that I implemented the turn functions after I wrote this loop.
	for _, token := range tokens {
		if token == "R" {
			// Think of it as multiplying by complex number -i.
			current = [2]int{current[1], -current[0]}
		} else if token == "L" {
			// Think of it as multiplying by complex number i.
			current = [2]int{-current[1], current[0]}
			// Token is a number
		} else {
			value, _ := strconv.Atoi(token)
			for i := 0; i < value; i++ {
				directions = append(directions, current)
			}
		}
	}
	return directions
}

/*
	Rather than storing the direction at each move, it's better to store whether
	there's a change in direction at each move, and which one if there is.
	That's because unlike (a), at part (b) wrapping around the board not only
	changes your current direction but also the following directions. However,
	direction changes stay the same.

	This function actually leads to a nicer implementation and can also be
	applied at part (a).
*/

// parsePasswordDirectionChangeVector also parses the input string, but converts
// it to a direction change vector. That is, instead of storing all directions
// and applying all changes, it stores whether you need to change the direction
// or not.
func parsePasswordDirectionChangeVector(password string) []byte {

	directions := make([]byte, 0)

	re := regexp.MustCompile("([0-9]+)|R|L")
	tokens := re.FindAllString(password, -1)

	for _, token := range tokens {
		if token == "R" {
			directions = append(directions, 'R')
		} else if token == "L" {
			directions = append(directions, 'L')
		} else {
			value, _ := strconv.Atoi(token)
			for i := 0; i < value; i++ {
				directions = append(directions, '.')
			}
		}
	}

	return directions
}

// move takes in the current position, the direction of the move and the board,
// and returns the next position after the move. In this function, wrapping
// around the board happens in the same line. The function returns the location
// after the move.
func move(current [2]int, direction [2]int, board map[[2]int]byte,
	rowBorders, colBorders map[int][2][2]int) [2]int {

	next := [2]int{}

	// Make the move
	next[0] = current[0] + direction[0]
	next[1] = current[1] + direction[1]

	_, ok := board[next]

	// Wrap around the map in the same line.
	if !ok {
		if direction == LEFT {
			row := next[0]
			next = rowBorders[row][1]
		} else if direction == RIGHT {
			row := next[0]
			next = rowBorders[row][0]
		} else if direction == UP {
			col := next[1]
			next = colBorders[col][1]
		} else if direction == DOWN {
			col := next[1]
			next = colBorders[col][0]
		}
	}
	tile := board[next]

	if tile == '.' {
		return next
		// If the next location is blocked, don't make the move.
	} else if tile == '#' {
		return current
	} else {
	}

	return next
}

// moveCube is similar to move, except that when it reaches a border, it wraps
// around the cube, not the line. It also returns the new location and direction
// you're facing.
func moveCube(current [2]int, currentDir [2]int, directionNumber int, directions []byte, board map[[2]int]byte, cubeMappings map[[2][2]int][2][2]int) ([2]int, [2]int) {

	next := [2]int{}
	nextDir := currentDir

	// If at this move a change in direction is detected, don't move and just
	// return the new direction.
	if directions[directionNumber] == 'L' {
		return current, turnLeft(currentDir)
	} else if directions[directionNumber] == 'R' {
		return current, turnRight(currentDir)
	}

	// Make the move
	next[0] = current[0] + currentDir[0]
	next[1] = current[1] + currentDir[1]

	_, ok := board[next]

	// Wrap around the map
	if !ok {
		reachedWithDir := cubeMappings[[2][2]int{current, currentDir}]
		next = reachedWithDir[0]
		nextDir = reachedWithDir[1]
	}
	tile := board[next]

	if tile == '.' {
		return next, nextDir
	} else if tile == '#' {
		return current, currentDir
	} else {
	}

	return current, currentDir

}

// makeMoves makes all of the moves from the list of directions. It returns the
// last position reached on the board.
func makeMoves(directions [][2]int, board map[[2]int]byte,
	rowBorders, colBorders map[int][2][2]int) [2]int {

	// Dirty trick but sue me.
	start := rowBorders[1][0]
	current := start

	for _, direction := range directions {
		current = move(current, direction, board, rowBorders, colBorders)
	}

	return current
}

// makeMovesCube does the same as makeMoves, except that the next move is
// computed using moveCube.
func makeMovesCube(directions []byte, board map[[2]int]byte,
	cubeMappings map[[2][2]int][2][2]int, start [2]int) ([2]int, [2]int) {

	current := start
	currentDir := RIGHT

	for dirNum := range directions {
		current, currentDir = moveCube(current, currentDir, dirNum, directions, board, cubeMappings)
	}

	return current, currentDir
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	board := make(map[[2]int]byte)

	rows := 0
	columns := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		processLine(scanner.Text(), rows, board)
		if len(scanner.Text()) > columns {
			columns = len(scanner.Text())
		}
		rows++
	}

	// Read password
	scanner.Scan()
	password := scanner.Text()

	part := 2

	if part == 1 {
		directions := parsePassword(password)
		rowBorders, colBorders := findBorders(rows, columns, board)

		final := makeMoves(directions, board, rowBorders, colBorders)
		fmt.Printf("Reached at position %v with facing %v\n", final, directions[len(directions)-1])
	} else {

		// Dirty trick to find the starting point because I'm lazy.
		rowBorders, _ := findBorders(rows, columns, board)

		directions := parsePasswordDirectionChangeVector(password)

		cubeMappings := findCubeMapping()
		final, dir := makeMovesCube(directions, board, cubeMappings, rowBorders[1][0])
		fmt.Printf("Reached at position %v with facing %v\n", final, dir)
	}

}
