package main

import "strings"

func wordPattern(pattern string, s string) bool {

	parts := strings.Split(s, " ")

	if len(parts) != len(pattern) {
		return false
	}
	words := make(map[byte]string)
	usedWords := make(map[string]struct{})
	for i := 0; i < len(pattern); i++ {
		if _, ok := words[pattern[i]]; !ok {
			if _, ok := usedWords[parts[i]]; ok {
				return false
			}
			words[pattern[i]] = parts[i]
			usedWords[parts[i]] = struct{}{}
		} else {
			if words[pattern[i]] != parts[i] {
				return false
			}
		}
	}
	return true
}

func main() {

}
