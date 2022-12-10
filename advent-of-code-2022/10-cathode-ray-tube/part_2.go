package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	NOOP = "noop"
	ADDX = "addx"
)

const WIDTH = 40

// abs returns the absolute value of an integer.
func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

// processLine takes as input the value of a line as a string and returns the
// instruction and the argument value.
func processLine(line string) (string, int) {
	parts := strings.Split(line, " ")
	if parts[0] == "noop" {
		return NOOP, 0
	} else {
		arg, _ := strconv.Atoi(parts[1])
		return ADDX, arg
	}
}

// drawPixel draws a pixel on the screen, based on the cycle number and the
// register value, which dictataes the stripe position.
func drawPixel(cycle, register int) {
	pixelPosition := (cycle - 1) % WIDTH
	if pixelPosition == 0 {
		fmt.Print("\n")
	}
	if abs(pixelPosition-register) < 2 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	// current cycle
	cycle := 0
	register := 1

	for scanner.Scan() {
		instruction, arg := processLine(scanner.Text())
		if instruction == NOOP {
			cycle += 1
			drawPixel(cycle, register)
		} else if instruction == ADDX {
			cycle += 1
			drawPixel(cycle, register)
			cycle += 1
			drawPixel(cycle, register)
			register += arg
		}
	}

}
