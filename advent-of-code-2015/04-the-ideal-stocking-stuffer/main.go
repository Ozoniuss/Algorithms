package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
)

// readPart reads the part number from the standard input.
func readPart() int {
	var part int
	fmt.Print("part: ")
	fmt.Scanf("%d\n", &part)
	return part
}

func checkInput(data []byte, in int) bool {
	inData := []byte(strconv.Itoa(in))
	data = append(data, inData...)

	md5enc := fmt.Sprintf("%x", md5.Sum(data))
	for i := 0; i < 6; i++ {
		if md5enc[i] != '0' {
			return false
		}
	}
	return true
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	part := readPart()
	if part == 1 {
		scanner.Scan()
		input := scanner.Text()
		data := []byte(input)

		i := 1
		for {
			if checkInput(data, i) {
				fmt.Println(i)
				break
			}
			i++
		}
	}
}
