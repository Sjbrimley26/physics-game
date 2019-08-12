package polygons

import (
	"math"
	"physics-game/circle"
	"physics-game/line"
	"physics-game/point"
)

func toRadians(d float64) float64 {
	return d * (math.Pi / 180)
}

// RegularPolygon is a polygon where all sides are the same length
type RegularPolygon struct {
	Center               point.Point
	Sides                int
	SideLength, Rotation float64
}

// Apothem returns the radius of the largest circle that will fit in the shape.
func (p *RegularPolygon) Apothem() float64 {
	sL, sides := p.SideLength, float64(p.Sides)
	return math.Abs(sL / math.Tan(toRadians(180/sides)))
}

// Area returns the area of the shape.
func (p *RegularPolygon) Area() float64 {
	perim, apoth := p.Perimeter(), p.Apothem()
	return perim * apoth / 2
}

// Circumcircle returns the smallest circle which contains all of the shape's vertices.
func (p *RegularPolygon) Circumcircle() circle.Circle {
	center, sL, sides := p.Center, p.SideLength, p.Sides
	return circle.Circle{
		Center: center,
		Radius: sL / 2 * math.Sin(toRadians(180/float64(sides))),
	}
}

// Vertices returns an array of the shape's vertices.
func (p *RegularPolygon) Vertices() []point.Point {
	sides, rotation, circle := p.Sides, p.Rotation, p.Circumcircle()
	a := float64(360 / sides)
	start := float64(a/2) - rotation

	var vertices []point.Point

	for angle := start; angle < 360+start; angle += a {
		r := toRadians(angle)
		vertices = append(vertices, circle.GetPointOnCircle(r))
	}

	return vertices
}

// Edges returns the lines that comprise the perimeter of the shape.
func (p *RegularPolygon) Edges() []line.Line {
	edges := make([]line.Line, p.Sides)
	vertices := p.Vertices()
	for i := 0; i < len(vertices); i++ {
		if i == len(vertices)-1 {
			edges[i] = line.Line{
				Start: vertices[i],
				End:   vertices[0],
			}
		} else {
			edges[i] = line.Line{
				Start: vertices[i],
				End:   vertices[i+1],
			}
		}
	}
	return edges
}

// Perimeter returns the sum of the length of the edges.
func (p *RegularPolygon) Perimeter() float64 {
	edges := p.Edges()
	total := 0.0
	for _, edge := range edges {
		total += edge.Length()
	}
	return total
}

// IrregularApothem returns the radius of the largest circle that will fit in the shape originating from it's center.
func (p *RegularPolygon) IrregularApothem() float64 {
	edges, center := p.Edges(), p.Center
	possApothems := make([]line.Line, len(edges))
	for i := 0; i < len(edges); i++ {
		possApothems[i] = line.Line{
			Start: center,
			End:   edges[i].Center(),
		}
	}

	shortest := possApothems[0].Length()
	for i := 1; i < len(possApothems); i++ {
		a, b := shortest, possApothems[i].Length()
		if b < a {
			shortest = b
		}
	}

	return shortest
}

// Midpoints returns an array of the center points of each edge
func (p *RegularPolygon) Midpoints() []point.Point {
	edges := p.Edges()
	midpoints := make([]point.Point, len(edges))
	for i := 0; i < len(edges); i++ {
		midpoints[i] = edges[i].Center()
	}
	return midpoints
}

// Normals returns an array of the lines perpendicular to the shape's edges
func (p *RegularPolygon) Normals() []line.Line {
	edges := p.Edges()
	normals := make([]line.Line, len(edges))
	for i := 0; i < len(edges); i++ {
		normals[i] = edges[i].GetPerpendicular()
	}
	return normals
}

// Bottom returns y-value of the lowest Point of the shape.
func (p *RegularPolygon) Bottom() float64 {
	vertices := p.Vertices()
	bottom := vertices[0].Y
	for _, p := range vertices {
		if p.Y > bottom {
			bottom = p.Y
		}
	}
	return bottom
}
