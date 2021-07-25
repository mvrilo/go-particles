package particles

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Vector [2]float64

func randf(min, max int) float64 {
	return float64(rand.Intn(max) + min)
}

// Distance returns a distance between two vectors
func (v Vector) Distance(v2 Vector) float64 {
	return math.Sqrt(math.Pow(v2[0]-v[0], 2) + math.Pow(v2[1]-v[1], 2))
}

func randVector(min, max int) Vector {
	x := randf(min, max) / 100.0
	y := randf(min, max) / 100.0
	return Vector{x, y}
}
