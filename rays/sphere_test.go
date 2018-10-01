package rays

import (
	"testing"

	"github.com/bas-velthuizen/go-raytracer/transformations"

	"github.com/bas-velthuizen/go-raytracer/matrix"
	"github.com/bas-velthuizen/go-raytracer/tuples"
)

// Scenario: A sphere's default transformation
// Given s ← sphere()
// Then s.transform = identity_matrix
func Test_Sphere_has_Default_Transformation(t *testing.T) {
	// Given
	s := NewUnitSphere()
	// Expected
	wanted := matrix.Identity(4)
	// Then
	if !wanted.Equals(s.transform) {
		t.Errorf("Transform of %v = %v, expected %v", s, s.transform, wanted)
	}
}

// Scenario: Changing a sphere's transformation
// Given s ← sphere()
// And t ← translation(2, 3, 4)
// When set_transform(s, t)
// Then s.transform = t
func Test_Changing_a_Sphere_s_Transformation(t *testing.T) {
	// Given
	s := NewUnitSphere()
	// And
	trans := transformations.Translation(2, 3, 4)
	// When
	s.SetTransform(trans)
	// Then
	if !trans.Equals(s.transform) {
		t.Errorf("Transform of %v = %v, expected %v", s, s.transform, trans)
	}
}

// Scenario: Intersecting a scaled sphere with a ray
// Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
// And s ← sphere()
// When set_transform(s, scaling(2, 2, 2))
// And xs ← intersect(s, r)
// Then xs.count = 2
// And xs[0].t = 3
// And xs[1].t = 7
func Test_Intersecting_a_Scaled_Sphere_with_a_Ray(t *testing.T) {
	// Given
	r := NewRay(tuples.Point(0,0,-5), tuples.Vector(0,0,1))
	// And
	s := NewUnitSphere()
	// When
	s.SetTransform(transformations.Scaling(2, 2, 2))
	// And
	xs := r.Intersect(s)
	// Expected
	wantedCount := 2
	wanted0 := 3.0
	wanted1 := 7.0
	// Then
	if wantedCount != len(xs) {
		t.Errorf("len(%v) = %d, expected %d", xs, len(xs), wantedCount)
	}
	// And
	if wanted0 != (*xs[0]).Time {
		t.Errorf("(%v).Time = %9.6f, expected %9.6f", *xs[0], (*xs[0]).Time, wanted0)
	}
	// And
	if wanted1 != (*xs[1]).Time {
		t.Errorf("(%v).Time = %9.6f, expected %9.6f", *xs[1], (*xs[1]).Time, wanted1)
	}
}

// Scenario: Intersecting a translated sphere with a ray
// Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
// And s ← sphere()
// When set_transform(s, translation(5, 0, 0))
// And xs ← intersect(s, r)
// Then xs.count = 0
func Test_Intersecting_a_Translated_Sphere_with_a_Ray(t *testing.T) {
	// Given
	r := NewRay(tuples.Point(0,0,-5), tuples.Vector(0,0,1))
	// And
	s := NewUnitSphere()
	// When
	s.SetTransform(transformations.Translation(5, 0, 0))
	// And
	xs := r.Intersect(s)
	// Expected
	wantedCount := 0
	// Then
	if wantedCount != len(xs) {
		t.Errorf("len(%v) = %d, expected %d", xs, len(xs), wantedCount)
	}
}
