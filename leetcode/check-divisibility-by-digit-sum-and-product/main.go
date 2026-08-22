package main

func main() {

}

func checkDivisibility(n int) bool {
	s := 0
	p := 1
	k := n
	for n > 0 {
		s += n % 10
		p *= n % 10
		n /= 10
	}
	return k%(s+p) == 0
}
