package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OuterCube struct {
	A [3]int
	B [3]int
	C [3]int
	D [3]int
	E [3]int
	F [3]int
	G [3]int
	H [3]int
}

// processLine adds a cube from a line to the set of cubes
func processLine(line string, cubes *map[[3]int]struct{}) {
	coords := strings.Split(line, ",")

	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	z, _ := strconv.Atoi(coords[2])

	(*cubes)[[3]int{x, y, z}] = struct{}{}
}

// disconnectedFaces returns the number of cube faces that are not directly
// connected to any other cube
func disconnectedFaces(cube [3]int, cubes *map[[3]int]struct{}) int {
	counter := 0

	if _, ok := (*cubes)[[3]int{cube[0] - 1, cube[1], cube[2]}]; !ok {
		counter++
	}
	if _, ok := (*cubes)[[3]int{cube[0] + 1, cube[1], cube[2]}]; !ok {
		counter++
	}
	if _, ok := (*cubes)[[3]int{cube[0], cube[1] - 1, cube[2]}]; !ok {
		counter++
	}
	if _, ok := (*cubes)[[3]int{cube[0], cube[1] + 1, cube[2]}]; !ok {
		counter++
	}
	if _, ok := (*cubes)[[3]int{cube[0], cube[1], cube[2] - 1}]; !ok {
		counter++
	}
	if _, ok := (*cubes)[[3]int{cube[0], cube[1], cube[2] + 1}]; !ok {
		counter++
	}
	return counter
}

// computeExtremes returns the coordinates of a parallelipipedic object that
// contains the cube within.
func computeExtremes(cubes *map[[3]int]struct{}) OuterCube {
	min_x := 100
	min_y := 100
	min_z := 100

	max_x := -100
	max_y := -100
	max_z := -100

	for c := range *cubes {
		if c[0] > max_x {
			max_x = c[0]
		}
		if c[1] > max_y {
			max_y = c[1]
		}
		if c[2] > max_z {
			max_z = c[2]
		}
		if c[0] < min_x {
			min_x = c[0]
		}
		if c[1] < min_y {
			min_y = c[1]
		}
		if c[2] < min_z {
			min_z = c[2]
		}
	}

	return OuterCube{
		A: [3]int{min_x - 1, min_y - 1, min_z - 1},
		B: [3]int{max_x + 1, min_y - 1, min_z - 1},
		C: [3]int{max_x + 1, min_y - 1, max_z + 1},
		D: [3]int{min_x - 1, min_y - 1, max_z + 1},
		E: [3]int{min_x - 1, max_y + 1, min_z - 1},
		F: [3]int{max_x + 1, max_y + 1, min_z - 1},
		G: [3]int{max_x + 1, max_y + 1, max_z + 1},
		H: [3]int{min_x - 1, max_y + 1, max_z + 1},
	}
}

func isOutside(coords [3]int, outer OuterCube) bool {
	if coords[0] < outer.A[0] || coords[0] > outer.B[0] {
		return true
	}
	if coords[1] < outer.A[1] || coords[1] > outer.E[1] {
		return true
	}
	if coords[2] < outer.A[2] || coords[2] > outer.D[2] {
		return true
	}
	return false
}

func bfs(outer OuterCube, cubes *map[[3]int]struct{}) int {
	visited := make(map[[3]int]struct{})
	q := [][3]int{
		outer.A, outer.B, outer.C, outer.D, outer.E, outer.F, outer.G, outer.H,
	}

	for _, el := range q {
		visited[el] = struct{}{}
	}

	counter := 0

	idx := 0

	for len(q) > 0 {
		current := q[0]

		// In this case, when faces touch it means that gas reached
		// faces of the cube. So for disconnected faces, it means that
		// they didn't actually reach the cube yet.
		//
		// Also, keep in mind that there is a unique intersection
		// between any two cubes, if any.

		if 6-disconnectedFaces(current, cubes) > 0 {
			fmt.Println(current, 6-disconnectedFaces(current, cubes))
		}
		counter += (6 - disconnectedFaces(current, cubes))

		q = q[1:]

		for _, n := range getNeighbours(current) {
			_, inVisited := visited[n]
			_, inCube := (*cubes)[n]
			if !inVisited && !inCube {
				if !isOutside(n, outer) {
					visited[n] = struct{}{}
					q = append(q, n)
				}
			}
		}

		idx += 1
	}

	return counter

}

func getNeighbours(cube [3]int) [][3]int {

	neighbourds := [][3]int{}

	neighbourds = append(neighbourds, [3]int{cube[0] - 1, cube[1], cube[2]})
	neighbourds = append(neighbourds, [3]int{cube[0] + 1, cube[1], cube[2]})
	neighbourds = append(neighbourds, [3]int{cube[0], cube[1] - 1, cube[2]})
	neighbourds = append(neighbourds, [3]int{cube[0], cube[1] + 1, cube[2]})
	neighbourds = append(neighbourds, [3]int{cube[0], cube[1], cube[2] - 1})
	neighbourds = append(neighbourds, [3]int{cube[0], cube[1], cube[2] + 1})

	return neighbourds
}

func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	cubes := make(map[[3]int]struct{}, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		processLine(line, &cubes)
	}

	counter := 0
	for c := range cubes {
		counter += disconnectedFaces(c, &cubes)
	}

	fmt.Println(counter)

	counterExterior := bfs(computeExtremes(&cubes), &cubes)
	fmt.Println(counterExterior)

}
