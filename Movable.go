package main

import (
	"math"
)

// Movable represents a moving object, affected by gravity.
type Movable struct {
	Mass, Elasticity, Acceleration float64
	Velocity                       Vector
	Shape                          RegularPolygon
}

// Momentum returns the object's momentum
func (m *Movable) Momentum() float64 {
	return m.Mass * m.Velocity.Magnitude
}

// Move translates the current position and decreases the velocity a bit.
func (m *Movable) Move() {
	ap := m.Shape.Apothem()
	m.Shape.Center = m.Shape.Center.AddVector(m.Velocity)
	m.Velocity = Vector{
		Direction: m.Velocity.Direction,
		Magnitude: m.Velocity.Magnitude * 0.95,
	}
	if m.Shape.Center.Y+ap > 500 {
		m.Shape.Center = Point{
			X: m.Shape.Center.X,
			Y: 500 - ap,
		}
	}
	if m.Shape.Center.X < 0 {
		m.Velocity = m.Velocity.Add(Vector{0, 1})
	}
	if m.Shape.Center.X > 800 {
		m.Velocity = m.Velocity.Add(Vector{math.Pi, 1})
	}
}

// Fall applies gravity to the object.
func (m *Movable) Fall() {
	if m.Shape.Center.Y+m.Shape.Apothem() < 500 {
		m.Velocity = m.Velocity.Add(Vector{math.Pi / 2, 0.5})
	}
}
