package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func roundTrip(value string) bool {
	return value == leveledToString(stringToLeveled(value))
}

func TestLeveled() {
	values := []string{"4", "[4]", "[4,4]", "[[4]]", "[[4],4]", "[[4],[4]]", "[4,[4,[4]]]", "[]", "[[]]"}
	for _, v := range values {
		if !roundTrip(v) {
			panic(fmt.Sprintf("Test failed for value %s", v))
		}
	}
}

type element int

// leveledArray can either be an element (a level 0 array) or an array of
// leveledArrays.
type leveledArray struct {
	element int
	array   []leveledArray
}

// leveledToString converts a leveled array to its string form.
func leveledToString(l leveledArray) string {
	if l.array == nil {
		return fmt.Sprint(l.element)
	} else {
		out := "["
		for _, a := range l.array {
			out += leveledToString(a) + ","
		}
		out = strings.TrimSuffix(out, ",")
		out += "]"
		return out
	}
}

// stringToLeveled converts a leveled array represented as a string into a
// leveledArray object. It is assumed that the string represents a valid
// array.
func stringToLeveled(leveledString string) leveledArray {

	// The string represents an array, level 0 array (an element). This
	// always means the last parentheses of the string is also ']'.
	if leveledString[0] == '[' {
		// Empty array special case.
		if leveledString[1] == ']' {
			return leveledArray{
				array: make([]leveledArray, 0, 0),
			}
		}
		positions := [][2]int{}
		start := 1
		// Open bracket counter. It is increased by 1 if an open bracket is
		// found, and decreased by 1 if a closed bracket is found. Whenever
		// this is 0, we know that the string we're currently parsing from the
		// top level items is an element (level 0 array), otherwise when this
		// bracket gets closed we know we just read an array item.
		inArray := 0

		// The idea here is, call this function recursively for all top-level
		// elements of the current string representation, whether they are
		// level 0 arrays or higher level arrays. What this does is determine
		// a list of starting and ending positions: each such pairs of positions
		// denotes the starting and ending point of the string representation
		// of a top-level item inside the provided array, and we will call this
		// function recursively to determine the leveled array objects each of
		// the items represent in order to add them as items to the current
		// leveled array.
		for i := 1; i < len(leveledString)-1; i++ {
			if leveledString[i] == '[' {
				inArray += 1
			} else if leveledString[i] == ']' {
				inArray -= 1
			} else if inArray == 0 {
				// Bracket counter is 0, we either read the string
				// representation of an element or an array.
				if leveledString[i] == ',' {
					// Add the positions of that string to the positions list.
					positions = append(positions, [2]int{start, i})
					start = i + 1
				}
			}
		}
		// The first-to-last character is the ending position of the last
		// top-level item.
		positions = append(positions, [2]int{start, len(leveledString) - 1})
		arr := leveledArray{
			array: []leveledArray{},
		}
		// Convert all the top-level items from string to leveledArray.
		for _, p := range positions {
			arr.array = append(arr.array, stringToLeveled(leveledString[p[0]:p[1]]))
		}

		return arr

		// If the current string is an element, just return its string
		// representation.
	} else {
		val, _ := strconv.Atoi(leveledString)
		return leveledArray{
			element: val,
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// compare compares two leved arrays as described in the statement, and returns
// true if they are in the right order.
func compare(l1, l2 leveledArray) int {

	// Both leveled arrays are in fact elements.
	if l1.array == nil && l2.array == nil {
		if l1.element == l2.element {
			return 0
		} else if l1.element < l2.element {
			return -1
		} else {
			return 1
		}

		// Both are non-element leveled arrays
	} else if l1.array != nil && l2.array != nil {
		t1 := len(l1.array)
		t2 := len(l2.array)

		if t1 == 0 && t2 == 0 {
			return 0
		}

		for i := 0; i < min(t1, t2); i++ {
			if val := compare(l1.array[i], l2.array[i]); val != 0 {
				return val
			}
		}

		if t1 == t2 {
			return 0
		} else if t1 < t2 {
			return -1
		} else {
			return 1
		}

	} else if l1.array == nil {
		arr := leveledArray{
			array: []leveledArray{
				{
					element: l1.element,
				},
			},
		}
		return compare(arr, l2)
	} else {
		arr := leveledArray{
			array: []leveledArray{
				{
					element: l2.element,
				},
			},
		}
		return compare(l1, arr)
	}
}

func main() {

	TestLeveled()

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	part := 2

	if part == 1 {
		idx := 1
		rightOrder := []int{}
		out := ""
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				leveledArrayStrings := strings.Split(out, "\n")
				l1 := stringToLeveled(leveledArrayStrings[0])
				l2 := stringToLeveled(leveledArrayStrings[1])
				if compare(l1, l2) == -1 {
					rightOrder = append(rightOrder, idx)
				}
				idx += 1
				out = ""
			} else {
				out += line + "\n"
			}
		}

		fmt.Printf("The arrays that are in the \"right order\" are at positions %+v\n", rightOrder)

		sum := 0
		for i := 0; i < len(rightOrder); i++ {
			sum += rightOrder[i]
		}

		fmt.Printf("Their sum is %d\n", sum)

	} else if part == 2 {
		all := make([]leveledArray, 0)
		for scanner.Scan() {
			line := scanner.Text()
			if line != "" {
				all = append(all, stringToLeveled(line))
			}
		}

		all = append(all, leveledArray{
			array: []leveledArray{
				{
					array: []leveledArray{
						{element: 2},
					},
				},
			},
		},
			leveledArray{
				array: []leveledArray{
					{
						array: []leveledArray{
							{element: 6},
						},
					},
				},
			},
		)

		for i := 0; i < len(all)-1; i++ {
			for j := i; j < len(all); j++ {
				if compare(all[i], all[j]) == 1 {
					aux := all[i]
					all[i] = all[j]
					all[j] = aux
				}
			}
		}

		decoders := [2]int{}
		for idx, larr := range all {
			fmt.Printf("%d:%s\n", idx+1, leveledToString(larr))
			if leveledToString(larr) == "[[2]]" {
				decoders[0] = idx + 1
			}
			if leveledToString(larr) == "[[6]]" {
				decoders[1] = idx + 1
			}
		}

		fmt.Printf("Decoder positions: %v\n", decoders)

	}
}
