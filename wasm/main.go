// go:build js,wasm
package main

import (
	pcanvas "github.com/mvrilo/go-particles/canvas"
	"github.com/mvrilo/go-particles/particles"
)

func main() {
	// particle configuration
	conf := particles.DefaultConfig
	conf.Color = "#9999FF"
	conf.Speed = 1.5
	conf.Size = 1.4

	// canvas setup
	canvas := pcanvas.NewCanvas("particles", 60, "#222", 100, conf)
	canvas.ListenEvents()
	canvas.Fullscreen()
	canvas.Start()
	canvas.AppendElement()
	canvas.Render()
}
