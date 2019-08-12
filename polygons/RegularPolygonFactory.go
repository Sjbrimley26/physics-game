package polygons

import (
	"physics-game/point"
)

// RPFactory is a struct which creates new polygons.
type RPFactory struct{}

// CreatePolygon returns a RegularPolygon with the given size and number of sides.
func (f RPFactory) CreatePolygon(sideLength float64, sides int, center point.Point) RegularPolygon {
	return RegularPolygon{
		SideLength: sideLength,
		Sides:      sides,
		Center:     center,
	}
}

// RegularPolygonFactory returns an RPFactory.
func RegularPolygonFactory() RPFactory {
	return RPFactory{}
}
