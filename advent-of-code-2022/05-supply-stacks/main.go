package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// scanCrates scans the first part of the text file, which is representative of
// the crates and their position. Returns an error if it doesn't encounter the
// new line separator between the crates representation and the move
// instructions.
func scanCrates(s *bufio.Scanner) ([]string, error) {

	out := []string{}

	for s.Scan() {
		value := s.Text()
		if value == "" {
			return out, nil
		}
		out = append(out, value)
	}

	// Should be unnecessary
	return nil, errors.New("Did not encounter empty line separator.")
}

func scanMoves(s *bufio.Scanner) ([][3]byte, error) {
	out := [][3]byte{}

	for s.Scan() {
		value := s.Text()
		if value == "" {
			return out, nil
		}
		parts := strings.Split(value, " ")

		// Assume there will be no errors
		noCrates, _ := strconv.ParseInt(parts[1], 10, 8)
		src, _ := strconv.ParseInt(parts[3], 10, 8)
		dest, _ := strconv.ParseInt(parts[5], 10, 8)

		out = append(out, [3]byte{byte(src), byte(dest), byte(noCrates)})
	}

	return out, nil
}

// generateCratesStructure takes the input lines representing the location of
// the crates and returns a mapping between the line numbers and a string
// representing the crates on that line, from bottom to top.
func generateCratesStructure(cratesString []string) (map[byte]string, error) {

	crates := make(map[byte]string)

	// positions stores the position in the actual line where the crate numbers.
	//
	// If there is a crate on that column, its location in the line itself
	// will be the same as the location of that column number in the bottom
	// line.
	positions := make(map[byte]byte)

	lineNumbers := cratesString[len(cratesString)-1]

	for i := 0; i < len(lineNumbers); i++ {
		if value := string(lineNumbers[i]); value != " " {

			// We know that the character is a digit.
			digit, err := strconv.ParseInt(value, 10, 8)
			if err != nil {
				return nil, err
			}

			// Make the position entry.
			positions[byte(digit)] = byte(i)
		}
	}

	// Loop through the lines representing the crates' positions, starting from
	// the last one until the top one.
	for j := len(cratesString) - 2; j >= 0; j-- {
		for column := 1; column <= 9; column++ {

			lineLocation := positions[byte(column)]

			// No need to do anything if there is no crate
			if string(cratesString[j][lineLocation]) == " " {
				continue
			}

			// There is no crate on that column yet
			if _, ok := crates[byte(column)]; !ok {

				// From the current line, the string we want to add to this
				// column is at the same location where the number of the column
				// was in the last line
				crates[byte(column)] = string(cratesString[j][lineLocation])

				// There are already crates in that column, simply add a letter to the string
			} else {
				crates[byte(column)] += string(cratesString[j][lineLocation])
			}
		}
	}

	return crates, nil
}

// moveCrates moves a total of noCrates from the source column to the
// destination column, one by one.
func moveCrates(src, dest, noCrates byte, crates *map[byte]string) {

	cratesInColumn := byte(len((*crates)[src]))
	movedCrates := (*crates)[src][cratesInColumn-noCrates : cratesInColumn]
	(*crates)[src] = (*crates)[src][:cratesInColumn-noCrates]

	for i := byte(len(movedCrates) - 1); ; i-- {
		(*crates)[dest] += string(movedCrates[i])
		if i == 0 {
			break
		}
	}
}

// moveCratesSameOrder moves a total of noCrates from the source column to the
// destination column, all at a time.
func moveCratesSameOrder(src, dest, noCrates byte, crates *map[byte]string) {

	cratesInColumn := byte(len((*crates)[src]))
	movedCrates := (*crates)[src][cratesInColumn-noCrates : cratesInColumn]
	(*crates)[src] = (*crates)[src][:cratesInColumn-noCrates]

	(*crates)[dest] += movedCrates
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	cratesString, err := scanCrates(scanner)
	if err != nil {
		panic(err)
	}

	moves, err := scanMoves(scanner)
	if err != nil {
		panic(err)
	}

	crates, err := generateCratesStructure(cratesString)
	if err != nil {
		panic(err)
	}

	for _, move := range moves {
		//moveCrates(move[0], move[1], move[2], &crates)
		moveCratesSameOrder(move[0], move[1], move[2], &crates)
	}

	fmt.Printf("Crates after moving them: %v\n", crates)

	word := ""

	for i := 1; i <= 9; i++ {
		word += string(crates[byte(i)][len(crates[byte(i)])-1])
	}

	fmt.Println(word)
}
