package main

import (
	"fmt"
	"math"
)

type Test struct {
	size struct {
		width, height uint
	}

	point     Route
	obstacles []Obstacle
}

type Route struct {
	start, end Position
}

type Position struct {
	x, y uint
}

type Obstacle struct {
	point Route
}

type Possibilities struct {
	hops []Hop
}

type Hop struct {
	speed Velocity
	point Position
}

type Velocity struct {
	x, y int
}

func (o *Obstacle) Contains(p Position) bool {
	return o.point.start.x <= p.x && p.x <= o.point.start.x && o.point.start.y <= p.y && p.y <= o.point.start.y
}

func (ps *Possibilities) Contains(h Hop) bool {
	for _, ho := range ps.hops {
		if ho == h {
			return true
		}
	}

	return false
}

func (t *Test) Calc() (bool, uint) {
	var ps Possibilities
	var r, err bool
	var s, n uint

	ps.hops = append(ps.hops, Hop{})

	// Loop the steps.
	for !r {
		s++

		// Loop previous hops.
		for _, h := range ps.hops {
			// Loop the possible X of speed.
			for x := h.speed.x - 1; !r && x <= h.speed.x+1 && x >= -3 && x <= 3; x++ {
				// Loop the possible Y of speed.
				for y := h.speed.y - 1; !r && y <= h.speed.y+1 && y >= -3 && y <= 3; y++ {
					// Check if the point is on the grid.
					if int(h.point.x)+x >= 0 && int(h.point.x)+x < int(t.size.width) && int(h.point.y)+y >= 0 && int(h.point.y)+y < int(t.size.height) {
						ho := Hop{Velocity{x, y}, Position{uint(int(h.point.x) + x), uint(int(h.point.y) + y)}}

						// Check if previous hops contain current hop.
						if !ps.Contains(ho) {
							err = false

							// Check if the obstacles contain current point.
							for _, o := range t.obstacles {
								for x := math.Min(float64(o.point.start.x), float64(o.point.end.x)); !err && x <= math.Max(float64(o.point.start.x), float64(o.point.end.x)); x++ {
									for y := math.Min(float64(o.point.start.y), float64(o.point.end.y)); !err && y <= math.Max(float64(o.point.start.y), float64(o.point.end.y)); y++ {
										if ho.point == (Position{uint(x), uint(y)}) {
											err = true
										}
									}
								}
							}

							if !err {
								r = ho.point == t.point.end

								ps.hops = append(ps.hops, ho)
							}
						}
					}
				}
			}

			if r {
				break
			}
		}

		// Check the current length of the hops equal to the length of the previous hops.
		err = n == uint(len(ps.hops))

		if err {
			break
		}

		n = uint(len(ps.hops))
	}

	return r, s
}

func main() {
	var n uint

	err := true

	// Ask for number of test cases and validate.
	for err {
		n = scan("Number of test cases: ")

		err = n < 1

		if err {
			fmt.Println("Number of test cases must be greater than or equal to 1.")
		}
	}

	var ts []Test

	// Loop the number of test cases and ask properties.
	for i := uint(1); i <= n; i++ {
		var t Test

		err = true

		// Ask and validate width of test case's grid.
		for err {
			t.size.width = scan(fmt.Sprintf("Width of %d. test case's grid: ", i))

			err = t.size.width < 1 || t.size.width > 30

			if err {
				fmt.Printf("Width of %d. test case's grid must be greater than or equal to 1 and less than or equal to 30.\n", i)
			}
		}

		err = true

		// Ask and validate height of test case's grid.
		for err {
			t.size.height = scan(fmt.Sprintf("Height of %d. test case's grid: ", i))

			err = t.size.height < 1 || t.size.height > 30

			if err {
				fmt.Printf("Height of %d. test case's grid must be greater than or equal to 1 and less than or equal to 30.\n", i)
			}
		}

		err = true

		// Ask and validate the X of test case's start point.
		for err {
			t.point.start.x = scan(fmt.Sprintf("X of %d. test case's start point: ", i))

			err = t.point.start.x >= t.size.width

			if err {
				fmt.Printf("X of %d. test case's start point must be greater than or equal to 0 and less than %d.\n", i, t.size.width)
			}
		}

		err = true

		// Ask and validate the Y of test case's start point.
		for err {
			t.point.start.y = scan(fmt.Sprintf("Y of %d. test case's start point: ", i))

			err = t.point.start.y >= t.size.height

			if err {
				fmt.Printf("Y of %d. test case's start point must be greater than or equal to 0 and less than %d.\n", i, t.size.height)
			}
		}

		err = true

		// Ask and validate the X of test case's end point.
		for err {
			t.point.end.x = scan(fmt.Sprintf("X of %d. test case's end point: ", i))

			err = t.point.end.x >= t.size.width

			if err {
				fmt.Printf("X of %d. test case's end point must be greater than or equal to 0 and less than %d.\n", i, t.size.width)
			}
		}

		err = true

		// Ask and validate the Y of test case's end point.
		for err {
			t.point.end.y = scan(fmt.Sprintf("Y of %d. test case's end point: ", i))

			err = t.point.end.y >= t.size.height

			if err {
				fmt.Printf("Y of %d. test case's end point must be greater than or equal to 0 and less than %d.\n", i, t.size.height)
			}
		}

		on := scan(fmt.Sprintf("Number of %d. grid's obstacle: ", i))

		// Loop the number of obstacles and ask properties.
		for oi := uint(1); oi <= on; oi++ {
			err = true

			for err {
				var o Obstacle

				// Ask and validate the X of test case's obstacle's start point.
				for err {
					o.point.start.x = scan(fmt.Sprintf("X of %d. test case's %d. obstacle's start point: ", i, oi))

					err = o.point.start.x >= t.size.width

					if err {
						fmt.Printf("X of %d. test case's %d. obstacle's start point must be greater than or equal to 0 and less than %d.\n", i, oi, t.size.width)
					}
				}

				err = true

				// Ask and validate the Y of test case's obstacle's start point.
				for err {
					o.point.start.y = scan(fmt.Sprintf("Y of %d. test case's %d. obstacle's start point: ", i, oi))

					err = o.point.start.y >= t.size.height

					if err {
						fmt.Printf("Y of %d. test case's %d. obstacle's start point must be greater than or equal to 0 and less than %d.\n", i, oi, t.size.height)
					}
				}

				err = true

				// Ask and validate the X of test case's obstacle's end point.
				for err {
					o.point.end.x = scan(fmt.Sprintf("X of %d. test case's %d. obstacle's end point: ", i, oi))

					err = o.point.end.x >= t.size.width

					if err {
						fmt.Printf("X of %d. test case's %d. obstacle's end point must be greater than or equal to 0 and less than %d.\n", i, oi, t.size.width)
					}
				}

				err = true

				// Ask and validate the Y of test case's obstacle's end point.
				for err {
					o.point.end.y = scan(fmt.Sprintf("Y of %d. test case's %d. obstacle's end point: ", i, oi))

					err = o.point.end.y >= t.size.height

					if err {
						fmt.Printf("Y of %d. test case's %d. obstacle's end point must be greater than or equal to 0 and less than %d.\n", i, oi, t.size.height)
					}
				}

				// Check if the obstacle contains the starting point.
				err = o.Contains(t.point.start)

				if err {
					fmt.Println("Error")
				} else {
					t.obstacles = append(t.obstacles, o)
				}
			}
		}

		ts = append(ts, t)
	}

	// Loop the test cases and calculate optimal solutions of them.
	for i, t := range ts {
		r, s := t.Calc()

		if !r {
			fmt.Println("No solution.")
		} else {
			fmt.Printf("%d. test case's optimal solution takes %d hops.\n", i+1, s)
		}
	}
}

func scan(format string) uint {
	var a uint

	fmt.Printf(format)

	_, err := fmt.Scan(&a)

	if err != nil {
		panic(err)
	}

	return a
}
