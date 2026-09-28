package main

import (
	"fmt"
	"strings"
)

func evaluate(s string, knowledge [][]string) string {

	last := -1
	kmap := make(map[string]string, len(knowledge))
	for _, v := range knowledge {
		kmap[v[0]] = v[1]
	}
	b := strings.Builder{}
	for idx := range len(s) {
		if s[idx] == '(' {
			last = idx + 1
		} else if s[idx] == ')' {
			word := s[last:idx]
			if _, ok := kmap[word]; ok {
				b.WriteString(kmap[word])
			} else {
				b.WriteByte('?')
			}
			last = -1
		} else if last == -1 {
			b.WriteByte(s[idx])
		}
	}

	return b.String()
}

func main() {
	fmt.Println(evaluate("(name)aa(age)bb", [][]string{{"name", "x"}, {"age", "10"}}))
	fmt.Println(evaluate("(name)aa(age)", [][]string{{"name", "x"}}))
}
