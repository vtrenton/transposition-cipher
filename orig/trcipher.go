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

	populate(msg, grid, size)

	fmt.Println(dumpGrid(grid))

	fmt.Println(transposition(grid, size))
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

func populate(msg string, grid [][]rune, size int) {
	row := 0
	col := 0
	for i := 0; i < len(msg); i++ {
		grid[row][col] = rune(msg[i])
		col++
		if col >= size {
			row++
			col = 0
		}
		if row >= size {
			break
		}
	}
}

func transposition(grid [][]rune, size int) string {
	var ciphertext string
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			// invert the read
			if grid[col][row] != 0 {
				ciphertext += string(grid[col][row])
			}
		}
	}
	return ciphertext
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
