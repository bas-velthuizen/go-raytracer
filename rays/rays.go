package rays

import (
	"fmt"
	"github.com/bas-velthuizen/go-raytracer/tuples"
)

type ray struct {
	origin tuples.Tuple
	direction tuples.Tuple
}

// Ray creates a new ray from a origin and direction
func Ray(origin, direction tuples.Tuple) (* ray) {
	return &ray{origin:origin, direction:direction}
}

// String formats the ray as a string
func (r ray) String() string {
	return fmt.Sprintf("Ray( %v, %v )", r.origin, r.direction)
}

// Position calculates the position of the ray at time t
func (r ray) Position(t float64) *tuples.Tuple {
	result := r.origin.Add(r.direction.Multiply(t))
	return &result
}


