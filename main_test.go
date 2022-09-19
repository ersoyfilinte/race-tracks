package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	var c Case

	c = Case{
		grid: struct {
			width, height int
		}{
			width:  5,
			height: 5,
		},
		points: Path{
			start: Position{
				x: 4,
				y: 0,
			},
			end: Position{
				x: 4,
				y: 4,
			},
		},
	}

	c.obstacles = append(c.obstacles, Path{
		start: Position{
			x: 1,
			y: 3,
		},
		end: Position{
			x: 2,
			y: 4,
		},
	})

	s, err := c.Calc()

	if err || s != 3 {
		t.Fatal("Optimal solution of 1. test case must be 3.")
	}

	c = Case{
		grid: struct {
			width, height int
		}{
			width:  3,
			height: 3,
		},
		points: Path{
			start: Position{
				x: 0,
				y: 0,
			},
			end: Position{
				x: 2,
				y: 2,
			},
		},
	}

	c.obstacles = append(c.obstacles, Path{
		start: Position{
			x: 0,
			y: 1,
		},
		end: Position{
			x: 1,
			y: 2,
		},
	})

	s, err = c.Calc()

	if err || s != 3 {
		t.Fatal("Optimal solution of 2. test case must be 3.")
	}
}
