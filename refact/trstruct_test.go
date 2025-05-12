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
		g.populate()

		got := g.grid[3][2]
		want := 'l'

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("validate chars are correctly placed in grid", func(t *testing.T) {
		dump := g.dumpGrid()
		padlen := 7
		got := dump[:len(dump)-padlen] // trim off padding
		want := "attackthenorthwall"

		//	for _, i := range got {
		//		t.Log(i)
		//	}

		//	for _, i := range want {
		//		t.Log(i)
		//	}

		if got != want {
			t.Errorf("got %s, want %s - this is a test", got, want)
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
