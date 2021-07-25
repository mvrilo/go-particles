package particles

// Group represents multiple particles
type Group struct {
	Max       int
	Width     int
	Height    int
	Particles Particles
}

// NewGroup initializes a new group of particles
func NewGroup(width, height, max int, config Config) (group *Group) {
	var particles Particles

	for i := 0; i < max; i++ {
		particles = append(particles, NewParticle(i, config))
	}

	group = &Group{
		Max:       max,
		Width:     width,
		Height:    height,
		Particles: particles,
	}
	return
}

// Move moves particles
func (g *Group) Move() {
	for _, particle := range g.Particles {
		particle.Move()

		for _, particle2 := range g.Particles {
			dist := particle.Distance(particle2)
			size := particle.Size + particle2.Size

			if dist > particle.Area || dist > particle2.Area {
				continue
			}

			if particle.Config.Bounce && dist < size {
				particle.Reverse()
				particle2.Reverse()
			}
		}

		particle.Bounds(float64(g.Width), float64(g.Height))
	}
}
