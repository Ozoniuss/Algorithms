package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func manhattanDistance(x, y [2]int) int {
	return abs(x[0]-y[0]) + abs(x[1]-y[1])
}

// processLine takes a line as the input string, finds the sensor and beacon
// positions, and adds them to the sensors and beacons map. In addition to all
// sensors, the sensors map stores the maximum distance visible by the sensor.
func processLine(line string, sensors *map[[2]int]int, beacons *map[[2]int]struct{}) {
	r := regexp.MustCompile("-?[0-9]+")
	positions := r.FindAllString(line, 4)

	xs, _ := strconv.Atoi(positions[0])
	ys, _ := strconv.Atoi(positions[1])
	xb, _ := strconv.Atoi(positions[2])
	yb, _ := strconv.Atoi(positions[3])

	sensor := [2]int{xs, ys}
	beacon := [2]int{xb, yb}

	(*sensors)[sensor] = manhattanDistance(sensor, beacon)
	(*beacons)[beacon] = struct{}{}
}

// canBeBeacon returns whether a beacon can occupy the provided position.
func canBeBeacon(pos [2]int, sensors *map[[2]int]int, beacons *map[[2]int]struct{}) bool {

	// A beacon is already at pos.
	if _, ok := (*beacons)[pos]; ok {
		return true
	}

	for s, viewDistance := range *sensors {

		// This would mean that there is a closer beacon to senzor s.
		// The less or equal is a sneaky detail from the authors.
		if manhattanDistance(pos, s) <= viewDistance {
			return false
		}
	}
	return true
}

// canBeUndiscoveredBeacon returns whether a beacon can occupy the provided
// position, with the condition that the beacon had not been discovered already.
func canBeUndiscoveredBeacon(pos [2]int, sensors *map[[2]int]int, beacons *map[[2]int]struct{}) bool {

	// A beacon is already at pos.
	if _, ok := (*beacons)[pos]; ok {
		return false
	}

	for s, viewDistance := range *sensors {

		// This would mean that there is a closer beacon to senzor s.
		// The less or equal is a sneaky detail from the authors.
		if manhattanDistance(pos, s) <= viewDistance {
			return false
		}
	}
	return true
}

func isInBound(point [2]int, xminlimit, xmaxlimit, yminlimit, ymaxlimit int) bool {
	return point[0] >= xminlimit &&
		point[0] <= xmaxlimit &&
		point[1] >= yminlimit &&
		point[1] <= ymaxlimit
}

// getSensorOuterLayer appends the points representing the edges of a rectangle
// with smallest area that contains all the area visible by a sensor to a map.
// The point's coordinates values must not exceed the provided limits.
func getSensorOuterLayer(sensor [2]int, fov int, xminlimit, xmaxlimit, yminlimit, ymaxlimit int, outer *map[[2]int]struct{}) {

	// height of the sensor
	y := sensor[1]

	// closest points on the x axis that are not reached by the sensor.
	left := sensor[0] - fov - 1
	right := sensor[0] + fov + 1

	for i := 0; i <= fov+1; i++ {

		up := [2]int{left + i, y - i}
		down := [2]int{left + i, y + i}

		if isInBound(up, xminlimit, xmaxlimit, yminlimit, ymaxlimit) {
			(*outer)[up] = struct{}{}
		}

		if isInBound(down, xminlimit, xmaxlimit, yminlimit, ymaxlimit) {
			(*outer)[down] = struct{}{}
		}

	}
	for i := 0; i <= fov+1; i++ {
		up := [2]int{right - i, y - i}
		down := [2]int{right - i, y + i}

		if isInBound(up, xminlimit, xmaxlimit, yminlimit, ymaxlimit) {
			(*outer)[up] = struct{}{}
		}

		if isInBound(down, xminlimit, xmaxlimit, yminlimit, ymaxlimit) {
			(*outer)[down] = struct{}{}
		}
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	sensors := make(map[[2]int]int, 0)
	beacons := make(map[[2]int]struct{}, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		processLine(line, &sensors, &beacons)
	}

	part := 2

	if part == 1 {
		y := 2000000
		count := 0

		// This can be heavily optimised by computing a rectangle that bounds the
		// senzors and the max viewing distance, but it's not worth it.

		// Alternatively, one can project each sensor on the y line and find the
		// edges of the interval it covers.
		for x := -12000000; x < 16000000; x++ {
			if !canBeBeacon([2]int{x, y}, &sensors, &beacons) {
				count++
			}
		}

		fmt.Printf("%d positions cannot contain a beacon at line y=%d\n", count, y)

	} else if part == 2 {

		// This is just way to slow, it won't do it.
		// for x := 0; x <= 4000000; x++ {
		// 	fmt.Println(x)
		// 	for y := 0; y <= 4000000; y++ {
		// 		if canBeUndiscoveredBeacon([2]int{x, y}, &sensors, &beacons) {
		// 			fmt.Printf("Beacon undiscovered at position %d, %d\n", x, y)
		// 			fmt.Printf("Distress signal is %d\n", x*4000000+y)
		// 		}
		// 	}
		// }

		// Instead, we are going to walk around each sensor, on the rectangle
		// representing the outer bound of their field of view. The beacon is
		// supposed to be there, knowing that it's only going to be one.
		outer := make(map[[2]int]struct{})
		for s, fov := range sensors {
			getSensorOuterLayer(s, fov, 0, 4000000, 0, 4000000, &outer)
		}

		fmt.Println(len(outer))

		for pos := range outer {
			if canBeUndiscoveredBeacon(pos, &sensors, &beacons) {
				fmt.Printf("Beacon undiscovered at position %d, %d\n", pos[0], pos[1])
				fmt.Printf("Distress signal is %d\n", pos[0]*4000000+pos[1])
			}
		}

	}

}
