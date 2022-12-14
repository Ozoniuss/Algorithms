package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func roundTrip(value string) bool {
	return value == stringLeveled(fromLeveled(value))
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

func stringLeveled(l leveledArray) string {

	if l.array == nil {
		return fmt.Sprint(l.element)
	} else {
		out := "["
		for _, a := range l.array {
			out += stringLeveled(a) + ","
		}
		out = strings.TrimSuffix(out, ",")
		out += "]"
		return out
	}
}

// fromLeveled takes as input a string which could be a possible representation
// of a leveled array and turns it into a leveled array.
func fromLeveled(l string) leveledArray {

	// this always means the last parentheses is also ']'
	if l[0] == '[' {
		if l[1] == ']' {
			return leveledArray{
				array: make([]leveledArray, 0, 0),
			}
		}
		positions := [][2]int{}
		start := 1
		inArray := 0
		for i := 1; i < len(l)-1; i++ {
			if l[i] == '[' {
				inArray += 1
			} else if l[i] == ']' {
				inArray -= 1
			} else if inArray == 0 {
				if l[i] == ',' {
					positions = append(positions, [2]int{start, i})
					start = i + 1
				}
			}
		}
		positions = append(positions, [2]int{start, len(l) - 1})
		arr := leveledArray{
			array: []leveledArray{},
		}
		for _, p := range positions {
			arr.array = append(arr.array, fromLeveled(l[p[0]:p[1]]))
		}

		return arr

		// The only case left is number
	} else {
		val, _ := strconv.Atoi(l)
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
				l1 := fromLeveled(leveledArrayStrings[0])
				l2 := fromLeveled(leveledArrayStrings[1])
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
				all = append(all, fromLeveled(line))
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
			fmt.Printf("%d:%s\n", idx+1, stringLeveled(larr))
			if stringLeveled(larr) == "[[2]]" {
				decoders[0] = idx + 1
			}
			if stringLeveled(larr) == "[[6]]" {
				decoders[1] = idx + 1
			}
		}

		fmt.Printf("Decoder positions: %v\n", decoders)

	}
}
