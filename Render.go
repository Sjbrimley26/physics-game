package main

import (
	"github.com/nlepage/golang-wasm/js/bind"
	"math"
)

// Renderable represents an object that can be rendered on the canvas.
type Renderable interface {
	Render()
}

func renderShape(vertices []Point, fill string, stroke string) {
	ctx := TDContext{}

	if err := bind.BindGlobals(ctx); err != nil {
		panic(err)
	}

	ctx.SetFillStyle(fill)
	ctx.SetStrokeStyle(stroke)
	ctx.BeginPath()

	for i, p := range vertices {
		if i == 0 {
			ctx.MoveTo(p.X, p.Y)
			continue
		}
		ctx.LineTo(p.X, p.Y)
	}

	ctx.Fill()
	ctx.Stroke()
}

func renderCircle(c *Circle, fill string, stroke string) {
	ctx := TDContext{}

	if err := bind.BindGlobals(ctx); err != nil {
		panic(err)
	}

	x, y := c.Center.X, c.Center.Y
	ctx.BeginPath()
	ctx.SetFillStyle(fill)
	ctx.SetStrokeStyle(stroke)
	ctx.Arc(x, y, c.Radius, 0, 2*math.Pi)
	ctx.Fill()
	ctx.Stroke()
}
