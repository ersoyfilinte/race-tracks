package main

import (
	"testing"
)

func TestLabyrinth(t *testing.T) {
	var c = Case{
		grid: struct {
			width, height int
		}{
			width:  29,
			height: 29,
		},
		points: Path{
			start: Position{
				x: 0,
				y: 0,
			},
			end: Position{
				x: 28,
				y: 28,
			},
		},
	}

	c.obstacles = append(c.obstacles, Path{
		start: Position{
			x: 1,
			y: 0,
		},
		end: Position{
			x: 1,
			y: 27,
		},
	}, Path{
		start: Position{
			x: 3,
			y: 1,
		},
		end: Position{
			x: 3,
			y: 28,
		},
	}, Path{
		start: Position{
			x: 5,
			y: 0,
		},
		end: Position{
			x: 5,
			y: 27,
		},
	}, Path{
		start: Position{
			x: 7,
			y: 1,
		},
		end: Position{
			x: 7,
			y: 28,
		},
	}, Path{
		start: Position{
			x: 9,
			y: 0,
		},
		end: Position{
			x: 9,
			y: 27,
		},
	}, Path{
		start: Position{
			x: 11,
			y: 1,
		},
		end: Position{
			x: 11,
			y: 28,
		},
	}, Path{
		start: Position{
			x: 13,
			y: 0,
		},
		end: Position{
			x: 13,
			y: 27,
		},
	}, Path{
		start: Position{
			x: 15,
			y: 1,
		},
		end: Position{
			x: 15,
			y: 28,
		},
	}, Path{
		start: Position{
			x: 17,
			y: 0,
		},
		end: Position{
			x: 17,
			y: 27,
		},
	}, Path{
		start: Position{
			x: 19,
			y: 1,
		},
		end: Position{
			x: 19,
			y: 28,
		},
	}, Path{
		start: Position{
			x: 21,
			y: 0,
		},
		end: Position{
			x: 21,
			y: 27,
		},
	}, Path{
		start: Position{
			x: 23,
			y: 1,
		},
		end: Position{
			x: 23,
			y: 28,
		},
	}, Path{
		start: Position{
			x: 25,
			y: 0,
		},
		end: Position{
			x: 25,
			y: 27,
		},
	}, Path{
		start: Position{
			x: 27,
			y: 1,
		},
		end: Position{
			x: 27,
			y: 28,
		},
	})

	if s, err := c.Calc(); err || s != 23 {
		t.Errorf("Expected 23 do not match actual %d.", s)
	}
}

func TestReversedLabyrinth(t *testing.T) {
	var c = Case{
		grid: struct {
			width, height int
		}{
			width:  29,
			height: 29,
		},
		points: Path{
			start: Position{
				x: 28,
				y: 28,
			},
			end: Position{
				x: 0,
				y: 0,
			},
		},
	}

	c.obstacles = append(c.obstacles, Path{
		start: Position{
			x: 1,
			y: 0,
		},
		end: Position{
			x: 1,
			y: 27,
		},
	}, Path{
		start: Position{
			x: 3,
			y: 1,
		},
		end: Position{
			x: 3,
			y: 28,
		},
	}, Path{
		start: Position{
			x: 5,
			y: 0,
		},
		end: Position{
			x: 5,
			y: 27,
		},
	}, Path{
		start: Position{
			x: 7,
			y: 1,
		},
		end: Position{
			x: 7,
			y: 28,
		},
	}, Path{
		start: Position{
			x: 9,
			y: 0,
		},
		end: Position{
			x: 9,
			y: 27,
		},
	}, Path{
		start: Position{
			x: 11,
			y: 1,
		},
		end: Position{
			x: 11,
			y: 28,
		},
	}, Path{
		start: Position{
			x: 13,
			y: 0,
		},
		end: Position{
			x: 13,
			y: 27,
		},
	}, Path{
		start: Position{
			x: 15,
			y: 1,
		},
		end: Position{
			x: 15,
			y: 28,
		},
	}, Path{
		start: Position{
			x: 17,
			y: 0,
		},
		end: Position{
			x: 17,
			y: 27,
		},
	}, Path{
		start: Position{
			x: 19,
			y: 1,
		},
		end: Position{
			x: 19,
			y: 28,
		},
	}, Path{
		start: Position{
			x: 21,
			y: 0,
		},
		end: Position{
			x: 21,
			y: 27,
		},
	}, Path{
		start: Position{
			x: 23,
			y: 1,
		},
		end: Position{
			x: 23,
			y: 28,
		},
	}, Path{
		start: Position{
			x: 25,
			y: 0,
		},
		end: Position{
			x: 25,
			y: 27,
		},
	}, Path{
		start: Position{
			x: 27,
			y: 1,
		},
		end: Position{
			x: 27,
			y: 28,
		},
	})

	if s, err := c.Calc(); err || s != 23 {
		t.Errorf("Expected 23 do not match actual %d.", s)
	}
}

func TestUnsolvable(t *testing.T) {
	var c = Case{
		grid: struct {
			width, height int
		}{
			width:  30,
			height: 30,
		},
		points: Path{
			start: Position{
				x: 0,
				y: 0,
			},
			end: Position{
				x: 29,
				y: 29,
			},
		},
	}

	c.obstacles = append(c.obstacles, Path{
		start: Position{
			x: 26,
			y: 26,
		},
		end: Position{
			x: 28,
			y: 29,
		},
	}, Path{
		start: Position{
			x: 29,
			y: 26,
		},
		end: Position{
			x: 29,
			y: 28,
		},
	})

	if s, err := c.Calc(); !err {
		t.Errorf("Expected no solution do not match actual %d.", s)
	}
}

func TestSameStartAndEndPoint(t *testing.T) {
	var c = Case{
		grid: struct {
			width, height int
		}{
			width:  30,
			height: 30,
		},
		points: Path{
			start: Position{
				x: 0,
				y: 0,
			},
			end: Position{
				x: 0,
				y: 0,
			},
		},
	}

	if s, err := c.Calc(); err || s != 0 {
		t.Errorf("Expected 0 do not match actual %d.", s)
	}
}

func TestStartPointOutside(t *testing.T) {
	var c = Case{
		grid: struct {
			width, height int
		}{
			width:  30,
			height: 30,
		},
		points: Path{
			start: Position{
				x: 30,
				y: 30,
			},
			end: Position{
				x: 0,
				y: 0,
			},
		},
	}

	if s, err := c.Calc(); !err {
		t.Errorf("Expected no solution do not match actual %d.", s)
	}
}

func TestEndPointOutside(t *testing.T) {
	var c = Case{
		grid: struct {
			width, height int
		}{
			width:  30,
			height: 30,
		},
		points: Path{
			start: Position{
				x: 29,
				y: 29,
			},
			end: Position{
				x: -1,
				y: -1,
			},
		},
	}

	if s, err := c.Calc(); !err {
		t.Errorf("Expected no solution do not match actual %d.", s)
	}
}
