package main

// Rectangle is a quadrilateral with 4 right angles.
type Rectangle struct {
	Center                  Point
	Length, Width, Rotation float64
}

// Area returns the rectangle's area.
func (r *Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Vertices returns an array of the shape's vertices.
func (r *Rectangle) Vertices() []Point {
	x, y := r.Center.X, r.Center.Y
	l, w := r.Length, r.Width
	return []Point{
		Point{x - w/2, y - l/2},
		Point{x - w/2, y + l/2},
		Point{x + w/2, y - l/2},
		Point{x + w/2, y + l/2},
	}
}

// Edges returns the lines that comprise the perimeter of the shape.
func (r *Rectangle) Edges() []Line {
	edges := make([]Line, 4)
	vertices := r.Vertices()
	for i := 0; i < len(vertices); i++ {
		if i == len(vertices)-1 {
			edges[i] = Line{
				Start: vertices[i],
				End:   vertices[0],
			}
		} else {
			edges[i] = Line{
				Start: vertices[i],
				End:   vertices[i+1],
			}
		}
	}
	return edges
}

// Perimeter returns the sum of the length of the edges.
func (r *Rectangle) Perimeter() float64 {
	edges := r.Edges()
	total := 0.0
	for _, edge := range edges {
		total += edge.Length()
	}
	return total
}

// IrregularApothem returns the radius of the largest circle that will fit in the shape originating from it's center.
func (r *Rectangle) IrregularApothem() float64 {
	edges, center := r.Edges(), r.Center
	possApothems := make([]Line, len(edges))
	for i := 0; i < len(edges); i++ {
		possApothems[i] = Line{
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
func (r *Rectangle) Midpoints() []Point {
	edges := r.Edges()
	midpoints := make([]Point, len(edges))
	for i := 0; i < len(edges); i++ {
		midpoints[i] = edges[i].Center()
	}
	return midpoints
}

// Normals returns an array of the lines perpendicular to the shape's edges
func (r *Rectangle) Normals() []Line {
	edges := r.Edges()
	normals := make([]Line, len(edges))
	for i := 0; i < len(edges); i++ {
		normals[i] = edges[i].GetPerpendicular()
	}
	return normals
}

// Bottom returns y-value of the lowest Point of the shape.
func (r *Rectangle) Bottom() float64 {
	vertices := r.Vertices()
	bottom := vertices[0].Y
	for _, p := range vertices {
		if p.Y > bottom {
			bottom = p.Y
		}
	}
	return bottom
}

// Left returns the lowest x-value of the shapes vertices'
func (r *Rectangle) Left() float64 {
	vertices := r.Vertices()
	left := vertices[0].X
	for _, p := range vertices {
		if p.X < left {
			left = p.X
		}
	}
	return left
}

// Right returns the greatest x-value of the shapes vertices'
func (r *Rectangle) Right() float64 {
	vertices := r.Vertices()
	right := vertices[0].X
	for _, p := range vertices {
		if p.X > right {
			right = p.X
		}
	}
	return right
}

// Top returns the greatest x-value of the shapes vertices'
func (r *Rectangle) Top() float64 {
	vertices := r.Vertices()
	top := vertices[0].Y
	for _, p := range vertices {
		if p.Y < top {
			top = p.Y
		}
	}
	return top
}

// Render renders the shape on the canvas.
func (r *Rectangle) Render(stroke string, fill string) {
	renderShape(r.Vertices(), fill, stroke)
}

// Covers returns whether this rectangle overlaps the given rectangle
func (r *Rectangle) Covers(r1 *Rectangle) bool {
	l1, r0, b1, t1 := r.Left(), r.Right(), r.Bottom(), r.Top()
	l2, r2, b2, t2 := r1.Left(), r1.Right(), r1.Bottom(), r1.Top()
	if l2 >= l1 &&
		r2 <= r0 &&
		t2 >= t1 &&
		b2 <= b1 {
		return true
	}
	return false
}

// MBB returns the smallest rectangle which contains all the given points.
func MBB(points []Point) (Rectangle, bool) {
	if len(points) < 3 {
		return Rectangle{}, true
	}

	var l, r, t, b float64

	for i, p := range points {
		if i == 0 {
			l = p.X
			r = p.X
			t = p.Y
			b = p.Y
			continue
		}
		if l > p.X {
			l = p.X
		}
		if r < p.X {
			r = p.X
		}
		if t > p.Y {
			t = p.Y
		}
		if b < p.Y {
			b = p.Y
		}
	}

	center := Point{(l + r) / 2, (t + b) / 2}
	length := b - t
	width := r - l

	return Rectangle{center, length, width, 0}, false
}
