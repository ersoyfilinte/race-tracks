package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"strconv"
)

type Case struct {
	grid struct {
		width, height int
	}

	points    Path
	obstacles []Path
}

type Path struct {
	start, end Position
}

type Position struct {
	x, y int
}

type Hop struct {
	point, speed Position
}

func (c *Case) Calc() (s uint, err bool) {
	var hs []Hop

	// Append test case's start point with initial speed to hops if start point is inside grid.
	if c.grid.width > c.points.start.x && c.grid.height > c.points.start.y {
		hs = append(hs, Hop{c.points.start, Position{}})
	}

	// Loop steps if not number of current hops equals number of previous hops.
	for n := 0; n != len(hs); s++ {
		n = len(hs)

		// Calculate new hops.
		for _, h := range hs {
			// Check if current hop's point equals test case's end point.
			if h.point == c.points.end {
				return s, false
			}

			// Calculate possible X positions of new hop's speed and point.
			for x := h.speed.x - 1; x <= h.speed.x+1; x++ {
				px := h.point.x + x

				// Validate X position of current hop's speed and point.
				if -3 > x || x > 3 || 0 > px || px >= c.grid.width {
					continue
				}

				// Calculate possible Y positions of new hop's speed and point.
				for y := h.speed.y - 1; y <= h.speed.y+1; y++ {
					py := h.point.y + y

					// Validate Y position of current hop's speed and point.
					if -3 > y || y > 3 || 0 > py || py >= c.grid.height {
						continue
					}
					
					err = false
					
					// Check if test case's obstacles overlaps current hop's point.
					for _, o := range c.obstacles {
						if err = o.start.x <= px && px <= o.end.x && o.start.y <= py && py <= o.end.y; err {
							break
						}
					}
					
					// Append current hop to hops if test case's obstacles does not overlap current hop's point.
					if !err {
						ho := Hop{Position{px, py}, Position{x, y}}
						
						// Check if not hops contain current hop.
						if !slices.Contains(hs, ho) {
							hs = append(hs, ho)
						}
					}
				}
			}
		}
	}

	return 0, true
}

func main() {
	var n, on int
	var cs []Case

	// Ask for number of test cases and validate.
	for {
		if n = Scan("Number of test cases: "); n >= 1 {
			break
		}

		Print(0, "Number of test cases must be greater than or equal to 1.")
	}

	// Ask for properties of each test cases and validate.
	for i := 1; i <= n; i++ {
		var c Case

		// Ask for width of current test case's grid and validate.
		for {
			if c.grid.width = Scan("Width of %d. test case's grid: ", i); 1 <= c.grid.width && c.grid.width <= 30 {
				break
			}

			Print(0, "Width of test case's grid must be greater than or equal to 1 and less than or equal to 30.")
		}

		// Ask for height of current test case's grid and validate.
		for {
			if c.grid.height = Scan("Height of %d. test case's grid: ", i); 1 <= c.grid.height && c.grid.height <= 30 {
				break
			}

			Print(0, "Height of test case's grid must be greater than or equal to 1 and less than or equal to 30.")
		}

		// Ask for X position of current test case's start point and validate.
		for {
			if c.points.start.x = Scan("X position of %d. test case's start point: ", i); c.points.start.x >= 0 {
				break
			}

			Print(0, "X position of %d. test case's start point must be greater than or equal to 0.", i)
		}

		// Ask for Y position of current test case's start point and validate.
		for {
			if c.points.start.y = Scan("Y position of %d. test case's start point: ", i); c.points.start.y >= 0 {
				break
			}

			Print(0, "Y position of %d. test case's start point must be greater than or equal to 0.", i)
		}

		// Ask for X position of current test case's end point and validate.
		for {
			if c.points.end.x = Scan("X position of %d. test case's end point: ", i); c.points.end.x < c.grid.width {
				break
			}

			Print(0, "X position of %d. test case's end point must be less than %d.", i, c.grid.width)
		}

		// Ask for Y position of current test case's end point and validate.
		for {
			if c.points.end.y = Scan("Y position of %d. test case's end point: ", i); c.points.end.y < c.grid.height {
				break
			}

			Print(0, "Y position of %d. test case's end point must be less than %d.", i, c.grid.height)
		}

		// Ask for number of current test case's obstacles and validate.
		for {
			if on = Scan("Number of %d. test case's obstacles: ", i); 0 <= on && on < c.grid.width*c.grid.height {
				break
			}

			Print(0, "Number of %d. test case's obstacles must be greater than or equal to 0 and less than %d.", i, c.grid.width*c.grid.height)
		}

		// Ask for properties of current test case's each obstacle and validate.
		for oi := 1; oi <= on; oi++ {
			for {
				var o Path

				// Ask for X position of current test case's current obstacle's start point and validate.
				for {
					if o.start.x = Scan("X position of %d. test case's %d. obstacle's start point: ", i, oi); 0 <= o.start.x && o.start.x < c.grid.width {
						break
					}

					Print(0, "X position of %d. test case's obstacle's start point must be greater than or equal to 0 and less than %d.", i, c.grid.width)
				}

				// Ask for Y position of current test case's current obstacle's start point and validate.
				for {
					if o.start.y = Scan("Y position of %d. test case's %d. obstacle's start point: ", i, oi); 0 <= o.start.y && o.start.y < c.grid.height {
						break
					}

					Print(0, "Y position of %d. test case's obstacle's start point must be greater than or equal to 0 and less than %d.", i, c.grid.height)
				}

				// Ask for X position of current test case's current obstacle's end point and validate.
				for {
					if o.end.x = Scan("X position of %d. test case's %d. obstacle's end point: ", i, oi); o.start.x <= o.end.x && o.end.x < c.grid.width {
						break
					}

					Print(0, "X position of %d. test case's obstacle's end point must be greater than or equal to %d and less than %d.", i, o.start.x, c.grid.width)
				}

				// Ask for Y position of current test case's current obstacle's end point and validate.
				for {
					if o.end.y = Scan("Y position of %d. test case's %d. obstacle's end point: ", i, oi); o.start.y <= o.end.y && o.end.y < c.grid.height {
						break
					}

					Print(0, "Y position of %d. test case's obstacle's end point must be greater than or equal to %d and less than %d.", i, o.start.y, c.grid.height)
				}

				// Check if not current test case's current obstacle overlaps own start point.
				if !(o.start.x <= c.points.start.x && c.points.start.x <= o.end.x && o.start.y <= c.points.start.y && c.points.start.y <= o.end.y) {
					c.obstacles = append(c.obstacles, o)

					break
				}

				Print(0, "Test case's obstacle should not overlap own start point.")
			}
		}

		cs = append(cs, c)
	}

	// Calculate each test case's optimal solutions.
	for i, c := range cs {
		Print(2, "Calculating %d. test case.", i+1)

		if s, err := c.Calc(); err {
			Print(1, "%d. test case has no solution.", i+1)
		} else {
			if s > 1 {
				Print(3, "%d. test case's optimal solution takes %d hops.", i+1, s)
			} else {
				Print(3, "%d. test case's optimal solution takes %d hop.", i+1, s)
			}
		}
	}
}

func Scan(format string, a ...any) (n int) {
	var s string

	for {
		fmt.Printf(format, a...)

		if _, err := fmt.Scanln(&s); err == nil {
			if n, err = strconv.Atoi(s); err == nil {
				return n
			}

			Print(0, "Enter a valid number.")
		}
	}
}

func Print(color int, format string, a ...any) {
	cs := [4]string{"\033[1;31m%s\033[0m", "\033[1;33m%s\033[0m", "\033[1;34m%s\033[0m", "\033[1;36m%s\033[0m"}

	fmt.Printf(cs[color], fmt.Sprintf(format, a...)+"\n")
}
