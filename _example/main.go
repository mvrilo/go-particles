package main

import (
	"fmt"

	"github.com/mvrilo/go-particles"
)

func main() {
	group := particles.NewGroup(
		400.0,
		400.0,
		100,
		20.0,
		4.0,
		"red",
	)

	fmt.Printf("before %+v\n", group.Particles[0])
	group.Move()
	fmt.Printf("after %+v\n", group.Particles[0])
}
