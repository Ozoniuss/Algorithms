package main

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	x1 := rec1[0]
	x2 := rec1[2]
	if x1 > x2 {
		x1, x2 = x2, x1
	}

	y1 := rec1[1]
	y2 := rec1[3]
	if y1 > y2 {
		y1, y2 = y2, y1
	}

	m1 := rec2[0]
	m2 := rec2[2]

	if m1 > m2 {
		m1, m2 = m2, m1
	}

	n1 := rec2[1]
	n2 := rec2[3]

	if n1 > n2 {
		n1, n2 = n2, n1
	}

	if m1 >= x2 || m2 <= x1 {
		return false

	}

	if y1 >= n2 || y2 <= n1 {
		return false
	}

	return true

}

func main() {}
