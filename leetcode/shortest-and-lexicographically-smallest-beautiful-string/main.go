package main

import (
	"fmt"
	"math/bits"
)

func shortestBeautifulSubstring(s string, k int) string {
	if k < 1 || k > len(s) {
		panic("invalid k")
	}
	if len(s) == 0 {
		return ""
	}

	// total can be represented as two bytes
	total := [2]uint64{}
	ones := 0
	minval := [2]uint64{1<<64 - 1, 1<<64 - 1}
	stridx := [2]int{0, 0}

	l := 0
	r := 0
	for r < len(s) {
		// fmt.Println("iter", r, minval, l, r, total)
		if s[r] == '1' {
			// we need to start writing to total[0] in this case
			if (total[1]>>63 == 1 && total[0] == 0) || total[0] != 0 {
				l := total[1] >> 63
				total[0] = (total[0] << 1) | l
			}
			// and in either case, move the bits by 1
			total[1] = (total[1] << 1) | 1
			ones += 1
		} else {
			if (total[1]>>63 == 1 && total[0] == 0) || total[0] != 0 {
				l := total[1] >> 63
				total[0] = (total[0] << 1) | l
			}
			// same as above, but write 0 at the end and do not increase ones
			total[1] = total[1] << 1
		}

		// fmt.Printf("before shift, %b-%b (%d %d)\n", total[0], total[1], total[0], total[1])
		// fmt.Printf("total %064b_%064b %d\n ", total[0], total[1], ones)
		// no need to shift yet
		if ones < k {
			r++
			continue
		}

		// fmt.Println("ones", ones
		for s[l] == '0' || ones > k {
			if total[0] == 0 && total[1] == 0 {
				panic("total")
			}
			// no need to move further and will also make sure we stop at 1
			if s[l] == '1' {
				// we really only need to adjust when a 1 gets removed, otherwise
				// total stays the same.
				ones -= 1
				if total[0] != 0 {
					total[0] -= uint64(1) << (bits.Len64(total[0]) - 1)
				} else {
					total[1] -= uint64(1) << (bits.Len64(total[1]) - 1)
				}
			}
			l++
		}
		// fmt.Printf("after shift %064b_%064b %d\n ", total[0], total[1], ones)
		// fmt.Printf("after shift, %b-%b (%d %d)\n", total[0], total[1], total[0], total[1])

		if ones == k {
			// fmt.Printf("comparison, total, minval, %b-%b %b-%b\n", total[0], total[1], minval[0], minval[1])
			if total[0] < minval[0] {
				minval[0] = total[0]
				minval[1] = total[1]
				stridx[0] = l
				stridx[1] = r
			} else if total[0] == minval[0] && total[1] < minval[1] {
				minval[1] = total[1]
				stridx[0] = l
				stridx[1] = r
			}
			// fmt.Printf("l, r after comparison total minval  %b-%b %b-%b\n", total[0], total[1], minval[0], minval[1])
		}

		r++
	}
	if ones < k {
		return ""
	}
	if minval[1] == 0 && minval[0] == 0 {
		panic("should be caught by k")
	}
	// fmt.Println("l, r", l, r, s)
	return s[stridx[0] : stridx[1]+1]
}

func main() {
	fmt.Println(shortestBeautifulSubstring("000101", 2))
	fmt.Println(shortestBeautifulSubstring("10101", 2))
	fmt.Println(shortestBeautifulSubstring("10101", 3))
	fmt.Println(shortestBeautifulSubstring("100011001", 3))
	fmt.Println(shortestBeautifulSubstring("1011", 2))
	fmt.Println(shortestBeautifulSubstring("101101", 3))
	fmt.Println(shortestBeautifulSubstring("0111001110101100101111111101001001001101000000011111101010000000100100001000000000110101010110", 38))
}
