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

// addStrenght adds the strenght of a signal to the signals map, if that signal
// exists in the map. The value of the signal is taken from the registry and
// multiply by the cycle's number.
func addStrenght(signals *map[int]int, cycle int, register int) {
	if _, ok := (*signals)[cycle]; ok == true {
		(*signals)[cycle] = cycle * register
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	// current cycle
	cycle := 1
	register := 1

	signals := map[int]int{
		20:  0,
		60:  0,
		100: 0,
		140: 0,
		180: 0,
		220: 0,
	}

	for scanner.Scan() {
		instruction, arg := processLine(scanner.Text())
		if instruction == NOOP {
			cycle += 1
			addStrenght(&signals, cycle, register)
		} else if instruction == ADDX {
			cycle += 1
			addStrenght(&signals, cycle, register)
			register += arg
			cycle += 1
			addStrenght(&signals, cycle, register)
		}
	}

	sum := 0
	for _, strength := range signals {
		sum += strength
	}

	fmt.Printf("The sum of the six signal strenghts is %d\n", sum)
}
