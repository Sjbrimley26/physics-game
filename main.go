package main

import (
	"syscall/js"
)

var loop js.Func

func main() {

	done := make(chan struct{}, 0)

	rpFactory := RegularPolygonFactory()

	hex := rpFactory.CreatePolygon(100, 6, Point{
		X: 100,
		Y: 100,
	})

	hexagon := Movable{Shape: hex}

	shapes := []Movable{hexagon}

	update := func(objs []Movable) []Movable {
		ctx.clearRect(0, 0, 800, 600)
		var updated []Movable
		for _, obj := range objs {
			obj.Fall()
			obj.Move()
			obj.Shape.Render("#000000", "rgb(0,120,200)")
			updated = append(updated, obj)
		}

		return updated
	}

	loop = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		shapes = update(shapes)
		js.Global().Call("requestAnimationFrame", loop)
		return nil
	})

	defer loop.Release()

	js.Global().Call("requestAnimationFrame", loop)

	<-done

}
