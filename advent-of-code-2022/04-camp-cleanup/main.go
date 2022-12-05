package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Range models an interval of integers, containing all integers between the
// "start" value or "end" value, including "start" and "end".
type Range struct {
	Start byte
	End   byte
}

// getRangesFromLine processes a line from the input file and returns the two
// ranges associated with that string.
func getRangesFromLine(line string) [2]Range {
	parts := strings.Split(line, ",")

	leftValues := strings.Split(parts[0], "-")
	rightValues := strings.Split(parts[1], "-")

	leftStart, _ := strconv.ParseUint(leftValues[0], 10, 8)
	leftEnd, _ := strconv.ParseUint(leftValues[1], 10, 8)
	rightStart, _ := strconv.ParseUint(rightValues[0], 10, 8)
	rightEnd, _ := strconv.ParseUint(rightValues[1], 10, 8)

	return [2]Range{
		{
			Start: byte(leftStart),
			End:   byte(leftEnd),
		},
		{
			Start: byte(rightStart),
			End:   byte(rightEnd),
		},
	}
}

// isContained returns true whether one of the two ranges is contained within
// the other.
func isContained(ranges [2]Range) bool {

	if ranges[0].Start < ranges[1].Start {
		return ranges[0].End >= ranges[1].End
	} else if ranges[0].Start > ranges[1].Start {
		return ranges[0].End <= ranges[1].End

		// If the starting point is equal for both ranges, one is always going to be
		// contained within the other.
	} else {
		return true
	}
}

// isContained returns true whether the two ranges are overlapping in at least
// one point.
func isOverlapping(ranges [2]Range) bool {

	// One of the ends of the second range must be between the ends of the first
	// range to overlap, or both ends of the second range must be out of the
	// bounds of the first range.
	return (ranges[0].Start <= ranges[1].Start && ranges[0].End >= ranges[1].Start) ||
		(ranges[0].Start <= ranges[1].End && ranges[0].End >= ranges[1].End) ||
		(ranges[0].Start >= ranges[1].Start && ranges[0].End <= ranges[1].End)

}

func computeRangesWithProperty(property func([2]Range) bool) (int32, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return int32(0), err
	}

	// Default scanner has endline splitter.
	scanner := bufio.NewScanner(f)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Convert the line into ranges
		ranges := getRangesFromLine(line)

		// Find out if one range is self-contained
		if property(ranges) {
			count += 1
		}
	}

	return int32(count), nil
}

func main() {
	countContained, err := computeRangesWithProperty(isContained)
	if err != nil {
		panic(err)
	}

	fmt.Printf("There are a total of %d pairs where one range is contained within the other.\n", countContained)

	countOverlapping, err := computeRangesWithProperty(isOverlapping)
	if err != nil {
		panic(err)
	}

	fmt.Printf("There are a total of %d pairs where ranges overlap.\n", countOverlapping)
}
