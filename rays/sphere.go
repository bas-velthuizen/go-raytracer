package rays

import (
	"fmt"

	"github.com/bas-velthuizen/go-raytracer/matrix"
	"github.com/bas-velthuizen/go-raytracer/tuples"
)

// Sphere describes a sphere shape
type Sphere struct {
	center    tuples.Tuple
	radius    float64
	transform matrix.Matrix
}

// NewSphere creates a new Sphere instance
func NewSphere(center tuples.Tuple, radius float64) *Sphere {
	return &Sphere{center, radius, *matrix.Identity(4)}
}

// NewUnitSphere creates a new Sphere instance
func NewUnitSphere() *Sphere {
	return NewSphere(tuples.Point(0, 0, 0), 1.0)
}

// String formats Object to readable string
func (s Sphere) String() string {
	return fmt.Sprintf("Sphere( %v, %v, %v )", s.center, s.radius, s.transform)
}

// SetTransform sets the transform value of the sphere
func (s *Sphere) SetTransform(transform *matrix.Matrix) {
	s.transform = *transform
	fmt.Printf("sphere with new transform: %v\n\n", s)
}

// NormalAt calculates the normal vector on a sphere at a certain world point
func (s Sphere) NormalAt(worldPoint tuples.Tuple) *tuples.Tuple {
	objectPoint := s.transform.Inverse().MultiplyTuple(worldPoint)
	objectNormal := objectPoint.Subtract(s.center)
	worldNormal := s.transform.Inverse().Transpose().MultiplyTuple(objectNormal)
	worldNormal.W = 0
	normal := worldNormal.Normalize()
	return &normal
}
