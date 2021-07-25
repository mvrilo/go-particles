package particles

// Particles is a list of particles
type Particles []*Particle

// Particle represents a single particle
type Particle struct {
	Config
	ID        int
	Position  Vector
	Direction Vector
}

// Config contains data for particles configuration
type Config struct {
	Speed  float64
	Area   float64
	Size   float64
	Color  string
	Bounce bool
}

// DefaultConfig is a default value for config
var DefaultConfig = Config{2.0, 40.0, 2.0, "blue", true}

// NewParticle initializes a new particle
func NewParticle(id int, config Config) *Particle {
	return &Particle{
		ID:        id,
		Config:    config,
		Direction: randVector(-100, 100),
	}
}

// Bounce bounces particles when too close
func (p *Particle) Reverse() {
	p.Direction[0] = -p.Direction[0]
	p.Direction[1] = -p.Direction[1]
}

// Distance returns the distance between two particles
func (p *Particle) Distance(pa *Particle) float64 {
	return p.Position.Distance(pa.Position)
}

// Nearby returns a list of particles within its area
func (p *Particle) Nearby(particles Particles) (last Particles) {
	for _, particle := range particles {
		if p.Distance(particle) < p.Area {
			last = append(last, p)
		}
	}
	return
}

// Move moves a particle
func (p *Particle) Move() {
	p.Position[0] += (p.Direction[0] * p.Speed)
	p.Position[1] += (p.Direction[1] * p.Speed)
}

// Bounds checks if a particle is within bounds
func (p *Particle) Bounds(maxx, maxy float64) {
	if p.Position[0] > maxx {
		p.Position[0] = 0.0
	}
	if p.Position[1] > maxy {
		p.Position[1] = 0.0
	}

	if p.Position[0] < 0.0 {
		p.Position[0] = maxx
	}
	if p.Position[1] < 0.0 {
		p.Position[1] = maxy
	}
}
