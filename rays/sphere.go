package rays

import (
	"fmt"
	"github.com/bas-velthuizen/go-raytracer/tuples"
)

// Sphere describes a sphere shape
type Sphere struct {
	center tuples.Tuple
	radius float64
}

// NewSphere creates a new Sphere instance
func NewSphere(center tuples.Tuple, radius float64) *Sphere {
	return &Sphere{center, radius}
}

// String formats Object to readable string
func (s Sphere) String() string {
	return fmt.Sprintf("Sphere( %v, %v )", s.center, s.radius)
}
