package main

import "strings"

var builder strings.Builder

func longestCommonPrefixTwo(s1, s2 string) string {

	if len(s2) < len(s1) {
		s1, s2 = s2, s1
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			builder.WriteByte(s1[i])
		} else {
			break
		}
	}
	val := builder.String()
	builder.Reset()
	return val
}

func longestCommonPrefixTwoNewBuilder(s1, s2 string) string {

	if len(s2) < len(s1) {
		s1, s2 = s2, s1
	}

	var sb = strings.Builder{}

	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			sb.WriteByte(s1[i])
		} else {
			break
		}
	}
	val := sb.String()
	return val
}

func longestCommonPrefixTwoStringAddition(s1, s2 string) string {

	if len(s2) < len(s1) {
		s1, s2 = s2, s1
	}

	var val string

	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			builder.WriteByte(s1[i])
			val = val + string(s1[i])
		} else {
			break
		}
	}
	return val
}

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	if len(strs) == 2 {
		return longestCommonPrefixTwo(strs[0], strs[1])
	}

	// optimization to use same array
	prefix := longestCommonPrefixTwo(strs[0], strs[1])
	strs[1] = prefix
	strs = strs[1:]

	return longestCommonPrefix(strs)
}

// If I were a dork, I'd say use *strategy* pattern for benchamrking easier
func longestCommonPrefixB(strs []string, compute func(string, string) string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	if len(strs) == 2 {
		return compute(strs[0], strs[1])
	}

	// optimization to use same array
	prefix := compute(strs[0], strs[1])
	strs[1] = prefix
	strs = strs[1:]

	return longestCommonPrefixB(strs, compute)
}

func main() {

}
