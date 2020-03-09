//An array of n pairs of reals (ai,bi) defined the n lines yi=aix+b. The lines were ordered in the x-interval [0,1]
//in the sense that yi<yi+1 for all values  i between 0 and n-2 and all values of x in [0,1].
//Given a point (x,y) where 0=<x<=1 determine the two lines that bracket the point

package linesbracketpoint

import (
	"fmt"
)

type point struct {
	x float64
	y float64
}

func (p *point) isValid() bool {
	return p.x <= 1 && p.x >= 0
}

func (p *point) below(line line) bool {
	xBreakingPoint := (p.y - line.y) / line.x //Determine the x coordinate where the point is below or above the line
	if line.x > 0 {                           //the func is growing
		return p.x >= xBreakingPoint
	}

	return p.x <= xBreakingPoint
}

func (p *point) above(line line) bool {
	return !p.below(line)
}

type line point

func (l line) String() string {
	return fmt.Sprintf("%.3fx + %.3f", l.x, l.y)
}

//Assume the lines are ordered in the x-interval [0,1]
func getLinesBracketPoint(points []point, p point) (found bool, bracket [2]line) {
	numPoints := len(points)
	if numPoints < 2 || !p.isValid() {
		return false, bracket
	}
	interval := getInterval(points, p)
	if interval != ([2]point{}) {
		return true, [2]line{line(interval[0]), line(interval[1])}
	}
	return false, bracket
}

func getInterval(points []point, p point) [2]point {
	size := len(points)
	if size == 2 {
		if p.below(line(points[1])) && p.above(line(points[0])) {
			return [2]point{points[0], points[1]}
		}
		return [2]point{}
	}

	midpoint := (size - 1) / 2

	if p.below(line(points[midpoint])) {
		return getInterval(points[:midpoint+1], p)
	}

	return getInterval(points[midpoint:], p)

}
