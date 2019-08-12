package point

import (
	"math"
)

// Distance finds the distance between 2 points.
func Distance(a, b Point) float64 {
	x1, y1 := a.X, a.Y
	x2, y2 := b.X, b.Y
	return math.Sqrt(
		math.Pow(float64(x2-x1), 2) +
			math.Pow(float64(y2-y1), 2))
}

// Orientation finds the orientation of 3 points.
func Orientation(p1, p2, p3 Point) int {
	// https://www.geeksforgeeks.org/orientation-3-ordered-points/
	// 0 == colinear
	// 1 == clockwise
	// 2 == counter-clockwise
	val := (p2.Y-p1.Y)*(p3.X-p2.X) -
		(p2.X-p1.X)*(p3.Y-p2.Y)

	if val == 0 {
		return 0 // colinear
	} else if val > 0 {
		return 1 // clockwise
	} else {
		return 2 // counter-clockwise
	}
}
