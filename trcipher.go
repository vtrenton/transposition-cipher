package main

import (
	"fmt"
	"math"
)

func main() {
	msg := "attack the north wall"

	// determine the size of the square (grid)
	// based on the input method
	var sizef float64
	for int(math.Pow(sizef, 2)) < len(msg) {
		sizef++
	}
	size := int(sizef)

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
