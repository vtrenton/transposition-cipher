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
	for row := 0; row < size; row++ {
		// get the range of the substring
		start := row * size
		end := (row + 1) * size

		// dont allow read out of bounds
		var strblock string
		if start >= len(msg) {
			break
		} else if end >= len(msg) {
			strblock = string(msg[start:])
		} else {
			strblock = string(msg[start:end])
		}

		// add padding so there is no index out of bounds
		if len(strblock) < size {
			diff := size - len(strblock)
			for i := 0; i < diff; i++ {
				strblock += " "
			}
		}

		for col := 0; col < size; col++ {
			grid[row][col] = rune(strblock[col])

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
