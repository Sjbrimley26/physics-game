package main

func update(objs []Movable) []Movable {
	var updated []Movable
	for _, obj := range objs {
		obj.Fall()
		obj.Move()
		obj.Shape.Render("#000000", "rgb(0,120,200)")
		updated = append(updated, obj)
	}

	return updated
}

func main() {

	rpFactory := RegularPolygonFactory()

	hex := rpFactory.CreatePolygon(20, 6, Point{
		X: 100,
		Y: 100,
	})

	hexagon := Movable{Shape: hex}

	shapes := []Movable{hexagon}

	shapes = update(shapes)
}
