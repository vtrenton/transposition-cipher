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
	t.Run("test creation of new grid", func(t *testing.T) {
		g.newGrid()

		got := g.dumpGrid()
		want := ""

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("add a test message and validate the rows and columns match size", func(t *testing.T) {
		g.msg = strip("attack the north wall") // strip is tested above

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
