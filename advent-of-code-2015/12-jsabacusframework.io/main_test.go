package main

import (
	"bufio"
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"testing"
)

var result = 0

func BenchmarkPart1(b *testing.B) {
	var r int
	var f *os.File
	var err error
	for n := 0; n <= b.N; n++ {

		// ignore file reads
		b.StopTimer()
		f, err = os.Open("input.txt")
		if err != nil {
			panic(err)
		}
		b.StartTimer()

		b := bufio.NewReader(f)
		var c byte
		err = nil
		inNumber := false

		sb := &strings.Builder{}
		total := 0

		// assume ascii encoded
		for ; err == nil; c, err = b.ReadByte() {
			if !inNumber && isNumberStart(c) {
				inNumber = true
				sb.WriteByte(c)
			} else if inNumber && isNumberStart(c) {
				sb.WriteByte(c)
			} else if inNumber && !isNumberStart(c) {
				inNumber = false
				number := sb.String()
				sb.Reset()

				numberValue, err := strconv.Atoi(number)
				if err != nil {
					panic(err)
				}
				total += numberValue
			}
		}
		r = total
	}
	result = r
	f.Close()
}

func BenchmarkPart2(b *testing.B) {
	var r int
	var err error
	var f *os.File
	for n := 0; n <= b.N; n++ {
		f, err = os.Open("input.txt")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		var alldata map[string]any
		dec := json.NewDecoder(f)
		err = dec.Decode(&alldata)
		if err != nil {
			panic(err)
		}

		r = computeNodeSum(alldata)
	}
	result = r
	f.Close()
}
