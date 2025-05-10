package main

import "testing"

func TestFunc(t *testing.T) {
	t.Run("remove all the spaces from the output using strip", func(t *testing.T) {
		got := strip("this is a test")
		want := "thisisatest"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("test the power of a number", func(t *testing.T) {
		got := pow(5, 2)
		want := 25

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestGrid(t *testing.T) {
	var g Grid
	g.msg = strip("attack the north wall") // strip is tested above

	t.Run("validate the rows and columns match size", func(t *testing.T) {
		g.newGrid()
		got := g.size
		want := 5 // want squared is the total size of the slice.

		if got != want {
			t.Errorf("sizes dont match, got %d, want %d", got, want)
		}
	})

	t.Run("validate grid population", func(t *testing.T) {

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
