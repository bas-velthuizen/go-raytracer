package rays

import "github.com/bas-velthuizen/go-raytracer/tuples"

type ray struct {
	origin tuples.Tuple
	direction tuples.Tuple
}

// Ray creates a new ray from a origin and direction
func Ray(origin, direction tuples.Tuple) (* ray) {
	return &ray{origin:origin, direction:direction}
}
