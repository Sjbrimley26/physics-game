package main

import (
	"math"
)

// Circle is a circle.
type Circle struct {
	Center Point
	Radius float64
}

// GetPointOnCircle returns a point on the circumference of a circle at a given angle.
func (c *Circle) GetPointOnCircle(angle float64) Point {
	x, y := c.Center.X, c.Center.Y
	radius := c.Radius
	return Point{
		X: radius*math.Sin(angle) + x,
		Y: radius*math.Cos(angle) + y,
	}
}

// Area returns the area of the circle.
func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Perimeter returns the circumference of the circle.
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
