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
	if p.below(line(points[numPoints-1])) && p.above(line(points[0])) {
		interval := getInterval(points, p)
		return true, [2]line{line(interval[0]), line(interval[1])}
	}
	return false, bracket
}

func getInterval(points []point, p point) [2]point {
	size := len(points)
	if size == 2 {
		return [2]point{points[0], points[1]}
	}
	midpoint := (size - 1) / 2

	if p.below(line(points[midpoint])) {
		return getInterval(points[:midpoint+1], p)
	}

	return getInterval(points[midpoint:], p)

}
