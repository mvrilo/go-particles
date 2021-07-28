package canvas

import (
	"math"
	"syscall/js"
	"time"

	"github.com/go-playground/colors"
	"github.com/mvrilo/go-particles/particles"
)

// Canvas element and data
type Canvas struct {
	done chan struct{}

	fps          time.Duration
	width        int
	height       int
	background   string
	maxparticles int
	config       particles.Config

	window js.Value
	doc    js.Value
	body   js.Value

	canvas js.Value
	ctx    js.Value
	reqID  js.Value

	group *particles.Group
}

// NewCanvas initializes a Canvas element
func NewCanvas(id string, fps time.Duration, background string, maxparticles int, particleConfig particles.Config) *Canvas {
	win := js.Global()
	doc := win.Get("document")
	body := doc.Get("body")
	canvas := doc.Call("createElement", "canvas")
	ctx := canvas.Call("getContext", "2d")
	canvas.Set("id", id)

	return &Canvas{
		fps:          fps,
		maxparticles: maxparticles,
		body:         body,
		canvas:       canvas,
		config:       particleConfig,
		ctx:          ctx,
		doc:          doc,
		window:       win,
		background:   background,
		done:         make(chan struct{}),
	}
}

func (c *Canvas) Start() {
	c.group = particles.NewGroup(c.width, c.height, c.maxparticles, c.config)
}

// AppendElement append the canvas element to the body
func (c *Canvas) AppendElement() {
	c.body.Call("appendChild", c.canvas)
}

// onMousemove sticks a particle to mouse movement
func (c *Canvas) onMousemove(args []js.Value) {
	event := args[0]
	event.Call("preventDefault")

	x := event.Get("offsetX").Float()
	y := event.Get("offsetY").Float()

	mouse := c.group.Particles[0]
	if mouse.Config.Move {
		mouse.Config.Move = false
	}

	mouse.Position = particles.Vector{x, y}
}

// onResize is a window resize event handler
func (c *Canvas) onResize() {
	c.Fullscreen()
	c.Start()
}

// ListenEvents add events listener, such as: resize
func (c *Canvas) ListenEvents() {
	onresize := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		c.onResize()
		return nil
	})
	onmousemove := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		c.onMousemove(args)
		return nil
	})
	c.window.Call("addEventListener", "resize", onresize)
	c.window.Call("addEventListener", "mousemove", onmousemove)
}

// Size sets a size for the canvas
func (c *Canvas) Size(width, height int) {
	c.width = width
	c.height = height
	c.canvas.Set("width", width)
	c.canvas.Set("height", height)
}

// Fullscreen set the size of the canvas as the size of the window
func (c *Canvas) Fullscreen() {
	width := c.window.Get("innerWidth").Int()
	height := c.window.Get("innerHeight").Int()
	c.Size(width, height)
}

// Background fills the background with color
func (c *Canvas) Background() {
	if c.background != "" {
		c.ctx.Set("fillStyle", c.background)
		c.ctx.Call("fillRect", 0, 0, c.width, c.height)
	}
}

// Render renders via requestAnimationFrame
func (c *Canvas) Render() {
	var render js.Func
	render = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			c.group.Move()

			c.Clear()
			c.Background()
			c.Draw()

			time.Sleep((1000 / c.fps) * time.Millisecond)
			c.reqID = c.window.Call("requestAnimationFrame", render)
		}()
		return nil
	})
	c.window.Call("requestAnimationFrame", render)
	<-c.done
}

// Draw draws elements in the canvas
func (c *Canvas) Draw() {
	for _, particle := range c.group.Particles {
		c.DrawParticle(particle)
		for _, p2 := range c.group.Particles {
			if particle.Distance(p2) >= particle.Area {
				continue
			}
			c.DrawConnection(particle, p2)
		}
	}
}

// DrawConnection draws a line between two vectors
func (c *Canvas) DrawConnection(p1, p2 *particles.Particle) {
	dist := p1.Distance(p2)
	alpha := 1.0 - (dist / p1.Area)
	color, _ := colors.Parse(p1.Color)
	rgba := color.ToRGBA()
	rgba.A = alpha

	c.ctx.Set("strokeStyle", rgba.String())
	c.ctx.Set("lineWidth", 1)
	c.ctx.Call("beginPath")
	c.ctx.Call("moveTo", p1.Position[0], p1.Position[1])
	c.ctx.Call("lineTo", p2.Position[0], p2.Position[1])
	c.ctx.Call("stroke")
	c.ctx.Call("closePath")
}

// DrawParticle draws elements in the canvas
func (c *Canvas) DrawParticle(particle *particles.Particle) {
	c.ctx.Call("beginPath")
	c.ctx.Call("arc", particle.Position[0], particle.Position[1], particle.Size, 0, 2*math.Pi, true)
	c.ctx.Set("fillStyle", particle.Color)
	c.ctx.Call("fill")
	c.ctx.Call("closePath")
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
