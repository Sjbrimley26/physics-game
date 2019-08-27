package main

import (
	"syscall/js"
)

// Context allows easy access to canvas methods
type Context struct {
	ctx js.Value
}

func (c *Context) setFillStyle(fill string) {
	c.ctx.Set("fillStyle", fill)
}

func (c *Context) setStrokeStyle(stroke string) {
	c.ctx.Set("strokeStyle", stroke)
}

func (c *Context) moveTo(x float64, y float64) {
	c.ctx.Call("moveTo", x, y)
}

func (c *Context) lineTo(x float64, y float64) {
	c.ctx.Call("lineTo", x, y)
}

func (c *Context) fill() {
	c.ctx.Call("fill")
}

func (c *Context) stroke() {
	c.ctx.Call("stroke")
}

func (c *Context) beginPath() {
	c.ctx.Call("beginPath")
}

func (c *Context) arc(x float64, y float64, radius float64, startAngle float64, endAngle float64) {
	c.ctx.Call("arc", x, y, radius, startAngle, endAngle)
}

func (c *Context) clearRect(x float64, y float64, width float64, height float64) {
	c.ctx.Call("clearRect", x, y, width, height)
}

func getContext(canvas js.Value) Context {
	ctx := canvas.Call("getContext", "2d")
	return Context{
		ctx,
	}
}
