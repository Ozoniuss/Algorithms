package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func minseq(dist int, invert bool) []int {
	// ugly hack
	if invert {
		dist = -dist
	}
	seq := []int{}
	seq = append(seq, 0)

	padlen := dist / 2
	if padlen > 5 {
		padlen = 5
	}

	for i := range padlen {
		seq = append(seq, 5-i)
	}

	for range dist - 2*padlen {
		seq = append(seq, max(5-padlen, 1))
	}

	for i := range padlen {
		seq = append(seq, 5-padlen+i+1)
	}

	seq = append(seq, 0)

	if invert {
		for i := range len(seq) {
			seq[i] = -seq[i]
		}
	}

	return seq
}

func main() {

	// for i := range 20 {
	// 	fmt.Println(minseq(i))
	// }

	f, err := os.Open("in_level-3_2-large.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)

	// scan header
	s.Scan()
	for s.Scan() {
		line := s.Text()
		// fmt.Println(line)
		numstrs := strings.Split(line, " ")
		dist, _ := strconv.Atoi(numstrs[0])
		seq := minseq(dist, dist < 0)
		seqstr := fmt.Sprintf("%v", seq)
		fmt.Println(seqstr[1 : len(seqstr)-1])
	}
}
