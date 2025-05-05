package main

import "testing"

func TestGrid(t *testing.T) {
	var g Grid
	t.Run("test creation of new grid", func(t *testing.T) {
		g.newGrid()

		t.Logf("%v\n", g.dumpGrid())
	})
}

func (g *Grid) dumpGrid() string {
	var outstr string
	for _, r := range g.grid {
		for _, c := range r {
			outstr += string(c)
		}
	}
	return outstr
}
