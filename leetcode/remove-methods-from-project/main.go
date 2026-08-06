package main

import "fmt"

func remainingMethods(n int, k int, invocations [][]int) []int {
	g := make(map[int][]int, 0)
	for _, i := range invocations {
		s := i[0]
		d := i[1]
		if _, ok := g[s]; !ok {
			g[s] = []int{d}
		} else {
			g[s] = append(g[s], d)
		}
	}

	fmt.Println("g", g)

	q := []int{k}
	corrupted := make(map[int]struct{}, 0)
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		if _, ok := corrupted[cur]; ok {
			continue
		}

		for _, n := range g[cur] {
			if _, ok := corrupted[n]; !ok {
				q = append(q, n)
			}
		}

		corrupted[cur] = struct{}{}
	}

LOOP:
	for k := range n {
		if _, ok := corrupted[k]; ok {
			continue
		}
		for _, n := range g[k] {
			if _, ok := corrupted[n]; ok {
				corrupted = map[int]struct{}{}
				break LOOP
			}
		}
	}

	remaining := []int{}
	for k := range n {
		if _, ok := corrupted[k]; ok {
			continue
		}
		remaining = append(remaining, k)
	}
	return remaining
}

func main() {
	a := remainingMethods(5, 0, [][]int{{1, 2}, {0, 2}, {0, 1}, {3, 4}})
	fmt.Println(a)
}
