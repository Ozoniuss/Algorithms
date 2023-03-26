package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// readInput reads from standard input until newline.
func readInput() string {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	return s.Text()
}

func main() {
	line := readInput()
	lineTrimmed := strings.TrimLeft(line, "0")
	fmt.Println(lineTrimmed)
}
