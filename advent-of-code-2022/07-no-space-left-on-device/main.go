package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// processLine reads the current line and either:
// - changes the working directory;
// - increases the size of the working directory in the "directories" map
// (which contains the sizes of all directories); if the current line is a file
// with a specific size ( "<number> <name>" )
// - increases the size of the parent directory with the size of the current
// directory, if the current line is "cd .."
// - nothing, otherwise; if the current line is "dir <name>" or "$ ls"
//
// directories is a map storing the size of all different directories, including
// subdirectories.
// For example, if directory a also contains b, the map will store both the size
// of b (with /a/b as the key) as well as the size of a, which is greater or
// equal to the size of b (with /a as the key)
func processLine(line string, wd *[]string, directories *map[string]int) {

	// list command
	if line == "$ ls" {
		return
	}

	// showing a directory
	if line[:3] == "dir" {
		return
	}

	/* From this point all other commands take one argument. */

	parts := strings.Split(line, " ")

	// Change directory
	if parts[0] == "$" && parts[1] == "cd" {

		// Go to root directory
		// It is worth noting that for all other directories, cd only specifies
		// a relative path, so no argument will start with / after the first
		// one.
		if parts[2] == "/" {

			// Initialize the value of the outermost directory.
			wd = &[]string{}
			(*directories)[dirToStr(*wd)] = 0

			// go inside a directory
		} else if parts[2] != ".." {
			(*wd) = append((*wd), parts[2])

			// initialize the total size of this directory*
			(*directories)[dirToStr(*wd)] = 0

			// move outside a directory
		} else {

			dirSize := (*directories)[dirToStr(*wd)]

			// remove the last directory from the stack
			*wd = (*wd)[:len(*wd)-1]

			// Add the size of that directory to the counter
			(*directories)[dirToStr(*wd)] += dirSize
		}

		//Last remaining case if "<size> <filename>"
	} else {

		size, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			panic(err)
		}
		(*directories)[dirToStr(*wd)] += int(size)
	}

}

// dirToStr converts a directory from its array form to its string form.
func dirToStr(dir []string) string {

	// Root directory
	if len(dir) == 0 {
		return "/"
	}

	out := ""
	for _, d := range dir {
		out += fmt.Sprintf("/%s", d)
	}

	return out
}

// strToDir converts a directory from its string form to its array form.
func strToDir(dir string) []string {

	// Root directory
	if dir == "/" {
		return []string{}
	}

	// Remove beginning slash
	dir = dir[1:]
	return strings.Split(dir, "/")
}

// findSmallestDir finds the directory of smallest size that needs to be
// deleted in order to free up "enough space".
func findSmallestDir(directories *map[string]int) (string, int) {
	diskSize := 70_000_000
	needed := 30_000_000
	current := (*directories)["/"]

	// Maximum space allowed
	spaceRequired := current - (diskSize - needed)

	fmt.Println(spaceRequired)

	if spaceRequired < 0 {
		panic("why would you do this")
	}

	minSize := current
	dir := "/"

	for d, size := range *directories {
		if size > spaceRequired {
			if size < minSize {
				dir = d
				minSize = size
			}
		}
	}

	return dir, minSize

}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	directories := make(map[string]int)
	wd := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		processLine(line, &wd, &directories)
	}

	// Treat the case where the stupid example doesn't exit from the last
	// directory, which would have been nice.
	if len(wd) != 0 {

		// Add up all remaining sizes recursively
		for j := len(wd) - 1; j >= 0; j-- {
			size := directories[dirToStr(wd)]
			wd := wd[:j]
			directories[dirToStr(wd)] += size
		}
	}

	total := 0

	for _, size := range directories {
		if size < 100_000 {
			total += size
		}
	}

	dir, size := findSmallestDir(&directories)

	fmt.Printf("Total size of directories with size smaller than 100_000: %d\n", total)
	fmt.Printf("Directory that can be deleted is %s and has size: %d\n", dir, size)
}
