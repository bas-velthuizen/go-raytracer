package rays

import (
	"fmt"
	"math"

	"github.com/bas-velthuizen/go-raytracer/matrix"
	"github.com/bas-velthuizen/go-raytracer/tuples"
)

// Ray describes a ray of light with an origin and direction
type Ray struct {
	origin    tuples.Tuple
	direction tuples.Tuple
}

// NewRay creates a new ray from a origin and direction
func NewRay(origin, direction tuples.Tuple) *Ray {
	return &Ray{origin: origin, direction: direction}
}

// String formats the ray as a string
func (r Ray) String() string {
	return fmt.Sprintf("Ray( %v, %v )", r.origin, r.direction)
}

// Position calculates the position of the ray at time t
func (r Ray) Position(t float64) *tuples.Tuple {
	result := r.origin.Add(r.direction.Multiply(t))
	return &result
}

// Intersect calculates the intersections with a Sphere
func (r Ray) Intersect(s *Sphere) Intersections {
	rTransformed := r.Transform( *s.transform.Inverse())

	sphereToRay := rTransformed.origin.Subtract(s.center)

	a := rTransformed.direction.Dot(rTransformed.direction)
	b := 2 * rTransformed.direction.Dot(sphereToRay)
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
	newOrigin := m.MultiplyTuple(r.origin)
	newDirection := m.MultiplyTuple(r.direction)
	return NewRay(*newOrigin, *newDirection)
}
