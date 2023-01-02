package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Syntactic sugar.
type Dimensions [3]int

// mustAtoi does the same thing as atoi, except less anoying
func mustAtoi(number string) int {
	num, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}
	return num
}

func min(nums ...int) int {
	min := 1<<31 - 1
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

// parseLine parses one input line representing the present dimensions
// and returns its dimensions as a Dimensions type.
func parseLine(line string) Dimensions {
	parts := strings.Split(line, "x")
	return Dimensions{
		mustAtoi(parts[0]), mustAtoi(parts[1]), mustAtoi(parts[2]),
	}
}

// computeRequiredPaper computes the required amount of paper in order to wrap
// a present.
func computeRequiredPaper(dim Dimensions) int {
	area1 := dim[0] * dim[1]
	area2 := dim[1] * dim[2]
	area3 := dim[0] * dim[2]
	return 2*(area1+area2+area3) + min(area1, area2, area3)
}

// computeRequiredRibbon computes the required amount of ribbon in order to wrap
// a present.
func computeRequiredRibbon(dim Dimensions) int {

	// minimum perimeter
	presentRibbon := min(2*(dim[0]+dim[1]), 2*(dim[1]+dim[2]), 2*(dim[0]+dim[2]))
	bowRibbon := dim[0] * dim[1] * dim[2]

	return presentRibbon + bowRibbon
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
	total := 0

	part := readPart()

	for scanner.Scan() {
		line := scanner.Text()
		dimensions := parseLine(line)
		if part == 1 {
			total += computeRequiredPaper(dimensions)
		} else {
			total += computeRequiredRibbon(dimensions)
		}
	}

	if part == 1 {
		fmt.Printf("Total area of paper required is %d", total)
	} else {
		fmt.Printf("Total area of ribbon required is %d", total)
	}
}
