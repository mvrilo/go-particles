package particles

// Config contains data for particles configuration
type Config struct {
	Speed  float64
	Area   float64
	Size   float64
	Color  string
	Bounds bool
	Bounce bool
	Move   bool
}

// DefaultConfig is a default value for config
var DefaultConfig = Config{2.0, 80.0, 1.8, "#ccccFF", true, true, true}

// Particles is a list of particles
type Particles []*Particle

// Particle represents a single particle
type Particle struct {
	Config
	ID        int
	Position  Vector
	Direction Vector
}

// NewParticle initializes a new particle
func NewParticle(id int, config Config) *Particle {
	return &Particle{
		ID:     id,
		Config: config,
		Direction: Vector{
			randf(-100, 100) / 100.0,
			randf(-100, 100) / 100.0,
		},
	}
}

// RandomizePosition sets a random (within a given range) position for the particle
func (p *Particle) RandomizePosition(width, height int) {
	p.Position = Vector{
		randf(0, width),
		randf(0, height),
	}
}

// ReverseX makes particles go the other way
func (p *Particle) ReverseX() {
	p.Direction[0] = -p.Direction[0]
}

// ReverseY makes particles go the other way
func (p *Particle) ReverseY() {
	p.Direction[1] = -p.Direction[1]
}

// Reverse makes particles go the other way
func (p *Particle) Reverse() {
	p.ReverseX()
	p.ReverseY()
}

// Distance returns the distance between two particles
func (p *Particle) Distance(pa *Particle) float64 {
	return p.Position.Distance(pa.Position)
}

// Move moves a particle
func (p *Particle) Move() {
	p.Position[0] += (p.Direction[0] * p.Speed)
	p.Position[1] += (p.Direction[1] * p.Speed)
}

// Bounce bounces a particle
func (p *Particle) Bounce(maxx, maxy float64) {
	if p.Position[0] > maxx || p.Position[0] < 0.0 {
		p.ReverseX()
	}
	if p.Position[1] > maxy || p.Position[1] < 0.0 {
		p.ReverseY()
	}
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
