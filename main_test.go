package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	var te Test

	te = Test{
		size: struct {
			width, height uint
		}{
			width:  5,
			height: 5,
		},
		point: Route{
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

	te.obstacles = append(te.obstacles, Obstacle{
		point: Route{
			start: Position{
				x: 1,
				y: 4,
			},
			end: Position{
				x: 2,
				y: 3,
			},
		},
	})

	r, s := te.Calc()

	if !r || s != 3 {
		t.Fatal("Optimal solution of 1. test case must be 3.")
	}

	te = Test{
		size: struct {
			width, height uint
		}{
			width:  3,
			height: 3,
		},
		point: Route{
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

	te.obstacles = append(te.obstacles, Obstacle{
		point: Route{
			start: Position{
				x: 1,
				y: 1,
			},
			end: Position{
				x: 0,
				y: 2,
			},
		},
	}, Obstacle{
		point: Route{
			start: Position{
				x: 0,
				y: 2,
			},
			end: Position{
				x: 1,
				y: 1,
			},
		},
	})

	r, s = te.Calc()

	if !r || s != 3 {
		t.Fatal("Optimal solution of 2. test case must be 3.")
	}
}
