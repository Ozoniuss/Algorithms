package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Ozoniuss/go-aoc/twod"
)

type Action byte

const (
	TURN_ON Action = iota
	TURN_OFF
	TOGGLE
)

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

// parseLine reads an input line and returns the action to be performed, as well
// as the two locations
func parseLine(input string) (Action, [2]twod.Location) {

	var x1, y1, x2, y2 int
	var act Action

	if strings.Contains(input, "turn on") {
		fmt.Sscanf(input, "turn on %d,%d through %d,%d", &x1, &y1, &x2, &y2)
		act = TURN_ON
	} else if strings.Contains(input, "turn off") {
		fmt.Sscanf(input, "turn off %d,%d through %d,%d", &x1, &y1, &x2, &y2)
		act = TURN_OFF
	} else if strings.Contains(input, "toggle") {
		fmt.Sscanf(input, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)
		act = TOGGLE
	} else {
		panic("invalid input")
	}
	return act, [2]twod.Location{{x1, y1}, {x2, y2}}
}

// performAction performs the respective action, as described by the problem
// statement. The turned_on map holds all lights that are currently turned
// on.
func performAction(act Action, coords [2]twod.Location, turned_on map[twod.Location]struct{}) {

	left_x := coords[0][0]
	left_y := coords[0][1]

	right_x := coords[1][0]
	right_y := coords[1][1]

	for row := left_x; row <= right_x; row++ {
		for col := left_y; col <= right_y; col++ {
			loc := twod.Location{row, col}
			if act == TURN_ON {
				turned_on[loc] = struct{}{}
			} else if act == TURN_OFF {
				delete(turned_on, loc)
			} else {
				if _, ok := turned_on[loc]; ok {
					delete(turned_on, loc)
				} else {
					turned_on[loc] = struct{}{}
				}
			}
		}
	}
}

// performActionBrightness performs the respective action, as described by the
// problem statement. It is similar to performAction, except that it takes a
// brightness map as input.
func performActionBrightness(act Action, coords [2]twod.Location, brightness map[twod.Location]int) {

	left_x := coords[0][0]
	left_y := coords[0][1]

	right_x := coords[1][0]
	right_y := coords[1][1]

	for row := left_x; row <= right_x; row++ {
		for col := left_y; col <= right_y; col++ {
			loc := twod.Location{row, col}
			switch act {
			case TURN_ON:
				if _, ok := brightness[loc]; ok {
					brightness[loc]++
				} else {
					brightness[loc] = 1
				}
			case TURN_OFF:
				if _, ok := brightness[loc]; ok {
					brightness[loc]--
					if brightness[loc] == 0 {
						delete(brightness, loc)
					}
				}
			case TOGGLE:
				if _, ok := brightness[loc]; ok {
					brightness[loc] += 2
				} else {
					brightness[loc] = 2
				}
			}
		}
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
		turned_on := make(map[twod.Location]struct{})
		for scanner.Scan() {
			act, coords := parseLine(scanner.Text())
			performAction(act, coords, turned_on)
		}
		fmt.Printf("There are a total of %d lit lights.", len(turned_on))
	} else {
		brightness := make(map[twod.Location]int)
		for scanner.Scan() {
			act, coords := parseLine(scanner.Text())
			performActionBrightness(act, coords, brightness)
		}
		total := 0
		for _, br := range brightness {
			total += br
		}
		fmt.Printf("The total brightness is %d.", total)
	}
}
