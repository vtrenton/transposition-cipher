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
	row := 0
	col := 0
	for i := 0; i < len(g.msg); i++ {
		g.grid[row][col] = rune(g.msg[i])
		col++
		if col >= g.size {
			row++
			col = 0
		}
		if row >= g.size {
			break
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
