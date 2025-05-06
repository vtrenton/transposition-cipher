package main

import "testing"

func TestGrid(t *testing.T) {
	var g Grid
	t.Run("test creation of new grid", func(t *testing.T) {
		g.newGrid()

		t.Logf("%v\n", g.dumpGrid())
	})
}

func TestFunc(t *testing.T) {
	t.Run("remove all the spaces from the output using strip", func(t *testing.T){
		got := strip("this is a test")
		want := "thisisatest"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
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
