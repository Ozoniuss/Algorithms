package main

import (
	"fmt"
)

// readPart reads the part number from the standard input.
func readPart() int {
	var part int
	fmt.Print("part: ")
	fmt.Scanf("%d\n", &part)
	return part
}

func hasStraight(pass string) bool {
	if len(pass) < 3 {
		return false
	}

	for i := 2; i < len(pass); i++ {
		if pass[i]-pass[i-1] == 1 && pass[i-1]-pass[i-2] == 1 {
			return true
		}
	}
	return false
}

func hasForbidden(pass string) bool {
	for i := 0; i < len(pass); i++ {
		if pass[i] == 'i' || pass[i] == 'o' || pass[i] == 'l' {
			return true
		}
	}
	return false
}

func hasPairs(pass string) bool {
	if len(pass) < 4 {
		return false
	}
	p1end := -1
	for i := 1; i < len(pass); i++ {
		if pass[i] == pass[i-1] && p1end == i-1 {
			// only possible cases: aaabx, aaabb, aaaa
			p1end = i - 1
		} else if pass[i] == pass[i-1] && p1end == -1 {
			p1end = i
		} else if pass[i] == pass[i-1] && p1end != i-1 {
			return true
		}
	}
	return false
}

func notHasForbidden(pass string) bool {
	return !hasForbidden(pass)
}

// checkCriteria checks if a password is valid
func checkCriteria(pass string, criteria ...func(string) bool) bool {
	ret := true
	for _, c := range criteria {
		ret = ret && c(pass)
	}
	return ret
}

func validPassword(pass string) bool {
	return checkCriteria(pass, notHasForbidden, hasStraight, hasPairs)
}

func nextLetter(l byte) byte {
	return (l+1-'a')%26 + 'a'
}

func nextPassword(pass string) string {
	p := []byte(pass)
	for i := len(p) - 1; i > 0; i-- {
		p[i] = nextLetter(p[i])
		if p[i] != 'a' {
			break
		}
	}
	return string(p)
}

func main() {
	password := "hxbxwxba"
	for {
		password = nextPassword(password)
		if validPassword(password) {
			fmt.Println(password)
			break
		}
	}

	// I figured out part 2 without having to run any code at all through
	// simple reasoning.
}
