// +build js,wasm
package main

import (
	"math"
	"syscall/js"

	"github.com/mvrilo/go-particles/particles"
)

// Canvas struct holds the Javascript objects needed for the Canvas creation
type Canvas struct {
	done chan struct{}

	width  int
	height int

	window js.Value
	doc    js.Value
	body   js.Value

	// Canvas properties
	canvas js.Value
	ctx    js.Value
	reqID  js.Value

	group *particles.Group
}

// NewCanvas initializes a Canvas element
func NewCanvas(width, height, maxparticles int) *Canvas {
	var c Canvas
	c.window = js.Global()
	c.doc = c.window.Get("document")
	c.body = c.doc.Get("body")
	c.done = make(chan struct{})

	c.width = width
	c.height = height

	c.group = particles.NewGroup(width, height, maxparticles, particles.DefaultConfig)
	c.canvas = c.doc.Call("createElement", "canvas")
	c.ctx = c.canvas.Call("getContext", "2d")

	c.canvas.Set("id", "particles")
	c.canvas.Set("width", width)
	c.canvas.Set("height", height)
	c.body.Call("appendChild", c.canvas)

	return &c
}

// Render calls the `requestAnimationFrame` Javascript function in asynchronous mode
func (c *Canvas) Render() {
	var render js.Func
	render = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			c.reqID = c.window.Call("requestAnimationFrame", render)
			c.Clear()
			c.group.Move()
			c.Draw()
		}()
		return nil
	})
	c.window.Call("requestAnimationFrame", render)
	<-c.done
}

// Draw draws elements in the canvas
func (c *Canvas) Draw() {
	for _, particle := range c.group.Particles {
		// println(fmt.Sprintf("%+v\n", particle))
		c.DrawParticle(particle)
	}
}

// DrawParticle draws elements in the canvas
func (c *Canvas) DrawParticle(particle *particles.Particle) {
	c.ctx.Call("beginPath")
	c.ctx.Call("arc", particle.Position[0], particle.Position[1], particle.Size, 0, 2*math.Pi, true)
	c.ctx.Call("closePath")
	c.ctx.Set("fillStyle", particle.Color)
	c.ctx.Call("fill")
}

// Clear clears the canvas
func (c *Canvas) Clear() {
	c.ctx.Call("clearRect", 0, 0, c.width, c.height)
}

// Stop stops the rendering
func (c *Canvas) Stop() {
	c.window.Call("cancelAnimationFrame", c.reqID)
	c.done <- struct{}{}
	close(c.done)
}

func main() {
	canvas := NewCanvas(800, 300, 50)
	canvas.Render()
}
