package main

import (
	"encoding/json"
	"fmt"
	"physics-game/movable"
	"physics-game/point"
	"physics-game/polygons"
)

func main() {
	rpFactory := polygons.RegularPolygonFactory()

	hex := rpFactory.CreatePolygon(20, 6, point.Point{
		X: 100,
		Y: 100,
	})

	hexagon := movable.Movable{
		Shape: hex,
	}

	shapes := []movable.Movable{hexagon}

	var data []string

	for _, shape := range shapes {
		j, err := json.Marshal(shape)
		if err != nil {
			fmt.Println(err)
			return
		}
		s := string(j)
		data = append(data, s)
		fmt.Println(s)
	}

}
