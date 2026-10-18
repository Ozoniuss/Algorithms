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

	f, err := os.Open("in_level-4_2-large.txt")
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
		diststrs := strings.Split(numstrs[0], ",")
		distx, _ := strconv.Atoi(diststrs[0])
		disty, _ := strconv.Atoi(diststrs[1])

		seqx := minseq(distx, distx < 0)
		seqy := minseq(disty, disty < 0)
		seqstrx := fmt.Sprintf("%v", seqx)
		seqstry := fmt.Sprintf("%v", seqy)
		fmt.Println(seqstrx[1 : len(seqstrx)-1])
		fmt.Println(seqstry[1 : len(seqstry)-1])
		fmt.Println()
	}
}
