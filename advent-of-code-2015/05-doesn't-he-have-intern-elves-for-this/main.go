package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// isNice returns whether the input string is nice or not, based on the provided
// conditions that determine it.
func isNice(input string, conds ...func(string) bool) bool {
	out := true
	for _, cond := range conds {
		out = out && cond(input)
	}
	return out
}

// hasInvalidLetterCombos returns whether the input string has the combination
// "ab", "cd", "pq", "xy".
func hasInvalidLetterCombos(input string) bool {
	return strings.Contains(input, "ab") ||
		strings.Contains(input, "cd") ||
		strings.Contains(input, "pq") ||
		strings.Contains(input, "xy")
}

// hasThreeVowels returns whether the string has at least three vowels.
func hasThreeVowels(input string) bool {
	vowels := "aeiou"
	count := 0
	for _, letter := range input {
		if strings.Contains(vowels, string(letter)) {
			count += 1
		}
	}
	return count >= 3
}

// hasDoubleLetter returns whether the string has at least one occurence of
// double letters, such as "aa".
func hasDoubleLetter(input string) bool {
	if len(input) <= 1 {
		return false
	}
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			return true
		}
	}
	return false
}

// hasRepeatingTwoLetters returns whether the input string contains a pair of
// any two letters that appears at least twice in the string without
// overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa
// (aa, but it overlaps)
func hasRepeatingTwoLetters(input string) bool {

	if len(input) < 4 {
		return false
	}

	// Stores the encountered pairs of consecutive letters, as well as their
	// offsets. In order for the condition to be satisfied, the offsets between
	// the repeating pairs must differ by at least two, which means they don't
	// overlap.
	repeatingPairs := make(map[string]int)

	for i := 0; i < len(input)-1; i++ {
		pair := input[i : i+2]
		if idx, ok := repeatingPairs[pair]; ok {
			if abs(idx-i) >= 2 {
				return true
			} else {
				// Set the offset to -69. If the repeating pair is encountered
				// at least 3 times, it's impossible that we don't find two
				// locations that don't overlap.
				//
				// Note: there are actually 69 nice strings :0 I swear I didn't
				// know that before.
				repeatingPairs[pair] = -69
			}
		} else {
			repeatingPairs[pair] = i
		}
	}
	return false
}

// hasLetterInBetween returns whether the input contains at least one letter
// which repeats with exactly one letter between them, like xyx, abcdefeghi
// (efe), or even aaa.
func hasLetterInBetween(input string) bool {
	if len(input) < 3 {
		return false
	}
	for i := 0; i < len(input)-2; i++ {
		if input[i] == input[i+2] {
			return true
		}
	}
	return false
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
		niceCount := 0

		for scanner.Scan() {
			line := scanner.Text()
			if isNice(line,
				func(s string) bool { return !hasInvalidLetterCombos(s) },
				hasDoubleLetter,
				hasThreeVowels) {
				niceCount++
			}
		}
		fmt.Printf("There are a total of %d nice strings.", niceCount)
	} else if part == 2 {
		niceCount := 0

		for scanner.Scan() {
			line := scanner.Text()
			if isNice(line,
				hasLetterInBetween,
				hasRepeatingTwoLetters) {
				niceCount++
			}
		}
		fmt.Printf("There are a total of %d nice strings.", niceCount)
	}
}
