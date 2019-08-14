package main

import "math"

// Vector is a measure of force with direction and magnitude!
type Vector struct {
	Direction, Magnitude float64
}

// X returns the change in X from the start of the vector until the end;
func (v *Vector) X() float64 {
	return v.Magnitude * math.Cos(v.Direction)
}

// Y returns the change in Y from the start of the vector until the end;
func (v *Vector) Y() float64 {
	return v.Magnitude * math.Sin(v.Direction)
}

// Add returns a new vector created from adding 2 others.
func (v *Vector) Add(v2 Vector) Vector {
	x0, y0 := v.X(), v.Y()
	x1, y1 := v2.X(), v2.Y()
	x := x0 + x1
	y := y0 + y1
	dir := math.Atan(y / x)
	if math.IsNaN(dir) {
		if y1 > y0 {
			dir = 3 * math.Pi / 2
		} else {
			dir = math.Pi / 2
		}
	}
	mag := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	return Vector{
		Direction: dir,
		Magnitude: mag,
	}
}

// Inverse returns a new vector of the same magnitude, going the opposite direction.
func (v *Vector) Inverse() Vector {
	dir := v.Direction + math.Pi
	mag := v.Magnitude
	return Vector{
		Direction: dir,
		Magnitude: mag,
	}
}

// Subtract returns a new vector, the first subtracted by the second.
func (v *Vector) Subtract(v2 *Vector) Vector {
	return v.Add(v2.Inverse())
}

// Scale returns a new vector, with the magnitude multiped by the given scalar.
func (v *Vector) Scale(s float64) Vector {
	return Vector{
		Direction: v.Direction,
		Magnitude: v.Magnitude * s,
	}
}

// DotProduct returns a dot-product... idk
func (v *Vector) DotProduct(v2 Vector) float64 {
	return v.X()*v2.X() + v.Y()*v2.Y()
}

// CrossProduct returns an float64 which signifies the angle between the 2 vectors.
func (v *Vector) CrossProduct(v2 Vector) float64 {
	// sign represents whether 2nd vector is on the left or right
	// sin-1(abs(crossProduct)) === the angle between the two vectors
	return v.X()*v2.Y() - v.Y()*v2.X()
}

// GetPerpendicular returns a new Vector, moving perpendicular to the original.
func (v *Vector) GetPerpendicular() Vector {
	x, y := float64(v.Y()), float64(-v.X())
	mag := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	dir := math.Atan(y / x)
	return Vector{
		Direction: dir,
		Magnitude: mag,
	}
}
