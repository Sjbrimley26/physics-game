package main

// import "github.com/nlepage/golang-wasm/js/bind"

// TDContext returns the context js object responsible for manipulating the canvas.
type TDContext struct {
	SetFillStyle   func(string)                                      `js:"context.fillStyle"`
	SetStrokeStyle func(string)                                      `js:"context.strokeStyle"`
	MoveTo         func(float64, float64)                            `js:"context.moveTo()"`
	LineTo         func(float64, float64)                            `js:"context.lineTo()"`
	Fill           func()                                            `js:"context.fill()"`
	Stroke         func()                                            `js:"context.stroke()"`
	BeginPath      func()                                            `js:"context.beginPath()"`
	Arc            func(float64, float64, float64, float64, float64) `js:"context.arc()"`
	ClearRect      func(float64, float64, float64, float64)          `js:"context.clearRect()"`
}
