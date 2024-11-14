package main

import "fmt"

type Grid struct {
	grid [][]rune
	size int
}

// methods
func (g *Grid) populate(msg string) {
	var row int
	var col int
	for i := range len(msg) {
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

func (g *Grid) transposition() string {
	var ciphertext string
	for row := range g.size {
		for col := range g.size {
			if g.grid[col][row] != 0 {
				ciphertext += string(g.grid[col][row])
			}
		}
	}
	return ciphertext
}

// DEBUG: useful function for debugging but not essential for program functionality
//func (g *Grid) dumpGrid() string {
//	var outstr string
//	for _, r := range g.grid {
//		for _, c := range r {
//			outstr += string(c)
//		}
//	}
//	return outstr
//}

func (g *Grid) newGrid(msg string) {
	var size int
	for pow(size, 2) < len(msg) {
		size++
	}

	grid := make([][]rune, size)
	for i := range grid {
		grid[i] = make([]rune, size)
	}
	g.grid = grid
	g.size = size
}

// functions

// Not going to bring in math for pow
// write my own helper function
func pow(base int, exp int) int {
	result := 1
	for range exp {
		result *= base
	}
	return result
}

// yet another avoided import ;)
func strip(msg string) string {
	var msgSlice []rune
	for _, v := range msg {
		if v == ' ' {
			continue
		}
		msgSlice = append(msgSlice, v)
	}
	return string(msgSlice)
}

func main() {
	var g Grid
	msg := "attack the north wall"
	fmtMsg := strip(msg) //remove any spaces
	g.newGrid(fmtMsg)
	g.populate(fmtMsg)
	fmt.Println(g.transposition())
}
