package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
SHAPE 1
####

SHAPE 2
.#.
###
.#.

SHAPE 3
..#
..#
###

SHAPE 4
#
#
#
#

SHAPE 5
##
##
*/

type Shape byte

const LEFT_BOUND = 0
const RIGHT_BOUND = 6
const DOWN_BOUND = 1

const (
	HORIZONTAL = 0
	PLUS       = 1
	L          = 2
	VERTICAL   = 3
	SQUARE     = 4
)

type Piece struct {
	shape       Shape
	coordinates [][2]int
}

// PieceHash is a useful struct in finding repeating patterns. Basically, two
// pieces with the same shape start at some point with the wind direction at the
// same place in the stream, fall the same number of moves and the rocks on the
// line just below have the same pattern, it's guaranteed that the structure
// between them will repeat.
type PieceHash struct {
	shape       Shape
	streamDir   int
	movesFallen int
	lowerHash   string // consider the value of the binary representation of the
	// lower

	// The following are just helpers to compute the height based on the pattern.
	addedHeight int // how much height was added by the piece

}

func areSame(ph1, ph2 PieceHash) bool {
	return ph1.shape == ph2.shape && ph1.streamDir == ph2.streamDir &&
		ph1.movesFallen == ph2.movesFallen &&
		ph1.lowerHash == ph2.lowerHash
}

// initPiece adds a new piece to the board, as described in the statement
func initPiece(shape Shape, height int) Piece {

	y := height + 4
	switch shape {
	case HORIZONTAL:
		return Piece{
			shape: HORIZONTAL,
			coordinates: [][2]int{
				{2, y}, {3, y}, {4, y}, {5, y},
			},
		}
	case PLUS:
		return Piece{
			shape: PLUS,
			coordinates: [][2]int{
				{3, y}, {2, y + 1}, {3, y + 1}, {4, y + 1}, {3, y + 2},
			},
		}

	case L:
		return Piece{
			shape: L,
			coordinates: [][2]int{
				{2, y}, {3, y}, {4, y}, {4, y + 1}, {4, y + 2},
			},
		}

	case VERTICAL:
		return Piece{
			shape: VERTICAL,
			coordinates: [][2]int{
				{2, y}, {2, y + 1}, {2, y + 2}, {2, y + 3},
			},
		}

	case SQUARE:
		return Piece{
			shape: SQUARE,
			coordinates: [][2]int{
				{2, y}, {3, y}, {2, y + 1}, {3, y + 1},
			},
		}

	default:
		return Piece{}
	}

}

// newPiece takes the current turn number and creates a new piece based on that.
func newPiece(turn, height int) Piece {
	return initPiece(Shape(turn%5), height)
}

// outOfBounds determines if some coordinate is out of the table's bounds.
func outOfBounds(coordinate [2]int, left, right, down int) bool {
	if coordinate[0] < left || coordinate[0] > right || coordinate[1] < down {
		return true
	}
	return false
}

type Direction [2]int

var DOWN Direction = [2]int{0, -1}
var LEFT Direction = [2]int{-1, 0}
var RIGHT Direction = [2]int{1, 0}

// dirFromSymbol takes a single symbol from the pattern and converts it to a
// direction.
func dirFromSymbol(symbol byte) Direction {
	if symbol == '<' {
		return LEFT
	} else {
		return RIGHT
	}
}

// dirsFromPattern extracts all direction from a pattern.
func dirsFromPattern(pattern string) []Direction {
	dirs := make([]Direction, 0, len(pattern))
	for i := 0; i < len(pattern); i++ {
		dirs = append(dirs, dirFromSymbol(pattern[i]))
	}
	return dirs
}

func nextDir(current *int, dirs []Direction) {
	*current = (*current + 1) % len(dirs)
}

// moveDown moves a piece in the specified direction, and changes the maximum
// board height if relevant. Returns whether the piece can still fall.
//
// Additionally, it returns the rock pattern the row just behind it.
func (p *Piece) movePiece(board *map[[2]int]struct{}, maxHeight *int, dir Direction, downbound int) (bool, string) {

	next := make([][2]int, len(p.coordinates))
	copy(next, p.coordinates)

	hasMoved := true

	for i := 0; i < len(next); i++ {

		// Move each coordinate in the specified direction
		next[i][0] = next[i][0] + dir[0]
		next[i][1] = next[i][1] + dir[1]

		// Find if there is already some other piece on the board
		_, ok := (*board)[next[i]]
		if ok || outOfBounds(next[i], LEFT_BOUND, RIGHT_BOUND, downbound) {
			hasMoved = false
			break
		}
	}

	if hasMoved {
		// Change the piece's coordiantes.
		p.coordinates = next

	}

	// Generate the hash of the rock pattern just below the piece.

	// This is always the lowest block in the piece.
	lowest := p.coordinates[0][1] - 1
	lowerHash := ""
	for i := 0; i <= 7; i++ {

		// Find how much below the lowest tile is the lowest tile in each row
		diff := 0
		for j := lowest; j >= 1; j-- {
			diff++
			if _, ok := (*board)[[2]int{i, j}]; ok {
				lowerHash += fmt.Sprint(diff) + "."
				break
			}
		}
	}

	lowerHash = strings.TrimSuffix(lowerHash, ".")

	return hasMoved, lowerHash
}

// dropPiece creates a new piece and drops it on the board.
func dropPiece(turn int, height *int, currentDir *int, dirs []Direction, board *map[[2]int]struct{}, downbound *int) PieceHash {

	// Initialize a new piece at the correct location
	piece := newPiece(turn, *height)

	cont := true
	ph := PieceHash{
		shape:       piece.shape,
		streamDir:   *currentDir,
		movesFallen: 0,
		lowerHash:   "",
		addedHeight: 0,
	}

	lowerHash := ""

	for cont {

		// First, move the piece in the stream direction. We don't care in this
		// case if the piece moves or not.
		_, lowerHash = piece.movePiece(board, height, dirs[*currentDir], *downbound)

		// Then, move the piece down. This should stop if the piece reaches a
		// dead end, which will also take care of the board and maximum height.
		cont, _ = piece.movePiece(board, height, DOWN, *downbound)

		// Change direction
		nextDir(currentDir, dirs)

		ph.movesFallen++

	}

	ph.lowerHash = lowerHash

	// Once the piece has dropped, change the coordinates and max height, if
	// relevant.

	for _, coord := range piece.coordinates {
		(*board)[coord] = struct{}{}
		if coord[1] > *height {
			ph.addedHeight += (coord[1] - *height)
			*height = coord[1]
		}
	}

	return ph

}

func dropAll(turns int, directons []Direction) (map[[2]int]struct{}, int) {

	height := 0
	board := make(map[[2]int]struct{})
	currentDir := 0

	downbound := DOWN_BOUND

	hashesMap := make(map[PieceHash]int, 0)
	hashesArray := []PieceHash{}
	computeDuplicate := true
	duplicates := []PieceHash{}

	heightUntilDuplicates := 0

	turn := 0

	for computeDuplicate {

		hash := dropPiece(turn, &height, &currentDir, directons, &board, &downbound)
		//fmt.Println(hash)
		// If the hash already exists
		if _, ok := hashesMap[hash]; ok {
			fmt.Println(turn, "duplicate hash found", hash, hashesMap[hash])
			computeDuplicate = false
			heightUntilDuplicates = height - hash.addedHeight
			for x := hashesMap[hash]; x < turn; x++ {
				duplicates = append(duplicates, hashesArray[x])
			}

			// If the hash doesn't exist, just add it.
		} else {
			hashesMap[hash] = turn
			hashesArray = append(hashesArray, hash)
			//fmt.Println(hash)
		}
		turn++
	}

	// Disconsider last drop, will only count using duplicates now.

	turn--

	heightDuplicatesAdd := 0
	duplicatesLength := len(duplicates)
	for _, d := range duplicates {
		heightDuplicatesAdd += d.addedHeight
	}

	for turn+duplicatesLength < turns {
		heightUntilDuplicates += heightDuplicatesAdd
		turn += duplicatesLength
	}

	i := 0

	for turn < turns {
		turn++
		heightUntilDuplicates += duplicates[i].addedHeight
		i++
	}

	return board, heightUntilDuplicates
}

func drawBoard(board *map[[2]int]struct{}, height int) {

	out := ""

	for y := height; y >= 1; y-- {

		sheight := fmt.Sprint(y)
		for len(sheight) <= 4 {
			sheight += " "
		}

		out += sheight + "|"
		for x := 0; x < 7; x++ {
			if _, ok := (*board)[[2]int{x, y}]; ok {
				out += "#"
			} else {
				out += "."
			}
		}
		out += "|\n"
	}
	out += "    +-------+\n"
	fmt.Println(out)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	pattern := scanner.Text()
	directions := dirsFromPattern(pattern)

	//_, h := dropAll(2022, directions)
	_, h := dropAll(1000000000000, directions)
	fmt.Printf("The maximum height of the tower is %d\n", h)

	//drawBoard(&b, h)
}
