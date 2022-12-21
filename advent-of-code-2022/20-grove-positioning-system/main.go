package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mod(a, b int) int {

	if a%b == 0 {
		return 0
	}

	if a >= 0 {
		return a % b
	} else {
		return b + a%b
	}
}

// move executes a single move on the list of numbers.
func move(numbers *[][2]int, num [2]int) {

	if num[0] == 0 {
		return
	}

	length := len(*numbers)

	// Find the position of the number.
	pos := 0
	for (*numbers)[pos] != num {
		pos++
	}

	newPos := 0

	if num[0] < 0 {
		newPos = mod(pos+num[0], length-1)
	} else if num[0] > 0 {
		newPos = mod(pos+num[0], length-1)
	}

	if newPos == pos {
		return
	} else if newPos > pos {
		for i := pos; i < newPos; i++ {
			(*numbers)[i] = (*numbers)[(i + 1)]
		}
		(*numbers)[newPos] = num
		return
	} else {
		for i := pos; i > newPos; i-- {
			(*numbers)[i] = (*numbers)[(i - 1)]
		}
		(*numbers)[newPos] = num
		return
	}
}

type DLL struct {
	head   *DLLNode
	length int
}

func newDLL() DLL {
	return DLL{
		head:   nil,
		length: 0,
	}
}

type DLLNode struct {
	value [2]int
	next  *DLLNode
	prev  *DLLNode
}

func newDLLNode(val [2]int, next *DLLNode, prev *DLLNode) DLLNode {
	return DLLNode{
		value: val,
		next:  next,
		prev:  prev,
	}
}

/*
	Note that for the functions below, it only matters that we swap the values.
*/

// advance moves a node in a DLL forwards.
func (n *DLLNode) advance() {
	aux := n.value
	n.value = n.next.value
	n.next.value = aux
}

// regress moves a node backwards in a DLL.
func (n *DLLNode) regress() {
	aux := n.value
	n.value = n.prev.value
	n.prev.value = aux
}

func (d *DLL) print() {
	test := d.head
	vals := []int{}
	for i := 0; i < d.length; i++ {
		vals = append(vals, test.value[0])
		test = test.next
	}
	fmt.Println(vals)
}

// moveDLL moves a number in a cyclic DLL with a number of steps indicated by
// its value.
func moveDLL(dll *DLL, num [2]int) {
	//fmt.Println(num)
	if num[0] == 0 {
		return
	}

	pos := dll.head
	for pos.value != num {
		pos = pos.next
	}

	if num[0] < 0 {
		for i := 0; i > num[0]%(dll.length-1); i-- {
			pos.regress()
			pos = pos.prev
		}
	} else if num[0] > 0 {
		for i := 0; i < num[0]%(dll.length-1); i++ {
			pos.advance()
			pos = pos.next
		}
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	numbers := [][2]int{}
	i := 0

	dll := newDLL()
	head := newDLLNode([2]int{0, 0}, nil, nil)
	dll.head = &head
	current := dll.head

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, [2]int{num, i})

		if i == 0 {
			dll.head.value = [2]int{num, i}
			dll.length++
		} else {
			new := newDLLNode([2]int{num, i}, nil, current)
			current.next = &new
			current = &new
			dll.length++
		}
		i++
	}

	current.next = dll.head
	dll.head.prev = current

	part := 2

	if part == 1 {
		for _, n := range numbers {
			moveDLL(&dll, n)
		}

		pos := dll.head
		for pos.value[0] != 0 {
			pos = pos.next
		}

		for i := 0; i < len(numbers); i++ {
			pos = pos.next
		}

		vals := [3]int{}

		//fmt.Println("here")

		for i := 0; i < 3000; i++ {
			pos = pos.next
			if i == 999 {
				vals[0] = pos.value[0]
			}
			if i == 1999 {
				vals[1] = pos.value[0]
			}
			if i == 2999 {
				vals[2] = pos.value[0]
			}
		}

		fmt.Println(vals)

		// pos := 0
		// for numbers[pos][0] != 0 {
		// 	pos++
		// }

		// l := len(numbers)
		// fmt.Println(l)
		// fmt.Println(pos)

		// fmt.Println(numbers[(pos+1000)%l][0], numbers[(pos+2000)%l][0], numbers[(pos+3000)%l][0])
	} else {
		const DECRYPTION_KEY = 811589153

		for i := 0; i < len(numbers); i++ {
			numbers[i][0] = DECRYPTION_KEY * numbers[i][0]
		}

		current := dll.head
		for i := 0; i < dll.length; i++ {
			current.value[0] = DECRYPTION_KEY * current.value[0]
			current = current.next
		}

		for i := 0; i < 10; i++ {
			for _, n := range numbers {
				moveDLL(&dll, n)
			}
		}

		pos := dll.head
		for pos.value[0] != 0 {
			pos = pos.next
		}

		for i := 0; i < len(numbers); i++ {
			pos = pos.next
		}

		vals := [3]int{}

		//fmt.Println("here")

		for i := 0; i < 3000; i++ {
			pos = pos.next
			if i == 999 {
				vals[0] = pos.value[0]
			}
			if i == 1999 {
				vals[1] = pos.value[0]
			}
			if i == 2999 {
				vals[2] = pos.value[0]
			}
		}

		fmt.Println(vals)
		fmt.Println(vals[0] + vals[1] + vals[2])

	}

}
