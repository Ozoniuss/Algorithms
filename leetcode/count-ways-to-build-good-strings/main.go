package main

import (
	"fmt"
)

type efficientQueue struct {
	elems map[int]int
}

func (e efficientQueue) append(x int) {
	e.elems[x]++
}

func (e efficientQueue) delete(x int) {
	if _, ok := e.elems[x]; ok {
		e.elems[x]--
		if e.elems[x] <= 0 {
			delete(e.elems, x)
		}
	}
}

func NewEfficientQueue() efficientQueue {
	return efficientQueue{
		elems: make(map[int]int, 10000),
	}
}

func lenq(e efficientQueue) int {
	return len(e.elems)
}

func top(e efficientQueue) int {
	if lenq(e) == 0 {
		return -1
	}
	var key int
	for k := range e.elems {
		key = k
		break
	}
	e.delete(key)
	return key
}

func countGoodStringsBfsEfficientQueue(low int, high int, zero int, one int) int {

	if low > high {
		return 0
	}

	q := NewEfficientQueue()
	q.append(zero)
	q.append(one)
	count := 0
	for lenq(q) != 0 {
		top := top(q)
		if top == -1 {
			panic("tf")
		}

		if top > high {
			continue
		}
		if top <= high {

			q.append(top + zero)
			q.append(top + one)
		}
		if top >= low {
			count++
		}
	}

	return count % (int(10e9 + 7))
}

func countGoodStringsBfsNormal(low int, high int, zero int, one int) int {

	if low > high {
		return 0
	}

	q := []int{zero, one}
	count := 0
	for len(q) != 0 {
		top := q[0]
		q = q[1:]

		if top > high {
			continue
		}
		if top <= high {

			q = append(q, top+zero)
			q = append(q, top+one)
		}
		if top >= low {
			count++
		}
	}

	return count % (int(10e9 + 7))
}

func countGoodStringsDfs(low int, high int, zero int, one int) int {
	if low > high {
		return 0
	}
	existing := make(map[int]int)
	return countDfs(low, high, zero, one, 0, existing)
}

func countDfs(low int, high int, zero int, one int, current int, existing map[int]int) int {
	if current > high {

		return 0
	}
	if val, ok := existing[current]; ok {
		return val
	}
	isSolution := 0
	if current >= low {
		isSolution = 1
	}

	total := (isSolution + countDfs(low, high, zero, one, current+zero, existing) + countDfs(low, high, zero, one, current+one, existing)) % (int(1e9) + 7)
	existing[current] = total
	return total
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func countGoodStringsDp(low int, high int, zero int, one int) int {
	if low > high {
		return 0
	}
	solutions := make([]int, high+1, high+1)
	solutions[0] = 1
	for i := min(zero, one); i <= high; i++ {
		if i >= zero {
			solutions[i] = (solutions[i] + solutions[i-zero]) % (int(1e9) + 7)
		}
		if i >= one {
			solutions[i] = (solutions[i] + solutions[i-one]) % (int(1e9) + 7)
		}
	}
	total := 0
	for t := low; t <= high; t++ {
		total = (total + solutions[t]) % (int(1e9) + 7)
	}
	return total
}

func main() {
	fmt.Println(countGoodStringsBfsNormal(100, 100, 10, 1))
	// fmt.Println(countGoodStringsBfsEfficientQueue(100, 100, 10, 1))
	fmt.Println(countGoodStringsDfs(100, 100, 10, 1))
	fmt.Println(countGoodStringsDp(100, 100, 10, 1))

}
