package line

import (
	"math"
	"physics-game/physics"
	"physics-game/point"
)

// Line is a line between 2 points.
type Line struct {
	Start, End point.Point
}

// Slope returns the rise / run of the line.
func (l *Line) Slope() float64 {
	x1 := l.Start.X
	y1 := l.Start.Y
	x2 := l.End.X
	y2 := l.End.Y

	if x1 == x2 {
		// Vertical line
		if y1 < y2 {
			return math.Inf(1)
		}
		return math.Inf(-1)
	}

	if y1 == y2 {
		// Horizontal line
		return 0
	}

	return float64((y2 - y1) / (x2 - x1))
}

// Length returns the length of the line.
func (l *Line) Length() float64 {
	return point.Distance(l.Start, l.End)
}

// Center returns the center point of the line.
func (l *Line) Center() point.Point {
	return point.Point{
		X: (l.End.X + l.Start.X) / 2,
		Y: (l.End.Y + l.Start.Y) / 2,
	}
}

// YInt returns the y-intercept.
func (l *Line) YInt() float64 {
	if l.Start.X == l.End.X ||
		l.Start.Y == l.End.Y {
		return l.Start.Y
	}
	return l.Start.Y - l.Slope()*l.Start.X
}

// XInt returns the x-intercept.
func (l *Line) XInt() float64 {
	if l.Start.X == l.End.X ||
		l.Start.Y == l.End.Y {
		return l.Start.X
	}
	return (-l.Slope()*l.Start.X - l.Start.Y) / l.Slope()
}

// IsPointOnLine returns a bool of whether the point rests on the line.
func (l *Line) IsPointOnLine(p point.Point) bool {
	x, y := float64(l.Start.X), float64(l.Start.Y)
	x2, y2 := float64(p.X), float64(p.Y)
	return (x2-x)*l.Slope() == y2-y
}

// IntersectsWith returns a bool of whether the 2 lines intersect
func (l *Line) IntersectsWith(l2 *Line) bool {
	// https://www.geeksforgeeks.org/check-if-two-given-line-segments-intersect/
	p1, q1 := l.Start, l.End
	p2, q2 := l2.Start, l2.End
	o1 := point.Orientation(p1, q1, p2)
	o2 := point.Orientation(p1, q1, q2)
	o3 := point.Orientation(p2, q2, p1)
	o4 := point.Orientation(p2, q2, q1)

	// General case
	if o1 != o2 && o3 != o4 {
		return true
	}

	// Special Cases
	// p1, q1 and p2 are colinear and p2 lies on segment p1q1
	if o1 == 0 && l.IsPointOnLine(p2) {
		return true
	}
	// p1, q1 and q2 are colinear and q2 lies on segment p1q1
	if o2 == 0 && l.IsPointOnLine(q2) {
		return true
	}
	// p2, q2 and p1 are colinear and p1 lies on segment p2q2
	if o3 == 0 && l2.IsPointOnLine(p1) {
		return true
	}
	// p2, q2 and q1 are colinear and q1 lies on segment p2q2
	if o4 == 0 && l2.IsPointOnLine(q1) {
		return true
	}
	return false // Doesn't fall in any of the above cases
}

// GetPointOfIntersection returns the point where two lines cross
func (l *Line) GetPointOfIntersection(l2 *Line) (point.Point, bool) {
	if l.IntersectsWith(l2) != true || l.Slope() == l2.Slope() {
		return point.Point{}, false
	}

	x0, y0, m0 := float64(l.Start.X), float64(l.Start.Y), l.Slope()
	x1, y1, m1 := float64(l2.Start.X), float64(l2.Start.Y), l2.Slope()

	x := (m0*x0 - m1*x1 + y1 - y0) / (m0 - m1)
	y := (m0*m1*(x1-x0) + m1*y0 - m0*y1) / (m1 - m0)

	return point.Point{
		X: x,
		Y: y,
	}, true
}

/*
	Usage:
	p, intersecting := l.GetPointOfIntersection(l2)
	if intersecting != false
		continue...
*/

// GetPerpendicular returns a new line, perpendicular to the original that goes through the center of the first.
func (l *Line) GetPerpendicular() Line {
	slope, length, center := l.Slope(), l.Length(), l.Center()
	x, y := center.X, center.Y
	inv := -1 * (1 / slope)
	var x0, x1, y0, y1, b float64

	if inv == 0 {
		x0 = x
		x1 = x + length
		y0, y1 = y, y
	} else if inv == math.Inf(1) {
		y0 = y
		y1 = y + length
		x0, x1 = x, x
	} else if inv == math.Inf(-1) {
		y0 = y
		y1 = y - length
		x0, x1 = x, x
	} else {
		b = float64(y) - inv*float64(x)

		x0 = -l.YInt()
		y0 = x0*inv + b

		x1 = x0 + math.Cos(inv)*length
		y1 = x1*inv + b
	}

	return Line{
		Start: point.Point{X: x0, Y: y0},
		End:   point.Point{X: x1, Y: y1},
	}
	// The perpendicular is a tad shorter than the original but it should work
}

// ToVector creates a vector with the same length and direction as the given line.
func (l *Line) ToVector() physics.Vector {
	m, len := l.Slope(), l.Length()
	dir := math.Atan(m)
	if dir == math.Inf(1) {
		dir = 3 * math.Pi / 2
	}
	if dir == math.Inf(-1) {
		dir = math.Pi / 2
	}
	return physics.Vector{
		Direction: dir,
		Magnitude: len,
	}
}

// AddVector returns a new line, offset by a given vector.
func (l Line) AddVector(v physics.Vector) Line {
	start, end := l.Start, l.End
	start = start.AddVector(v)
	end = end.AddVector(v)
	return Line{
		Start: start,
		End:   end,
	}
}
