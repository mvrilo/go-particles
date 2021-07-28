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

	// canvas setup
	canvas := pcanvas.NewCanvas("particles", 80, "#222", 40, conf)
	go canvas.ListenEvents()
	canvas.Fullscreen()
	canvas.Start()
	canvas.AppendElement()
	canvas.Render()
}
