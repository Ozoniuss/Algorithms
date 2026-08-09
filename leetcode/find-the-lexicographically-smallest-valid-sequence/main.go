package main

import (
	"fmt"
)

func main() {
	// fmt.Println(validSequenceIdentical("xcavbcbcaabew", "bcbca", 0, 0, 3, cache))
	// fmt.Println(validSequenceIdentical("xcavbcbcaabew", "bcbca", 0, 0, 3, cache))
	// fmt.Println(validSequenceIdentical("xcavbcbcaabew", "bcbca", 0, 0, 3, cache))

	fmt.Println(validSequence("vbcaa", "abc"))
	fmt.Println(validSequence("bacdc", "abc"))
	fmt.Println(validSequence("aaaaaa", "aaabc"))
	fmt.Println(validSequence("abc", "ab"))

}

func validSequence(word1 string, word2 string) []int {

	maxSuffix := make([]int, len(word1))
	c := 0

	j := len(word2) - 1
	for i := len(word1) - 1; i >= 0; i -= 1 {
		maxSuffix[i] = c
		if j != -1 && word1[i] == word2[j] {
			c += 1
			j -= 1
		}
	}

	j = 0
	seq := []int{}
	hasSwitched := false
	for i := 0; i < len(word1); i++ {
		if word1[i] == word2[j] {
			seq = append(seq, i)
			j += 1
		} else if !hasSwitched && maxSuffix[i]+len(seq)+1 >= len(word2) {
			// eager switch
			fmt.Println("switch at", i)
			hasSwitched = true
			j += 1
			seq = append(seq, i)
		}
		if len(seq) == len(word2) {
			break
		}
	}
	if len(seq) < len(word2) {
		return []int{}
	}
	return seq
}
