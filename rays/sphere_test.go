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
	s := NewSphere(tuples.Point(0, 0, 0), 1.0)
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
	s := NewSphere(tuples.Point(0, 0, 0), 1.0)
	// And
	trans := transformations.Translation(2, 3, 4)
	// When
	s.SetTransform(trans)
	// Then
	if !trans.Equals(s.transform) {
		t.Errorf("Transform of %v = %v, expected %v", s, s.transform, trans)
	}
}
