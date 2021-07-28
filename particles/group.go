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
		particle := NewParticle(i, config)
		particle.RandomizePosition(width, height)
		particles = append(particles, particle)
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
		if !particle.Config.Move {
			continue
		}

		particle.Move()

		for _, particle2 := range g.Particles {
			dist := particle.Distance(particle2)
			if dist > particle.Area {
				continue
			}

			if particle.Config.Bounce && dist < particle.Size*2 {
				particle.Reverse()
				particle2.Reverse()
			}
		}

		w, h := float64(g.Width), float64(g.Height)

		if particle.Config.Bounds {
			particle.Bounds(w, h)
			continue
		}

		particle.Bounce(w, h)
	}
}
