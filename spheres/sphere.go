package spheres

import (
	"fmt"
	"math"

	"github.com/bas-velthuizen/go-raytracer/materials"

	"github.com/bas-velthuizen/go-raytracer/matrix"
	"github.com/bas-velthuizen/go-raytracer/tuples"
)

// Sphere describes a sphere shape
type Sphere struct {
	Center    tuples.Tuple
	Radius    float64
	Transform matrix.Matrix
	Material  materials.Material
}

// NewSphere creates a new Sphere instance
func NewSphere(center tuples.Tuple, radius float64) *Sphere {
	return &Sphere{center, radius, *matrix.Identity(4), materials.DefaultMaterial()}
}

// NewUnitSphere creates a new Sphere instance
func NewUnitSphere() *Sphere {
	return NewSphere(tuples.Point(0, 0, 0), 1.0)
}

// String formats Object to readable string
func (s Sphere) String() string {
	return fmt.Sprintf("Sphere( %v, %v, %v, %v )", s.Center, s.Radius, s.Transform, s.Material)
}

// SetTransform sets the transform value of the sphere
func (s *Sphere) SetTransform(transform *matrix.Matrix) {
	s.Transform = *transform
	// fmt.Printf("sphere with new transform: %v\n\n", s)
}

// NormalAt calculates the normal vector on a sphere at a certain world point
func (s Sphere) NormalAt(worldPoint tuples.Tuple) *tuples.Tuple {
	objectPoint := s.Transform.Inverse().MultiplyTuple(worldPoint)
	objectNormal := objectPoint.Subtract(s.Center)
	worldNormal := s.Transform.Inverse().Transpose().MultiplyTuple(objectNormal)
	worldNormal.W = 0
	normal := worldNormal.Normalize()
	return &normal
}

// Equals checks if another sphere is equal to the current sphere
func (s Sphere) Equals(other Sphere) bool {
	return s.Center.Equals(other.Center) &&
		s.Material.Equals(other.Material) &&
		(math.Abs(s.Radius-other.Radius) < tuples.Epsilon) &&
		s.Transform.Equals(other.Transform)
}
