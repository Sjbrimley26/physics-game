package main

import (
	"math"
	"syscall/js"
)

// Renderable represents an object that can be rendered on the canvas.
type Renderable interface {
	Render()
}

var ctx Context
var canvas js.Value

func init() {
	canvas = js.
		Global().
		Get("document").
		Call("getElementById", "canvas")

	ctx = getContext(canvas)
}

func renderShape(vertices []Point, fill string, stroke string) {
	ctx.setFillStyle(fill)
	ctx.setStrokeStyle(stroke)
	ctx.beginPath()

	for i, p := range vertices {
		if i == 0 {
			ctx.moveTo(p.X, p.Y)
			continue
		}
		ctx.lineTo(p.X, p.Y)
	}

	ctx.fill()
	ctx.stroke()
}

func renderCircle(c *Circle, fill string, stroke string) {
	x, y := c.Center.X, c.Center.Y
	ctx.beginPath()
	ctx.setFillStyle(fill)
	ctx.setStrokeStyle(stroke)
	ctx.arc(x, y, c.Radius, 0, 2*math.Pi)
	ctx.fill()
	ctx.stroke()
}
