package main

import (
	"fmt"
	"math/rand"
	"os"
)

const filesize = 100000000

func main() {
	f, err := os.OpenFile(fmt.Sprintf("numsgo%d.txt", filesize), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	rand.Seed(3)

	var buf = make([]byte, filesize, filesize)
	for i := 0; i < filesize; i++ {
		num := rand.Int() % 10
		buf[i] = byte(num) + '0'
	}

	f.Write(buf)
	f.WriteString("\n")
	for i := 0; i < filesize; i++ {
		num := rand.Int() % 10
		buf[i] = byte(num) + '0'
	}
	f.Write(buf)
}
