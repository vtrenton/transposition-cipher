package main

import "fmt"

type Grid struct {
	grid [][]rune
	size int
	msg  string
}

// constructor method
func (g Grid) new(msg string) Grid {
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
	return Grid{
		grid: grid,
		size: size,
		msg:  msg,
	}
}

func (g Grid) populate() {
	for row := 0; row < g.size; row++ {
		// get the range of the substring
		start := row * g.size
		end := (row + 1) * g.size

		// dont allow read out of bounds
		var strblock string
		if start >= len(g.msg) {
			break
		} else if end >= len(g.msg) {
			strblock = string(g.msg[start:])
		} else {
			strblock = string(g.msg[start:end])
		}

		// add padding so there is no index out of bounds
		if len(strblock) < g.size {
			diff := g.size - len(strblock)
			for i := 0; i < diff; i++ {
				strblock += " "
			}
		}

		for col := 0; col < g.size; col++ {
			g.grid[row][col] = rune(strblock[col])
		}
	}
}

func (g Grid) transposition() string {
	var ciphertext string
	for row := 0; row < g.size; row++ {
		for col := 0; col < g.size; col++ {
			// invert the read
			if g.grid[col][row] != 0 {
				ciphertext += string(g.grid[col][row])
			}
		}
	}
	return ciphertext
}

func main() {
	msg := "attack the north wall"
	var g Grid
	g = g.new(msg)
	g.populate()
	fmt.Println(g.transposition())
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
