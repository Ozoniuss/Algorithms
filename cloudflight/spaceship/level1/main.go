package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, err := os.Open("in_level-1_2-large.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)

	// scan header
	s.Scan()
	for s.Scan() {
		line := s.Text()
		// fmt.Println(line)
		total := 0
		numstrs := strings.Split(line, " ")
		for _, nstr := range numstrs {
			num, _ := strconv.Atoi(nstr)
			total += num
		}
		fmt.Println(total)
	}
}
