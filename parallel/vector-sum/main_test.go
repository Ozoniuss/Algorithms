package main

import (
	"fmt"
	"os"
	"sync"
	"testing"
)

const SIZE = 5000000

var FILE = fmt.Sprintf("nums%d.txt", SIZE)

// readNums reads the numbers from the given file in two arrays. For simplicity,
// all numbers in the file are digits; they are not separated by any separator.
//
// For example, if the file's content is
// 12345
// 67890
// Then arr1 would have elements 1,2,3,4,5 and arr2 would have elements 6,7,8,
// 9,0.
func readNums() (arr1 [SIZE]int32, arr2 [SIZE]int32) {

	f, err := os.Open(FILE)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, SIZE, SIZE)
	f.Read(buf)
	for i := 0; i < len(buf); i++ {
		// Since the file is UTF-8 encoded, decrease the ASCII value of the 0
		// character to obtain the digit value.
		arr1[i] = int32(buf[i] - '0')
	}

	// Skip Linux newline to get the next array of digits (the two digit sets
	// are separated via '\n')
	f.Seek(2, 1)

	f.Read(buf)
	for i := 0; i < len(buf); i++ {
		arr2[i] = int32(buf[i] - '0')
	}
	return
}

// ComputeClassic1 adds the numbers of arr1 and arr2 into arr3. The sum of the
// elements with the same index i in arr1 and arr2 is placed in the ith location
// of arr3.
// i.e. arr3[i] = arr1[i] + arr2[i]
func ComputeClassic1(arr1, arr2, arr3 *[SIZE]int32) {
	for i := 0; i < SIZE; i++ {
		(arr3)[i] = ((arr2)[i] + (arr1)[i]) % 10
	}
}

// computeParallel1 does the same as ComputeClassic1, except it separates
// the array position into blocks, executing the computation of the elements
// in each block in parallel.
func ComputeParallel1(arr1, arr2, arr3 *[SIZE]int32, nrThreads int) {
	w := &sync.WaitGroup{}
	w.Add(nrThreads)
	for i := 0; i < nrThreads; i++ {
		go computeParallel1Thread(arr1, arr2, arr3, i, nrThreads, w)
	}
	w.Wait()
}

// ComputeParallel1Thread does the computation of a single thread for the
// function ComputeParallel1.
//
// Note that a mutex is not required for the arrays since any two blocks are
// disjoint.
func computeParallel1Thread(arr1, arr2, arr3 *[SIZE]int32, threadIdx, nrThreads int, w *sync.WaitGroup) {
	start := (SIZE / nrThreads) * threadIdx
	end := (SIZE / nrThreads) * (threadIdx + 1)
	for i := start; i < end; i++ {
		arr3[i] = (arr2[i] + arr1[i]) % 10
	}
	w.Done()
}

func BenchmarkComputeClassic1(b *testing.B) {
	b.StopTimer()
	arr1, arr2 := readNums()
	var arr3 [SIZE]int32
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		ComputeClassic1(&arr1, &arr2, &arr3)
	}
}

func BenchmarkComputeParallel1(b *testing.B) {
	b.StopTimer()
	arr1, arr2 := readNums()
	var arr3 [SIZE]int32
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		ComputeParallel1(&arr1, &arr2, &arr3, 8)
	}
}
