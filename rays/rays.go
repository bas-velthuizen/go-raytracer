package rays

import (
	"fmt"
	"math"

	"github.com/bas-velthuizen/go-raytracer/matrix"
	"github.com/bas-velthuizen/go-raytracer/spheres"
	"github.com/bas-velthuizen/go-raytracer/tuples"
)

// Ray describes a ray of light with an origin and direction
type Ray struct {
	Origin    tuples.Tuple
	Direction tuples.Tuple
}

// NewRay creates a new ray from a origin and direction
func NewRay(origin, direction tuples.Tuple) *Ray {
	return &Ray{Origin: origin, Direction: direction}
}

// String formats the ray as a string
func (r Ray) String() string {
	return fmt.Sprintf("Ray( %v, %v )", r.Origin, r.Direction)
}

// Position calculates the position of the ray at time t
func (r Ray) Position(t float64) *tuples.Tuple {
	result := r.Origin.Add(r.Direction.Multiply(t))
	return &result
}

// Intersect calculates the intersections with a Sphere
func (r Ray) Intersect(s *spheres.Sphere) Intersections {
	rTransformed := r.Transform(*s.Transform.Inverse())

	sphereToRay := rTransformed.Origin.Subtract(s.Center)

	a := rTransformed.Direction.Dot(rTransformed.Direction)
	b := 2 * rTransformed.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1.0

	discriminant := b*b - 4*a*c

	if discriminant < 0.0 {
		return *NewIntersections([]*Intersection{})
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	if t1 > t2 {
		t1, t2 = t2, t1
	}

	intersection1 := NewIntersection(t1, s)
	intersection2 := NewIntersection(t2, s)

	return *NewIntersections([]*Intersection{intersection1, intersection2})
}

// Transform transforms a ray with a matrix, returnning a new ray
func (r Ray) Transform(m matrix.Matrix) *Ray {
	newOrigin := m.MultiplyTuple(r.Origin)
	newDirection := m.MultiplyTuple(r.Direction)
	return NewRay(*newOrigin, *newDirection)
}
