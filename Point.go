package main

import (
	"math"
)

// Point is a 2d point in space.
type Point struct {
	X, Y float64
}

// ToVector creates a vector from a point.
func (p *Point) ToVector() Vector {
	magnitude := math.Sqrt(
		math.Pow(p.X, 2) +
			math.Pow(p.Y, 2))

	direction := math.Atan(p.Y / p.X)

	return Vector{
		Direction: direction,
		Magnitude: magnitude,
	}
}

// AddVector returns a new point offset by a vector.
func (p *Point) AddVector(v Vector) Point {
	x := v.X()
	y := v.Y()
	return Point{
		X: x + p.X,
		Y: y + p.Y,
	}
}
