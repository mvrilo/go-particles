// +build js,wasm
package main

import (
	pcanvas "github.com/mvrilo/go-particles/canvas"
)

func main() {
	canvas := pcanvas.NewCanvas("particles", 60, "#FFF", 20)
	go canvas.ListenEvents()
	canvas.Fullscreen()
	canvas.Start()
	canvas.AppendElement()
	canvas.Render()
}
