package main

import "fmt"

type Grid struct {
	grid [][]rune
	size int
}

// methods
func (g Grid) populate(msg string) {
	row := 0
	col := 0
	for i := 0; i < len(msg); i++ {
		g.grid[row][col] = rune(msg[i])
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

// functions
func newGrid(msg string) Grid {
	var size int
	for pow(size, 2) < len(msg) {
		size++
	}

	grid := make([][]rune, size)
	for i := range grid {
		grid[i] = make([]rune, size)
	}

	return Grid{
		grid: grid,
		size: size,
	}
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

func main() {
	msg := "attack the north wall"
	g := newGrid(msg)
	g.populate(msg)
	fmt.Println(g.transposition())
}
