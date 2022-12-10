package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	UP    = "U"
	DOWN  = "D"
	LEFT  = "L"
	RIGHT = "R"
)

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

// manhattanDistance returns the manhattan distance between two points.
func manhattanDistance(p1 [2]int, p2 [2]int) int {
	return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
}

// needsMove returns wether the tail needs to be moved as well, after head and
// tail reached in the provided positions.
func needsMove(head [2]int, tail [2]int) bool {
	if manhattanDistance(head, tail) > 2 {
		return true
	}
	if manhattanDistance(tail, head) == 2 {
		// If they are located diagonally, no need to move.
		if abs(head[0]-tail[0]) == 1 && abs(head[1]-tail[1]) == 1 {
			return false

			// They're on a line and there's at least one empty spot between
			// them.
		} else {
			return true
		}
	}
	return false
}

// moveIfNeedsMove makes a move if it's possible to make a move and returns the
// new position. It's possible to determine the new location of the head and
// tail based on this position.
func moveIfNeedsMove(head [2]int, tail [2]int) ([2]int, [2]int) {
	if needsMove(head, tail) {
		// todo: explain
		if manhattanDistance(head, tail) == 4 {
			tail[0] = (head[0] + tail[0]) / 2
			tail[1] = (head[1] + tail[1]) / 2
		} else if head[0]-tail[0] == 2 {
			tail[0] = head[0] - 1
			tail[1] = head[1]
			return head, tail
		} else if head[0]-tail[0] == -2 {
			tail[0] = head[0] + 1
			tail[1] = head[1]
			return head, tail
		} else if head[1]-tail[1] == 2 {
			tail[0] = head[0]
			tail[1] = head[1] - 1
			return head, tail
		} else if head[1]-tail[1] == -2 {
			tail[0] = head[0]
			tail[1] = head[1] + 1
			return head, tail
		}
	}
	// Doesn't move
	return head, tail
}

// move takes the current position of the head and the tail and the direction
// of the move, and will return the new position of the head and tail after the
// move.
func move(direction string, head [2]int, tail [2]int) ([2]int, [2]int) {
	if direction == UP {
		head[0] -= 1

		// Move the tail up behind the head
		if needsMove(head, tail) {
			tail[0] = head[0] + 1
			tail[1] = head[1]
		}
	} else if direction == DOWN {
		head[0] += 1

		// Move the tail down behind the head
		if needsMove(head, tail) {
			tail[0] = head[0] - 1
			tail[1] = head[1]
		}
	}
	if direction == LEFT {
		head[1] -= 1

		// Move the tail left behind the head
		if needsMove(head, tail) {
			tail[0] = head[0]
			tail[1] = head[1] + 1
		}
	}
	if direction == RIGHT {
		head[1] += 1

		// Move the tail right behind the head
		if needsMove(head, tail) {
			tail[0] = head[0]
			tail[1] = head[1] - 1
		}
	}

	return head, tail
}

// move rope moves the entire rope in one direction and returns the new
// positions of the rope.
func moveRope(direction string, rope [10][2]int) [10][2]int {
	newRope := [10][2]int{}
	head := rope[0]

	fmt.Println("before ", head)

	// Move the head
	if direction == UP {
		head[0] -= 1
	} else if direction == DOWN {
		head[0] += 1
	}
	if direction == LEFT {
		head[1] -= 1
	}
	if direction == RIGHT {
		head[1] += 1
	}

	newRope[0] = head

	// move all other knots based on the head.
	for i := 0; i <= 8; i++ {
		_, next := moveIfNeedsMove(newRope[i], rope[i+1])
		newRope[i+1] = next
	}

	return newRope
}

func processLine(line string) (string, int) {
	parts := strings.Split(line, " ")
	direction := parts[0]
	noMoves, _ := strconv.ParseInt(parts[1], 10, 32)
	return direction, int(noMoves)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	// Initial tail positions.
	// head := [2]int{0, 0}
	// tail := [2]int{0, 0}

	// The locations visited by the tail
	visited := make(map[[2]int]struct{})
	visited[[2]int{0, 0}] = struct{}{}

	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	direction, noMoves := processLine(line)
	// 	for i := 0; i < noMoves; i++ {
	// 		head, tail = move(direction, head, tail)
	// 		visited[tail] = struct{}{}
	// 	}
	// }

	//fmt.Printf("The tail travelled through %d places.\n", len(visited))

	rope := [10][2]int{}

	for scanner.Scan() {
		line := scanner.Text()
		direction, noMoves := processLine(line)
		for i := 0; i < noMoves; i++ {
			newRope := moveRope(direction, rope)
			rope = newRope
			visited[rope[len(rope)-1]] = struct{}{}
		}
	}

	fmt.Printf("The tail travelled through %d places.\n", len(visited))

}
