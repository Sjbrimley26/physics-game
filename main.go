package main

import (
	"fmt"
)

func update(objs []Movable) []Movable {
	var updated []Movable
	for _, obj := range objs {
		obj.Fall()
		obj.Move()
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

	hexagon := Movable{
		Shape: hex,
	}

	shapes := []Movable{hexagon}

	shapes = update(shapes)

	var data [][]Point

	for _, shape := range shapes {
		data = append(data, shape.Shape.Vertices())
	}

	fmt.Println(data)

}
