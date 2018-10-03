package rays

import (
	"math"
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
	r := NewRay(tuples.Point(0, 0, -5), tuples.Vector(0, 0, 1))
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
	r := NewRay(tuples.Point(0, 0, -5), tuples.Vector(0, 0, 1))
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

// Scenario: The normal on a sphere at a point on the x axis
// Given s ← sphere()
// When n ← normal_at(s, point(1, 0, 0))
// Then n = vector(1, 0, 0)
func Test_Normal_of_a_Sphere_at_a_Point_on_the_X_Axis(t *testing.T) {
	// Given
	s := NewUnitSphere()
	// And
	point := tuples.Point(1, 0, 0)
	// When
	n := s.NormalAt(point)
	// Expected
	wanted := tuples.Vector(1, 0, 0)
	// Then
	if !wanted.Equals(*n) {
		t.Errorf("normal_at( %v, %v ) = %v , wanted %v", s, point, n, wanted)
	}
}

// Scenario: The normal on a sphere at a point on the y axis
// Given s ← sphere()
// When n ← normal_at(s, point(0, 1, 0))
// Then n = vector(0, 1, 0)
func Test_Normal_of_a_Sphere_at_a_Point_on_the_Y_Axis(t *testing.T) {
	// Given
	s := NewUnitSphere()
	// And
	point := tuples.Point(0, 1, 0)
	// When
	n := s.NormalAt(point)
	// Expected
	wanted := tuples.Vector(0, 1, 0)
	// Then
	if !wanted.Equals(*n) {
		t.Errorf("normal_at( %v, %v ) = %v , wanted %v", s, point, n, wanted)
	}
}

// Scenario: The normal on a sphere at a point on the z axis
// Given s ← sphere()
// When n ← normal_at(s, point(0, 0, 1))
// Then n = vector(0, 0, 1)
func Test_Normal_of_a_Sphere_at_a_Point_on_the_Z_Axis(t *testing.T) {
	// Given
	s := NewUnitSphere()
	// And
	point := tuples.Point(0, 0, 1)
	// When
	n := s.NormalAt(point)
	// Expected
	wanted := tuples.Vector(0, 0, 1)
	// Then
	if !wanted.Equals(*n) {
		t.Errorf("normal_at( %v, %v ) = %v , wanted %v", s, point, n, wanted)
	}
}

// Scenario: The normal on a sphere at a non-axial point
// Given s ← sphere()
// When n ← normal_at(s, point(√3/3, √3/3, √3/3))
// Then n = vector(√3/3, √3/3, √3/3)
func Test_Normal_of_a_Sphere_at_a_NonAxial_Point(t *testing.T) {
	// Given
	s := NewUnitSphere()
	// And
	val := math.Sqrt(3) / 3.0
	point := tuples.Point(val, val, val)
	// When
	n := s.NormalAt(point)
	// Expected
	wanted := tuples.Vector(val, val, val)
	// Then
	if !wanted.Equals(*n) {
		t.Errorf("normal_at( %v, %v ) = %v , wanted %v", s, point, n, wanted)
	}
}

// Scenario: The normal is a normalized vector
// Given s ← sphere()
// When n ← normal_at(s, point(√3/3, √3/3, √3/3))
// Then n = normalize(n)
func Test_Normal_is_a_Normalized_Vector(t *testing.T) {
	// Given
	s := NewUnitSphere()
	// And
	val := math.Sqrt(3) / 3.0
	point := tuples.Point(val, val, val)
	// When
	n := s.NormalAt(point)
	// Expected
	wanted := n.Normalize()
	// And
	if !n.IsVector() {
		t.Errorf("normal_at( %v, %v ) is Point , wanted Vector", s, point)
	}
	// Then
	if !wanted.Equals(*n) {
		t.Errorf("normal_at( %v, %v ) = %v (normalized) , wanted %v", s, point, n, wanted)
	}
}

// Scenario: Computing the normal on a translated sphere
// Given s ← sphere()
// And set_transform(s, translation(0, 1, 0))
// When n ← normal_at(s, point(0, 1.70711, -0.70711))
// Then n = vector(0, 0.70711, -0.70711)
func Test_Computing_the_Normal_on_a_Translated_Sphere(t *testing.T) {
	// Given
	s := NewUnitSphere()
	// And
	s.SetTransform(transformations.Translation(0, 1, 0))
	// And
	point := tuples.Point(0, 1.70711, -0.70711)
	// When
	n := s.NormalAt(point)
	// Expected
	wanted := tuples.Vector(0, 0.70711, -0.70711)
	// Then
	if !wanted.Equals(*n) {
		t.Errorf("normal_at( %v, %v ) = %v , wanted %v", s, point, n, wanted)
	}
}

// Scenario: Computing the normal on a scaled sphere
// Given s ← sphere()
// And set_transform(s, scaling(1, 0.5, 1))
// When n ← normal_at(s, point(0, √2/2, -√2/2))
// Then n = vector(0, 0.97014, -0.24254)
func Test_Computing_the_Normal_on_a_scaled_Sphere(t *testing.T) {
	// Given
	s := NewUnitSphere()
	// And
	s.SetTransform(transformations.Scaling(1, 0.5, 1))
	// And
	point := tuples.Point(0, math.Sqrt(2)/2.0, -math.Sqrt(2)/2.0)
	// When
	n := s.NormalAt(point)
	// Expected
	wanted := tuples.Vector(0, 0.97014, -0.24254)
	// Then
	if !wanted.Equals(*n) {
		t.Errorf("normal_at( %v, %v ) = %v , wanted %v", s, point, n, wanted)
	}
}
