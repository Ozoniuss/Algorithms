package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToWire(name string) [2]byte {
	var wire [2]byte

	// characters are always ASCII so they occupy 1 byte
	if len(name) == 1 {
		wire = [2]byte{0, name[0]}
	} else {
		wire = [2]byte{name[0], name[1]}
	}
	return wire
}

func lshift(x int, n int) int {
	return (x << n) % 65536
}

func rshift(x int, n int) int {
	return (x >> n) % 65536

}

// processLine processes one line of text from the input file.
func processLine(line string, signals map[[2]byte]int) error {
	content := strings.Split(line, " ")
	// send simple signal
	if content[1] == "->" {
		wire := stringToWire(content[2])
		sl, err := strconv.Atoi(content[0])
		if err != nil {
			strength, ok := signals[stringToWire(content[0])]
			if !ok {
				return errors.New("not found")
			}
			signals[wire] = strength
			return nil
		}

		fmt.Println("aiciii", sl)
		signals[wire] = sl
		return nil
	}

	// send NOT signal
	if content[2] == "->" {
		if content[0] != "NOT" {
			panic("NOT not encountered")
		}
		wire1 := stringToWire(content[1])
		wire2 := stringToWire(content[3])

		if _, ok := signals[wire1]; !ok {
			return errors.New("not found")
		}

		signals[wire2] = 65535 - signals[wire1]
		return nil
	}

	// Gate
	if content[3] == "->" {

		var leftPart, rightPart int
		var ok bool

		// left or right may be number
		if val, err := strconv.Atoi(content[0]); err == nil {
			leftPart = val
		} else {
			leftPart, ok = signals[stringToWire(content[0])]
			if !ok {
				return errors.New("not found")
			}
		}

		if val, err := strconv.Atoi(content[2]); err == nil {
			rightPart = val
		} else {
			rightPart, ok = signals[stringToWire(content[2])]
			if !ok {
				return errors.New("not found")
			}
		}

		wire := stringToWire(content[4])
		switch content[1] {
		case "OR":
			signals[wire] = (leftPart | rightPart) % 65536
		case "AND":
			signals[wire] = (leftPart & rightPart) % 65536
		case "LSHIFT":
			signals[wire] = lshift(leftPart, rightPart)
		case "RSHIFT":
			signals[wire] = rshift(leftPart, rightPart)
		default:
			panic("invalid gate")
		}
	}
	return nil
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

	signals := make(map[[2]byte]int, 100)
	scanner := bufio.NewScanner(f)
	part := readPart()
	if part == 1 {
		lines := make(map[int]string, 100)
		i := 0
		for scanner.Scan() {
			lines[i] = scanner.Text()
			i++
		}

		// I know, not the best but whatever
		var keepgoing = true
		for keepgoing {
			for i, line := range lines {
				err := processLine(line, signals)
				if err == nil {
					delete(lines, i)
				}
			}
			if len(lines) == 0 {
				keepgoing = false
			}
			fmt.Println(len(lines))
		}
		for sig := range signals {
			fmt.Println(string(sig[:]), signals[sig])
		}
		fmt.Println(signals[[2]byte{0, 'a'}])
	} else {
		panic("unimplemented")
	}
}
