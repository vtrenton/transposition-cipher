package main

import "fmt"

func main() {
	msg := "attack the north wall"

	// determine the size of the square (grid)
	// based on the input method
	var size int
	for pow(size, 2) < len(msg) {
		size++
	}

	// initialize a multidementional fixed-sized slice
	grid := make([][]rune, size)
	for i := range grid {
		grid[i] = make([]rune, size)
	}

	// populate with junk data
	for ri := 0; ri < size; ri++ {
		for ci := 0; ci < size; ci++ {
			grid[ri][ci] = 'x'
		}
	}

	fmt.Println(dumpGrid(grid))

}

func dumpGrid(grid [][]rune) string {
	var outstr string
	for _, r := range grid {
		for _, c := range r {
			outstr += string(c)
		}
	}
	return outstr
}

// Not going to bring in math for pow
// write my own helper function
func pow(base int, exp int) int {
	result := 1
	for i := 1; i <= exp; i++ {
		result *= base
	}
	return result
}
