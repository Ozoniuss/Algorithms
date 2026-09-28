package main

func maxDepth(s string) int {
	cnt := 0
	m := 0
	for _, c := range s {
		if c == '(' {
			cnt += 1
		}
		if c == ')' {
			cnt -= 1
		}
		m = max(cnt, m)
	}
	return m
}

func main() {

}
